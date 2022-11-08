/**
 * @Author: dingQingHui
 * @Description:
 * @File: tcpServer
 * @Version: 1.0.0
 * @Date: 2022/10/28 10:53
 */

package mznet

import (
	"net"
)

type tcpServer struct {
	config   *ServerConfig
	listener net.Listener
	ICloser
}

func NewTcpServer(config *ServerConfig) IServer {
	s := &tcpServer{
		config:  config,
		ICloser: defaultCloser,
	}
	return s
}

func (s *tcpServer) RunEventLoop() error {
	if err := s.listen(); err != nil {
		return err
	}
	go s.accept()
	return nil
}

func (s *tcpServer) listen() error {
	listener, err := net.Listen(s.config.Network, s.config.ListenAddress)
	if err != nil {
		return err
	}
	s.listener = listener
	return nil
}

func (s *tcpServer) accept() {
	for true {
		c, err := s.listener.Accept()
		if err != nil {
			s.config.EventListener.OnError(nil, err)
			return
		}
		_ = c
		con := newTcpConnection(c, AcceptConnection, s.config.EventListener)
		s.config.EventListener.OnConnected(con)
	}
}

func (s *tcpServer) Close() error {
	if err := s.ICloser.Close(); err != nil {
		return err
	}
	if s.listener != nil {
		return s.listener.Close()
	}
	return nil
}
