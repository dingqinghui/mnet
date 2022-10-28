/**
 * @Author: dingQingHui
 * @Description:
 * @File: connection
 * @Version: 1.0.0
 * @Date: 2022/10/28 14:03
 */

package mznet

import (
	"net"
	"sync/atomic"
)

type ConnectionType int

var (
	AcceptConnection  ConnectionType = 1
	ConnectConnection ConnectionType = 2
)

var connectionId int32

func genId() int32 {
	return atomic.AddInt32(&connectionId, 1)
}

type IConnection interface {
	GetId() int32
	GetType() ConnectionType
	GetCodec() ICodec
	SetCodec(codec ICodec)
	SetEventListener(eventListener IEventListener)
	GetEventListener() IEventListener
	GetWriteChan() chan interface{}
	GetCon() net.Conn
	Send(interface{}) error
	ICloser
}

type connection struct {
	id            int32
	con           net.Conn
	cType         ConnectionType
	codec         ICodec
	eventListener IEventListener
	writeChan     chan interface{}
	ICloser
}

func newConnection(con net.Conn, cType ConnectionType) IConnection {
	return &connection{
		id:        genId(),
		con:       con,
		cType:     cType,
		writeChan: make(chan interface{}, 1024),
		ICloser:   defaultCloser,
	}
}

func (t *connection) GetType() ConnectionType {
	return t.cType
}

func (t *connection) GetId() int32 {
	return t.id
}

func (t *connection) GetCodec() ICodec {
	if t.codec == nil {
		return DefaultCodec
	}
	return t.codec
}

func (t *connection) SetCodec(codec ICodec) {
	t.codec = codec
}

func (t *connection) SetEventListener(eventListener IEventListener) {
	t.eventListener = eventListener
}

func (t *connection) GetEventListener() IEventListener {
	return t.eventListener
}
func (t *connection) GetWriteChan() chan interface{} {
	return t.writeChan
}
func (t *connection) GetCon() net.Conn {
	return t.con
}
func (t *connection) Send(data interface{}) error {
	t.writeChan <- data
	return nil
}

func (t *connection) Close() error {
	if err := t.ICloser.Close(); err != nil {
		return err
	}
	return t.con.Close()
}
