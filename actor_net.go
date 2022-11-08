/**
 * @Author: dingQingHui
 * @Description:
 * @File: net
 * @Version: 1.0.0
 * @Date: 2022/10/28 23:49
 */

package main

import (
	"github.com/dingqinghui/mz/actor"
	"github.com/dingqinghui/mz/mznet"
)

type NetConnected struct {
	connection mznet.IConnection
}

type NetClosed struct {
	connection mznet.IConnection
}

type NetProcess struct {
	connection mznet.IConnection
	msg        interface{}
}

type NetError struct {
	connection mznet.IConnection
	err        error
	users      []interface{}
}

type eventForwardListener struct {
	pid actor.IPid
}

func newEventForwardListener(pid actor.IPid) mznet.IEventListener {
	return &eventForwardListener{pid: pid}
}

func (e *eventForwardListener) OnConnected(connection mznet.IConnection) bool {
	_ = e.pid.SendUserMessage(&NetConnected{connection: connection})
	return true
}
func (e *eventForwardListener) OnProcess(connection mznet.IConnection, msg interface{}) bool {
	_ = e.pid.SendUserMessage(&NetProcess{connection: connection, msg: msg})
	return true
}
func (e *eventForwardListener) OnClosed(connection mznet.IConnection) bool {
	_ = e.pid.SendUserMessage(&NetClosed{connection: connection})
	return true
}
func (e *eventForwardListener) OnError(connection mznet.IConnection, err error, users ...interface{}) bool {
	_ = e.pid.SendUserMessage(&NetError{connection: connection, err: err, users: users})
	return false
}

func NewClient(pid actor.IPid, config *mznet.ClientConfig) mznet.IClient {
	listener := newEventForwardListener(pid)
	config.EventListener = listener
	return mznet.NewClient(config)
}

func NewServer(pid actor.IPid, config *mznet.ServerConfig) mznet.IServer {
	listener := newEventForwardListener(pid)
	config.EventListener = listener
	return mznet.NewServer(config)
}
