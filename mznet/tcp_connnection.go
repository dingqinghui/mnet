/**
 * @Author: dingQingHui
 * @Description:
 * @File: tcp_connnection
 * @Version: 1.0.0
 * @Date: 2022/10/28 13:59
 */

package mznet

import "net"

type tcpConnection struct {
	IConnection
}

func newTcpConnection(con net.Conn, cType ConnectionType, eventListener IEventListener) *tcpConnection {
	t := &tcpConnection{
		IConnection: newConnection(con, cType, eventListener),
	}
	go t.read()
	go t.write()
	return t
}

func (t *tcpConnection) read() {
	for !t.IsClose() {
		msg, err := t.GetCodec().UnPack(t.GetCon())
		if err != nil {
			if !t.GetEventListener().OnError(t, err) {
				return
			}
			continue
		}
		if !t.GetEventListener().OnProcess(t, msg) {
			return
		}
	}
}

func (t *tcpConnection) write() {
	for !t.IsClose() {
		select {
		case msg, _ := <-t.GetWriteChan():
			if err := t.GetCodec().Pack(t.GetCon(), msg); err != nil {
				if !t.GetEventListener().OnError(t, err) {
					return
				}
			}
		}
	}
}
