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
	IConnection
	config *ClientConfig
}

func newUdpClient(config *ClientConfig) IClient {
	s := &udpClient{
		config: config,
	}

	return s
}

func (c *udpClient) Connect() error {
	radar, err := net.ResolveUDPAddr("udp", c.config.Address)
	if err != nil {
		return err
	}

	conn, err := net.DialUDP("udp", nil, radar)
	if err != nil {
		return err
	}

	c.IConnection = newUdpConnection(conn, ConnectConnection, c.config.Address, c.config.Network, c.config.EventListener)
	c.config.EventListener.OnConnected(c.IConnection)

	go c.read()
	return nil
}

func (c *udpClient) read() {
	for true {
		b := make([]byte, 1024)
		udpCon := c.GetCon().(*net.UDPConn)
		n, _, err := udpCon.ReadFromUDP(b)
		if err != nil {
			c.config.EventListener.OnError(c.IConnection, err)
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
