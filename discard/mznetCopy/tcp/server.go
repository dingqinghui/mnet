/**
 * @Author: dingQingHui
 * @Description:
 * @File: server
 * @Version: 1.0.0
 * @Date: 2022/7/7 15:29
 */

package tcp

import (
	"github.com/dingqinghui/mz/mznet/core"
	"github.com/dingqinghui/mz/mznet/miface"
	"log"
	"net"
)

type (
	server struct {
		*core.Server
		listener net.Listener
	}
)

func NewServer(options core.Options) (miface.IServer, error) {
	s := &server{}
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

func (s *server) listen() error {
	listener, err := net.Listen(s.Options.Network, s.Options.Address)
	if err != nil {
		return err
	}

	s.listener = listener

	log.Printf("server lisent address:%s network:%s", s.Options.Network, s.Options.Address)
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

func (s *server) accept() (miface.IConnection, error) {
	c, err := s.listener.Accept()
	if err != nil {
		return nil, err
	}
	return NewConnection(s.Options.Network, c, miface.TypeConnectionAccept, s.genOptions()), nil
}

func (s *server) genOptions() core.Options {
	options := s.Options
	// 重新赋值路由，回调到server进行连接管理
	options.Router = s
	// context
	options.ParentCtx = s.Ctx
	return options
}
