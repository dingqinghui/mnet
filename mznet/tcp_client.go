/**
 * @Author: dingQingHui
 * @Description:
 * @File: tcp_client
 * @Version: 1.0.0
 * @Date: 2022/10/28 16:26
 */

package mznet

import (
	"net"
)

type tpcClient struct {
	config *ClientConfig
	IConnection
}

func newTcpClient(config *ClientConfig) IClient {
	s := &tpcClient{
		config: config,
	}
	return s
}

func (c *tpcClient) Connect() error {
	con, err := net.Dial(c.config.Network, c.config.Address)
	if err != nil {
		return err
	}
	c.IConnection = newTcpConnection(con, ConnectConnection, c.config.EventListener)
	c.config.EventListener.OnConnected(c.IConnection)
	return nil
}

func (c *tpcClient) Close() error {
	return c.IConnection.Close()
}
