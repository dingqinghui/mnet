/**
 * @Author: dingQingHui
 * @Description:
 * @File: server
 * @Version: 1.0.0
 * @Date: 2022/7/8 18:57
 */

package core

import (
	"context"
	"github.com/dingqinghui/mz/mznet/miface"
	"sync/atomic"
)

type (
	Server struct {
		Options Options
		Stop    atomic.Value
		Ctx     context.Context // 通知子连接退出
		Cancel  context.CancelFunc
	}
)

func NewServer(options Options) *Server {
	s := &Server{
		Options: options,
	}
	s.init()

	return s
}
func (s *Server) init() {
	s.Stop.Store(false)
	s.Ctx, s.Cancel = context.WithCancel(s.Options.ParentCtx)
}

func (s *Server) OnConnected(connection miface.IConnection) {
	s.Options.Router.OnConnected(connection)
}
func (s *Server) OnDisconnect(connection miface.IConnection) {
	s.Options.Router.OnDisconnect(connection)
}
func (s *Server) OnProcess(connection miface.IConnection, message miface.IPackage) {
	s.Options.Router.OnProcess(connection, message)
}

func (s *Server) GetRouter() miface.IRouter {
	return s.Options.Router
}

func (s *Server) IsClose() bool {
	return s.Stop.CompareAndSwap(true, true)
}

//
// Destroy
// @Description: 析构
// @receiver s
// @return bool
//
func (s *Server) Destroy() bool {
	if s.IsClose() {
		return false
	}

	s.Stop.Store(true)

	return true
}

//
// Close
// @Description: 主动关闭
// @receiver s
//
func (s *Server) Close() bool {
	if s.IsClose() {
		return false
	}

	s.Cancel()
	return true
}
