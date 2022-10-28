/**
 * @Author: dingQingHui
 * @Description:
 * @File: iclient
 * @Version: 1.0.0
 * @Date: 2022/7/7 18:24
 */

package miface

import "net"

type (
	IConnect interface {
		Connect() error
	}

	IAccept interface {
		Accept() error
	}

	ICloser interface {
		Close() error
	}

	IConnection interface {
		GetId() int64
		GetType() TypeConnection

		Send(message *Message) bool

		GetLocalAddr() net.Addr
		GetRemoteAddr() net.Addr
		IsClose() bool

		ICloser
	}

	IClient interface {
		IConnect
		IConnection
	}

	IServer interface {
		IAccept
		ICloser
	}

	MessageHead struct {
		Len uint16
	}
	Message struct {
		MessageHead
		Body []byte
	}
)

func NewMessage(Len uint16, Body []byte) *Message {
	return &Message{
		MessageHead: MessageHead{
			Len: Len,
		},
		Body: Body,
	}
}
