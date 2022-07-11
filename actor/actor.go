/**
 * @Author: dingQingHui
 * @Description:
 * @File: actor
 * @Version: 1.0.0
 * @Date: 2022/7/11 16:49
 */

package actor

import (
	"log"
	"mz/iface"
	"mz/mznet/miface"
)

type (
	Protocol struct {
		msgType iface.ActorMessageType
		parse   iface.IParse
		handler iface.IHandler
	}

	Message struct {
		msgType    iface.ActorMessageType
		server     miface.IServer
		connection miface.IConnection
		pack       miface.IPackage
	}

	Actor struct {
		id        uint64
		revChan   chan *iface.Message
		protocols map[iface.ActorMessageType]*Protocol
	}
)

func NewMessage(msgType iface.ActorMessageType, server miface.IServer, connection miface.IConnection, pack miface.IPackage) *Message {
	return &Message{
		msgType:    msgType,
		server:     server,
		connection: connection,
		pack:       pack,
	}
}

func New() iface.IActor {
	return &Actor{
		revChan:   make(chan *iface.Message, 0),
		protocols: make(map[iface.ActorMessageType]*Protocol),
	}
}
func (a *Actor) Init(args ...interface{}) {

}
func (a *Actor) GetId() uint64 {
	return a.id
}

func (a *Actor) RegistryProtocol(msgType iface.ActorMessageType, parse iface.IParse, handler iface.IHandler) {
	a.protocols[msgType] = &Protocol{
		parse:   parse,
		handler: handler,
	}
}

func (a *Actor) PutMessage(msg *iface.Message) {
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
func (a *Actor) Dispatch(msg *iface.Message) {
	protocol, ok := a.protocols[msg.MsgType]
	if !ok {
		log.Printf("not registry protocol type:%d", msg.MsgType)
		return
	}

	if protocol.parse == nil {
		log.Printf("protocol not parse type:%d", msg.MsgType)
		return
	}

	if protocol.handler == nil {
		log.Printf("protocol not handler type:%d", msg.MsgType)
		return
	}

	useMsg, err := protocol.parse.UnMarshal(msg.Pack.GetData())
	if err != nil {
		log.Printf("protocol parse fail type:%d", msg.MsgType)
		return
	}

	f, err := protocol.handler.GetHandler(useMsg.MsgId())
	if err != nil {
		log.Printf("protocol handler is nil msgId:%d", useMsg.MsgId())
		return
	}

	f(msg.Server, msg.Connection, useMsg)
}

func (a *Actor) Destroy() {

}
