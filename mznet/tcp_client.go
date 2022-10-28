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
	opts *clientOptions
	IConnection
}

func newTcpClient(opts ...ClientOptionFun) IClient {
	s := &tpcClient{
		opts: defaultClientOption,
	}
	for _, opt := range opts {
		opt(s.opts)
	}
	return s
}

func (c *tpcClient) Connect() error {
	con, err := net.Dial(c.opts.network, c.opts.address)
	if err != nil {
		return err
	}
	c.IConnection = newTcpConnection(con, ConnectConnection)
	c.opts.eventListener.OnConnected(c.IConnection)
	return nil
}

func (c *tpcClient) Close() error {
	return c.IConnection.Close()
}
