/**
 * @Author: dingQingHui
 * @Description:
 * @File: IActor
 * @Version: 1.0.0
 * @Date: 2022/7/11 16:46
 */

package iface

import "github.com/dingqinghui/mz/mznet/miface"

var (
	ActorMessageNetWork ActorMessageType = 1
	ActorMessageRpc     ActorMessageType = 2
)

type (
	ActorMessageType int

	IActor interface {
		GetId() uint64
		Init(args ...interface{})
		Run()
		PutMessage(msg *Message)
		RegistryProtocol(msgType ActorMessageType, parse IParse, handler IHandler)
	}

	Message struct {
		MsgType    ActorMessageType
		Server     miface.IServer
		Connection miface.IConnection
		Pack       miface.IPackage
	}
)

func NewMessage(msgType ActorMessageType, server miface.IServer, connection miface.IConnection, pack miface.IPackage) *Message {
	return &Message{
		MsgType:    msgType,
		Server:     server,
		Connection: connection,
		Pack:       pack,
	}
}
