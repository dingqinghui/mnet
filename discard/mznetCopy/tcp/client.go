/**
 * @Author: dingQingHui
 * @Description:
 * @File: client
 * @Version: 1.0.0
 * @Date: 2022/7/7 18:13
 */

package tcp

import (
	"github.com/dingqinghui/mz/mznet/core"
	"github.com/dingqinghui/mz/mznet/miface"
	"net"
)

type client struct {
	options core.Options
	miface.IConnection
}

func NewClient(options core.Options) miface.IClient {
	s := &client{
		options: options,
	}
	return s
}
func (c *client) Connect() error {
	con, err := net.Dial(c.options.Network, c.options.Address)
	if err != nil {
		return err
	}

	c.IConnection = NewConnection(c.options.Network, con, miface.TypeConnectionConnect, c.options)
	return nil
}
