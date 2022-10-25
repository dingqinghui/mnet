/**
 * @Author: dingQingHui
 * @Description:
 * @File: net
 * @Version: 1.0.0
 * @Date: 2022/9/27 10:05
 */

package actor

import (
	"github.com/dingqinghui/mz/actor/iface"
	"github.com/dingqinghui/mz/mznet"
	"github.com/dingqinghui/mz/mznet/core"
	"github.com/dingqinghui/mz/mznet/miface"
	"log"
)

type ActorMessageType int

var (
	ActorMessageSocket ActorMessageType = 1
	ActorMessageRpc    ActorMessageType = 2
)

type IActorMessage interface {
	GetType() ActorMessageType
	GetData() []byte
}
type (
	Protocol struct {
		msgType  iface.ActorMessageType
		parse    iface.IParse
		dispatch iface.DispatchFun
	}
)

func (a *BaseActor) OnConnected(connection miface.IConnection) {
	msg := NewSocketMessage(SocketActConnected, connection, nil)
	a.PutMessage(msg)
}

func (a *BaseActor) OnDisconnect(connection miface.IConnection) {
	msg := NewSocketMessage(SocketActDisconnect, connection, nil)
	a.PutMessage(msg)
}

func (a *BaseActor) OnProcess(connection miface.IConnection, pack miface.IPackage) {
	msg := NewSocketMessage(SocketActData, connection, pack.GetData())
	a.PutMessage(msg)
}

func (a *BaseActor) NetListen(options ...core.Option) miface.IServer {
	// 网络消息 回调到 actor基类
	options = append(options, core.WithRouter(a))
	s := mznet.NewServer(options...)
	if err := s.Run(); err != nil {
		log.Printf("server start fail")
		return nil
	}
	return s
}

func (a *BaseActor) NetConnect(options ...core.Option) miface.IClient {
	options = append(options, core.WithRouter(a))
	c := mznet.NewClient(options...)
	if err := c.Connect(); err != nil {
		log.Printf("client connect  fail")
		return nil
	}
	return c
}

func (a *BaseActor) Dispatch(msg iface.IActorMessage) {
	protocol, ok := a.protocols[1]
	if !ok {
		log.Printf("not registry protocol type:%d", msg.GetType())
		return
	}

	if protocol.parse == nil {
		log.Printf("protocol not parse type:%d", msg.GetType())
		return
	}

	if protocol.dispatch == nil {
		log.Printf("protocol not handler type:%d", msg.GetType())
		return
	}

	rets, err := protocol.parse.UnMarshal(msg.GetData())
	if err != nil {
		log.Printf("protocol parse fail type:%d", msg.GetType())
		return
	}
	protocol.dispatch(msg, rets...)
}
