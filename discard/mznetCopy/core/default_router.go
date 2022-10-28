/**
 * @Author: dingQingHui
 * @Description:
 * @File: irouter
 * @Version: 1.0.0
 * @Date: 2022/7/7 17:32
 */

package core

import (
	miface "github.com/dingqinghui/mz/mznet/miface"
)

type defaultRouter struct {
}

func NewDefaultRouter() miface.IRouter {
	return &defaultRouter{}
}
func (r *defaultRouter) OnConnected(connection miface.IConnection) {
}
func (r *defaultRouter) OnDisconnect(connection miface.IConnection) {
}
func (r *defaultRouter) OnProcess(connection miface.IConnection, iPackage miface.IPackage) {
}
