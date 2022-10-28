/**
 * @Author: dingQingHui
 * @Description:
 * @File: serverUdp
 * @Version: 1.0.0
 * @Date: 2022/7/8 13:44
 */

package udp

import (
	"errors"
	"github.com/dingqinghui/mz/mznet/core"
	"github.com/dingqinghui/mz/mznet/miface"
	"log"
	"net"
)

type server struct {
	*core.Server
	listener *net.UDPConn
	udpMap   map[string]miface.IConnection
}

func NewUdpServer(options core.Options) (miface.IServer, error) {
	s := &server{
		udpMap: make(map[string]miface.IConnection),
	}
	s.Server = core.NewServer(options)

	if err := s.listen(); err != nil {
		return nil, err
	}
	go s.waitExit()
	return s, nil
}

func (s *server) Accept() error {
	return s.accept()
}

func (s *server) Close() error {
	if s.Server.Destroy() {
		return nil
	}
	_ = s.listener.Close()
	return nil
}

func (s *server) waitExit() {
	select {
	case <-s.Options.ParentCtx.Done(): // 父节点退出
		s.Destroy()
	case <-s.Ctx.Done(): // 本节点退出
		s.Destroy()
	}
}

func (s *server) listen() error {
	udpAddr, err := net.ResolveUDPAddr(s.Options.Network, s.Options.Address)
	if err != nil {
		return err
	}
	listener, err := net.ListenUDP(s.Options.Network, udpAddr)
	if err != nil {
		return err
	}
	s.listener = listener
	log.Printf("server lisent address:%s network:%s", s.Options.Network, s.Options.Address)
	return nil
}

func (s *server) accept() error {
	b := make([]byte, 1024)
	n, addr, err := s.listener.ReadFrom(b)
	if err != nil {
		return err
	}
	if _, ok := s.udpMap[addr.String()]; !ok {
		c := newConnection(s.Options.Network, s.listener, miface.TypeConnectionAccept, addr, s.Options)
		s.udpMap[addr.String()] = c
	}
	con := s.udpMap[addr.String()]
	udpCon, ok := con.(*connection)
	if !ok {
		return errors.New("change type fail")
	}
	udpCon.RevMsg(miface.NewMessage(uint16(n), b[:n]))
	return nil
}
