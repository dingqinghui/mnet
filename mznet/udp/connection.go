/**
 * @Author: dingQingHui
 * @Description:
 * @File: connection_udp
 * @Version: 1.0.0
 * @Date: 2022/7/8 14:34
 */

package udp

import (
	"github.com/dingqinghui/mz/mznet/core"
	"github.com/dingqinghui/mz/mznet/miface"
	"net"
)

type connection struct {
	*core.Connection
	con      *net.UDPConn
	readChan chan miface.IPackage
}

func newConnection(_ string, con net.Conn, conType miface.TypeConnection, options core.Options) miface.IConnection {
	c := &connection{
		con:        con.(*net.UDPConn),
		Connection: core.NewConnection(conType, options, con.LocalAddr(), con.RemoteAddr()),
		readChan:   make(chan miface.IPackage, 0),
	}
	c.start()
	return c
}

func (u *connection) start() {
	go u.waitExit()
	go u.write()
	go u.read()
	u.Options.Router.OnConnected(u)
}

func (u *connection) waitExit() {
	select {
	case <-u.Options.ParentCtx.Done(): // 父节点退出
		u.destroy()
	case <-u.Ctx.Done(): // 本节点退出
		u.destroy()
	}
}

func (u *connection) write() {
	defer u.Close()
	for true {
		select {
		case msg, ok := <-u.WriteChan:
			if !ok {
				return
			}
			switch u.GetType() {
			case miface.TypeConnectionAccept:
				if _, err := u.con.WriteTo(msg.GetData(), u.Options.UdpAddr); err != nil {
					return
				}
			case miface.TypeConnectionConnect:
				if _, err := u.con.Write(msg.GetData()); err != nil {
					return
				}
			}
		}
	}
}

func (u *connection) RevMsg(message miface.IPackage) bool {
	if u.readChan == nil {
		return false
	}
	u.readChan <- message
	return true
}

func (u *connection) read() {
	defer u.Close()
	for true {
		select {
		case msg, ok := <-u.readChan:
			if !ok {
				return
			}
			if msg.GetDataLen() == 0 {
				return
			}
			u.Options.Router.OnProcess(u, msg)
		}
	}
}

func (u *connection) destroy() {
	if !u.SetClose() {
		return
	}
	close(u.readChan)
	close(u.WriteChan)
}
