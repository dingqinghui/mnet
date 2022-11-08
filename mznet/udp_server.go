/**
 * @Author: dingQingHui
 * @Description:
 * @File: udp_server
 * @Version: 1.0.0
 * @Date: 2022/10/28 14:48
 */

package mznet

import (
	"net"
)

type udpServer struct {
	config     *ServerConfig
	listener   *net.UDPConn
	udpAddrMap map[string]IConnection
	ICloser
}

func newUdpServer(config *ServerConfig) IServer {
	s := &udpServer{
		config:     config,
		udpAddrMap: make(map[string]IConnection),
		ICloser:    defaultCloser,
	}
	return s
}

func (s *udpServer) RunEventLoop() error {
	if err := s.listen(); err != nil {
		return err
	}
	go s.accept()
	return nil
}

func (s *udpServer) listen() error {
	udpAddr, err := net.ResolveUDPAddr(s.config.Network, s.config.ListenAddress)
	if err != nil {
		return err
	}
	listener, err := net.ListenUDP(s.config.Network, udpAddr)
	if err != nil {
		return err
	}
	s.listener = listener
	return nil
}

func (s *udpServer) Close() error {
	if err := s.ICloser.Close(); err != nil {
		return err
	}
	if s.listener != nil {
		return s.listener.Close()
	}
	return nil
}

func (s *udpServer) accept() {
	for true {
		b := make([]byte, 1024)
		n, addr, err := s.listener.ReadFromUDP(b)
		if err != nil {
			s.config.EventListener.OnError(nil, err)
			return
		}
		if _, ok := s.udpAddrMap[addr.String()]; !ok {
			c := newUdpConnection(s.listener, AcceptConnection, addr.String(), s.config.Network, s.config.EventListener)
			s.udpAddrMap[addr.String()] = c
			s.config.EventListener.OnConnected(c)
		}
		con := s.udpAddrMap[addr.String()]
		udpCon, ok := con.(*udpConnection)
		if !ok {
			s.config.EventListener.OnError(nil, err)
			return
		}
		udpCon.recUdpStream(b[:n])
	}
}
