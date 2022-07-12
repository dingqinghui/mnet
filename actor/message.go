/**
 * @Author: dingQingHui
 * @Description:
 * @File: message
 * @Version: 1.0.0
 * @Date: 2022/7/12 14:40
 */

package actor

import (
	"github.com/dingqinghui/mz/actor/iface"
	"github.com/dingqinghui/mz/mznet/miface"
)

var (
	SocketActConnected  SocketActType = 1
	SocketActData       SocketActType = 2
	SocketActDisconnect SocketActType = 3
)

type (
	Message struct {
		MsgType iface.ActorMessageType
		Data    []byte
	}

	SocketActType int
	SocketMessage struct {
		iface.IActorMessage
		Act        SocketActType
		Connection miface.IConnection
		Pack       miface.IPackage
	}
)

func NewMessage(msgType iface.ActorMessageType, Data []byte) *Message {
	return &Message{
		MsgType: msgType,
		Data:    Data,
	}
}

func (m *Message) GetType() iface.ActorMessageType {
	return m.MsgType
}

func (m *Message) GetData() []byte {
	return m.Data
}

func NewSocketMessage(act SocketActType, connection miface.IConnection, data []byte) iface.IActorMessage {
	return &SocketMessage{
		IActorMessage: NewMessage(iface.ActorMessageSocket, data),
		Act:           act,
		Connection:    connection,
	}
}
