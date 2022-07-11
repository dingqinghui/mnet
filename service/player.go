/**
 * @Author: dingQingHui
 * @Description:
 * @File: player
 * @Version: 1.0.0
 * @Date: 2022/7/11 17:49
 */

package service

import (
	"mz/actor"
	"mz/handler"
	"mz/iface"
	"mz/mznet/miface"
	"mz/parser"
)

var (
	netParser  iface.IParse
	netHandler iface.IHandler
)

func init() {
	netHandler = handler.NewHandler()
	netParser = parser.NewJsonParser()
}

func EchoHandle(server miface.IServer, connection miface.IConnection, msg iface.IMessage) {
	println(msg.Msg().(*MessageEcho).Msg)
}

func registryMessage(msgId uint32, message interface{}, handlerFun iface.HandlerFun) {
	netHandler.SetHandler(msgId, handlerFun)
	netParser.Register(msgId, message)
}

type (
	MessageEcho struct {
		Msg string
	}

	Player struct {
		iface.IActor
	}
)

func NewPlayerActor() iface.IActor {
	return &Player{
		IActor: actor.New(),
	}
}

func (p *Player) Init(args ...interface{}) {
	registryMessage(1, &MessageEcho{}, EchoHandle)
	p.IActor.RegistryProtocol(iface.ActorMessageNetWork, netParser, netHandler)
}

func (p *Player) Destroy() {

}
