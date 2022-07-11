/**
 * @Author: dingQingHui
 * @Description:
 * @File: client_udp
 * @Version: 1.0.0
 * @Date: 2022/7/8 15:55
 */

package udp

import (
	"errors"
	"github.com/dingqinghui/mz/mznet/core"
	"github.com/dingqinghui/mz/mznet/miface"
	"net"
)

type client struct {
	options core.Options
}

func NewUdpClient(options core.Options) miface.IClient {
	s := &client{
		options: options,
	}
	return s
}
func (u *client) Connect() error {
	con, err := net.Dial("udp", u.options.Address)
	if err != nil {
		return err
	}

	udpCon, ok := con.(*net.UDPConn)
	if !ok {
		return errors.New("change type UDPConn fail")
	}

	c := newConnection(u.options.Network, udpCon, miface.TypeConnectionConnect, u.options)

	for true {
		b := make([]byte, 1024)
		n, _, err := udpCon.ReadFrom(b)
		if err != nil {
			return err
		}
		udpConnection, ok := c.(*connection)
		if !ok {
			return errors.New("change type connection fail")
		}
		udpConnection.RevMsg(core.NewPackage(uint32(n), b[:n]))
	}
	return nil
}
