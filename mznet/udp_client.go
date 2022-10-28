/**
 * @Author: dingQingHui
 * @Description:
 * @File: udpClient
 * @Version: 1.0.0
 * @Date: 2022/10/28 16:40
 */

package mznet

import (
	"net"
)

type udpClient struct {
	opts *clientOptions
	IConnection
}

func newUdpClient(opts ...ClientOptionFun) IClient {
	s := &udpClient{
		opts: defaultClientOption,
	}
	for _, opt := range opts {
		opt(s.opts)
	}
	return s
}

func (c *udpClient) Connect() error {
	raddr, err := net.ResolveUDPAddr("udp", c.opts.address)
	if err != nil {
		return err
	}

	conn, err := net.DialUDP("udp", nil, raddr)
	if err != nil {
		return err
	}

	c.IConnection = newUdpConnection(conn, ConnectConnection, c.opts.address, c.opts.network)
	c.opts.eventListener.OnConnected(c.IConnection)

	go c.read()
	return nil
}

func (c *udpClient) read() {
	for true {
		b := make([]byte, 1024)
		udpCon := c.GetCon().(*net.UDPConn)
		n, _, err := udpCon.ReadFromUDP(b)
		if err != nil {
			c.opts.eventListener.OnError(c.IConnection, err)
			return
		}
		uc := c.IConnection.(*udpConnection)
		uc.recUdpStream(b[:n])
	}
	return
}
func (c *udpClient) Close() error {
	return c.IConnection.Close()
}
