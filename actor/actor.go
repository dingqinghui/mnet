/**
 * @Author: dingQingHui
 * @Description:
 * @File: actor
 * @Version: 1.0.0
 * @Date: 2022/7/11 16:49
 */

package actor

import (
	"github.com/dingqinghui/mz/actor/iface"
	"github.com/dingqinghui/mz/mznet"
	"github.com/dingqinghui/mz/mznet/core"
	"github.com/dingqinghui/mz/mznet/miface"
	"log"
)

type (
	Protocol struct {
		msgType  iface.ActorMessageType
		parse    iface.IParse
		dispatch iface.DispatchFun
	}
	Actor struct {
		id        uint64
		revChan   chan iface.IActorMessage
		protocols map[iface.ActorMessageType]*Protocol
	}
)

func New() *Actor {
	return &Actor{
		revChan:   make(chan iface.IActorMessage, 1),
		protocols: make(map[iface.ActorMessageType]*Protocol),
	}
}
func (a *Actor) Init(_ ...interface{}) {

}
func (a *Actor) GetId() uint64 {
	return a.id
}

func (a *Actor) RegistryProtocol(msgType iface.ActorMessageType, parse iface.IParse, dispatch iface.DispatchFun) {
	a.protocols[msgType] = &Protocol{
		parse:    parse,
		dispatch: dispatch,
	}
}

func (a *Actor) PutMessage(msg iface.IActorMessage) {
	a.revChan <- msg
}

func (a *Actor) Run() {
	go func() {
		for true {
			select {
			case msg := <-a.revChan:
				a.Dispatch(msg)
			}
		}
	}()
}
func (a *Actor) Dispatch(msg iface.IActorMessage) {
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

func (a *Actor) Destroy() {

}

func (a *Actor) OnConnected(connection miface.IConnection) {
	msg := NewSocketMessage(SocketActConnected, connection, nil)
	a.PutMessage(msg)
}

func (a *Actor) OnDisconnect(connection miface.IConnection) {
	msg := NewSocketMessage(SocketActDisconnect, connection, nil)
	a.PutMessage(msg)
}

func (a *Actor) OnProcess(connection miface.IConnection, pack miface.IPackage) {
	msg := NewSocketMessage(SocketActData, connection, pack.GetData())
	a.PutMessage(msg)
}

func (a *Actor) NetListen(options ...core.Option) miface.IServer {
	// 网络消息 回调到 actor基类
	options = append(options, core.WithRouter(a))
	s := mznet.NewServer(options...)
	if err := s.Run(); err != nil {
		log.Printf("server start fail")
		return nil
	}
	return s
}

func (a *Actor) NetConnect(options ...core.Option) miface.IClient {
	options = append(options, core.WithRouter(a))
	c := mznet.NewClient(options...)
	if err := c.Connect(); err != nil {
		log.Printf("client connect  fail")
		return nil
	}
	return c
}
