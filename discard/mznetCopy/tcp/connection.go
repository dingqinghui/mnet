/**
 * @Author: dingQingHui
 * @Description:
 * @File: connection
 * @Version: 1.0.0
 * @Date: 2022/7/7 15:19
 */

package tcp

import (
	"github.com/dingqinghui/mz/mznet/core"
	"github.com/dingqinghui/mz/mznet/miface"
	"net"
)

type connection struct {
	*core.Connection
	con *net.TCPConn
}

func NewConnection(_ string, con net.Conn, conType miface.TypeConnection, options core.Options) miface.IConnection {
	c := &connection{
		Connection: core.NewConnection(conType, options, con.LocalAddr(), con.RemoteAddr()),
		con:        con.(*net.TCPConn),
	}
	if err := c.start(); err != nil {
		return nil
	}
	return c
}

func (t *connection) start() error {
	//if err := t.con.SetNoDelay(false); err != nil {
	//	return err
	//}
	go t.waitExit()
	go t.read()
	go t.write()
	t.Options.Router.OnConnected(t)
	return nil
}

func (t *connection) waitExit() {
	select {
	case <-t.Options.ParentCtx.Done(): // 父节点退出
		t.destroy()
	case <-t.Ctx.Done(): // 本节点退出
		t.destroy()
	}
}

func (t *connection) read() {
	defer t.Close()
	for true {
		msg := core.NewPackage(0, nil)
		err := t.Options.Codec.Unpack(t.con, msg)
		if err != nil {
			return
		}
		t.Options.Router.OnProcess(t, msg)
	}
}

func (t *connection) write() {
	defer t.Close()
	for true {
		select {
		case msg, ok := <-t.WriteChan:
			if !ok {
				break
			}
			err := t.Options.Codec.Pack(t.con, msg)
			if err != nil {
				break
			}
		}
	}
}

func (t *connection) destroy() {
	if !t.SetClose() {
		return
	}
	// 管理连接
	_ = t.con.Close()
	// 关闭管道
	close(t.WriteChan)
}
