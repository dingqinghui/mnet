/**
 * @Author: dingQingHui
 * @Description:
 * @File: client
 * @Version: 1.0.0
 * @Date: 2022/7/12 15:14
 */

package service

import (
	"github.com/dingqinghui/mz/actor"
	"github.com/dingqinghui/mz/actor/iface"
	"github.com/dingqinghui/mz/message"
	"github.com/dingqinghui/mz/mznet/codec"
	"github.com/dingqinghui/mz/mznet/core"
	"log"
)

type (
	Client struct {
		*actor.Actor
	}
)

func NewClient() iface.IActor {
	return &Client{
		Actor: actor.New(),
	}
}

func (p *Client) Init(args ...interface{}) {
	p.NetConnect(core.WithAddress("192.168.1.170:2100"),
		core.WithNetwork("tcp"),
		core.WithTcpCodec(codec.NewCommonCodec()),
	)

	p.RegistryProtocol(iface.ActorMessageSocket, netParser,
		func(msg iface.IActorMessage, args ...interface{}) {
			socketMsg := msg.(*actor.SocketMessage)

			switch socketMsg.Act {
			case actor.SocketActConnected:
				log.Print("SocketActConnected")
				var msgId uint32 = 1

				data, err := netParser.Marshal(msgId, &MessageEcho{Msg: make([]byte, 512, 512)})
				_ = err
				socketMsg.Connection.Send(core.NewPackage(uint32(len(data)), data))
			case actor.SocketActData:
				msgId := args[0].(uint32)
				//log.Print("SocketActData")
				f, _ := netHandler.GetHandler(msgId)
				if f != nil {
					f(socketMsg.Connection, message.NewMessage(msgId, args[1]))
				}
			case actor.SocketActDisconnect:
				log.Print("SocketActDisconnect")
			}
		})

}

func (p *Client) Destroy() {

}
