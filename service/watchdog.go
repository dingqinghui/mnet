/**
 * @Author: dingQingHui
 * @Description:
 * @File: player
 * @Version: 1.0.0
 * @Date: 2022/7/11 17:49
 */

package service

import (
	"github.com/dingqinghui/mz/actor"
	"github.com/dingqinghui/mz/actor/iface"
	"github.com/dingqinghui/mz/actor/parser"
	iface2 "github.com/dingqinghui/mz/actorNew"
	"github.com/dingqinghui/mz/handler"
	"log"

	"github.com/dingqinghui/mz/message"
	"github.com/dingqinghui/mz/mznet/codec"
	"github.com/dingqinghui/mz/mznet/core"
	"github.com/dingqinghui/mz/mznet/miface"
)

var (
	netParser  iface.IParse
	netHandler handler.IHandler
)

func init() {
	netHandler = handler.NewHandler()
	netParser = parser.NewJsonParser()

	registryMessage(1, &MessageEcho{}, EchoHandle)

	actor.RegistryActor("watchDog", NewWatchdog)
}

func EchoHandle(connection miface.IConnection, msg message.IMessage) {
	//println(msg.Msg().(*MessageEcho).Msg)

	var msgId uint32 = 1
	data, err := netParser.Marshal(msgId, &MessageEcho{Msg: make([]byte, 512, 512)})
	_ = err
	connection.Send(core.NewPackage(uint32(len(data)), data))

}

func registryMessage(msgId uint32, message interface{}, handlerFun handler.HandlerFun) {
	netHandler.SetHandler(msgId, handlerFun)
	netParser.Register(msgId, message)
}

type (
	MessageEcho struct {
		Msg []byte
	}

	Watchdog struct {
		*actor.BaseActor
	}
)

func NewWatchdog() iface2.IActor {
	return &Watchdog{
		BaseActor: actor.NewBase(),
	}
}

func (p *Watchdog) Init(args ...interface{}) {

	p.NetListen(core.WithAddress("192.168.1.170:2100"),
		core.WithNetwork("tcp"),
		core.WithTcpCodec(codec.NewCommonCodec()),
	)

	p.RegistryProtocol(iface.ActorMessageSocket, netParser,
		func(msg iface.IActorMessage, args ...interface{}) {
			socketMsg := msg.(*actor.SocketMessage)

			switch socketMsg.Act {
			case actor.SocketActConnected:
				log.Print("SocketActConnected")
			case actor.SocketActData:
				//log.Print("SocketActData")
				msgId := args[0].(uint32)
				f, _ := netHandler.GetHandler(msgId)
				if f != nil {
					f(socketMsg.Connection, message.NewMessage(msgId, args[1]))
				}
			case actor.SocketActDisconnect:
				log.Print("SocketActDisconnect")
			}
		})
}

func (p *Watchdog) Destroy() {

}
