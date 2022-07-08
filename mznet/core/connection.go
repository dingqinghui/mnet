/**
 * @Author: dingQingHui
 * @Description:
 * @File: connection
 * @Version: 1.0.0
 * @Date: 2022/7/8 17:37
 */

package core

import (
	"context"
	"mz/mznet/miface"
	"net"
	"sync/atomic"
)

type Connection struct {
	id                    int64
	WriteChan             chan miface.IMessage
	Stop                  atomic.Value
	Options               Options
	ConType               miface.TypeConnection
	localAddr, remoteAddr net.Addr

	Ctx    context.Context // 通知子连接退出
	Cancel context.CancelFunc
}

func NewConnection(conType miface.TypeConnection, options Options, localAddr, remoteAddr net.Addr) *Connection {
	c := &Connection{
		id:         GenId(),
		Options:    options,
		WriteChan:  make(chan miface.IMessage, 0),
		ConType:    conType,
		localAddr:  localAddr,
		remoteAddr: remoteAddr,
	}
	c.init()
	return c
}

func (c *Connection) init() {
	c.Stop.Store(false)
	c.Ctx, c.Cancel = context.WithCancel(c.Options.ParentCtx)
}

func (c *Connection) GetId() int64 {
	return c.id
}

func (c *Connection) GetLocalAddr() net.Addr {
	return c.localAddr
}

func (c *Connection) GetRemoteAddr() net.Addr {
	return c.remoteAddr
}

func (c *Connection) Send(message miface.IMessage) bool {
	if c.WriteChan == nil {
		return false
	}
	if c.Stop.CompareAndSwap(true, true) {
		return false
	}
	c.WriteChan <- message
	return true
}

func (c *Connection) GetType() miface.TypeConnection {
	return c.ConType
}

func (c *Connection) IsClose() bool {
	return c.Stop.CompareAndSwap(true, true)
}
func (c *Connection) SetClose() bool {
	return c.Stop.CompareAndSwap(false, true)
}

//
// Close
// @Description: 主动关闭
// @receiver s
//
func (c *Connection) Close() bool {
	if c.SetClose() {
		c.Cancel()
		return true
	}
	return false
}
