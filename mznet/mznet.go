/**
 * @Author: dingQingHui
 * @Description:
 * @File: mznet
 * @Version: 1.0.0
 * @Date: 2022/7/8 16:55
 */

package mznet

import (
	"log"
	"mz/mznet/core"
	"mz/mznet/miface"
	"mz/mznet/tcp"
	"mz/mznet/udp"
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
