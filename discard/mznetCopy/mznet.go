/**
 * @Author: dingQingHui
 * @Description:
 * @File: mznet
 * @Version: 1.0.0
 * @Date: 2022/7/8 16:55
 */

package mznet

import (
	"github.com/dingqinghui/mz/mznet/core"
	"github.com/dingqinghui/mz/mznet/miface"
	"github.com/dingqinghui/mz/mznet/tcp"
	"github.com/dingqinghui/mz/mznet/udp"
	"log"
)

var (
	tcpNetwork = "tcp"
	udpNetwork = "udp"
)

func NewServer(options ...core.Option) miface.IServer {
	opts := core.DefaultOptions
	for _, opt := range options {
		opt(&opts)
	}
	switch opts.Network {
	case tcpNetwork:
		return tcp.NewServer(opts)
	case udpNetwork:
		return udp.NewUdpServer(opts)
	default:
		log.Println("invalid network must is tcp or udp")
	}
	return nil
}

func NewClient(options ...core.Option) miface.IClient {
	opts := core.DefaultOptions
	for _, opt := range options {
		opt(&opts)
	}
	switch opts.Network {
	case tcpNetwork:
		return tcp.NewClient(opts)
	case udpNetwork:
		return udp.NewUdpClient(opts)
	default:
		log.Println("invalid network must is tcp or udp")
	}
	return nil
}
