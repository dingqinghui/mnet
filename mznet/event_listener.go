/**
 * @Author: dingQingHui
 * @Description:
 * @File: event_listener
 * @Version: 1.0.0
 * @Date: 2022/10/28 10:56
 */

package mznet

type IEventListener interface {
	OnConnected(IConnection) bool
	OnProcess(IConnection, interface{}) bool
	OnClosed(IConnection) bool
	OnError(IConnection, error, ...interface{}) bool
}

type DefaultEventListener struct{}

func (DefaultEventListener) OnConnected(IConnection) bool {
	//TODO implement me
	panic("implement me")
}

func (DefaultEventListener) OnProcess(IConnection, interface{}) bool {
	//TODO implement me
	panic("implement me")
}

func (DefaultEventListener) OnClosed(IConnection) bool {
	//TODO implement me
	panic("implement me")
}

func (DefaultEventListener) OnError(IConnection, error, ...interface{}) bool {
	//TODO implement me
	panic("implement me")
}
