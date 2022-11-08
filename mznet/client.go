/**
 * @Author: dingQingHui
 * @Description:
 * @File: client
 * @Version: 1.0.0
 * @Date: 2022/10/28 16:56
 */

package mznet

type (
	IClient interface {
		ICloser
		Connect() error
	}
)

func NewClient(config *ClientConfig) IClient {
	switch config.Network {
	case "tcp":
		return newTcpClient(config)
	case "udp":
		return newUdpClient(config)
	}
	return nil
}
