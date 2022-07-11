/**
 * @Author: dingQingHui
 * @Description:
 * @File: serverUdp
 * @Version: 1.0.0
 * @Date: 2022/7/8 13:44
 */

package udp

import (
	"log"
	"mz/mznet/core"
	"mz/mznet/miface"
	"net"
)

type server struct {
	*core.Server
	listener *net.UDPConn
	udpMap   map[string]miface.IConnection
}

func NewUdpServer(options core.Options) miface.IServer {
	s := &server{
		udpMap: make(map[string]miface.IConnection),
	}

	s.Server = core.NewServer(options)
	return s
}

func (s *server) Run() error {
	if err := s.listen(); err != nil {
		return err
	}
	go s.waitExit()
	go s.accept()
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

func (s *server) accept() {
	defer s.Close()
	for true {
		b := make([]byte, 1024)
		n, addr, err := s.listener.ReadFrom(b)
		if err != nil {
			return
		}

		if _, ok := s.udpMap[addr.String()]; !ok {
			options := s.Options
			options.UdpAddr = addr
			c := newConnection(s.Options.Network, s.listener, miface.TypeConnectionAccept, options)
			s.udpMap[addr.String()] = c
		}
		con := s.udpMap[addr.String()]
		udpCon, ok := con.(*connection)
		if !ok {
			log.Printf("change type connection fail")
			continue
		}
		udpCon.RevMsg(core.NewPackage(uint32(n), b[:n]))
	}
}

func (s *server) Destroy() {
	if s.Server.Destroy() {
		return
	}
	_ = s.listener.Close()
}
