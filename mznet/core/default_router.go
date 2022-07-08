/**
 * @Author: dingQingHui
 * @Description:
 * @File: irouter
 * @Version: 1.0.0
 * @Date: 2022/7/7 17:32
 */

package core

import (
	miface2 "mz/mznet/miface"
)

type defaultRouter struct {
}

func NewDefaultRouter() miface2.IRouter {
	return &defaultRouter{}
}
func (r *defaultRouter) OnConnected(connection miface2.IConnection) {
}
func (r *defaultRouter) OnDisconnect(connection miface2.IConnection) {
}
func (r *defaultRouter) OnProcess(connection miface2.IConnection, message miface2.IMessage) {
}
