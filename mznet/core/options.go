/**
 * @Author: dingQingHui
 * @Description:
 * @File: Options
 * @Version: 1.0.0
 * @Date: 2022/7/8 10:23
 */

package core

import (
	"context"
	"github.com/dingqinghui/mz/mznet/miface"
	"net"
)

type (
	Option func(o *Options)

	Options struct {
		ParentCtx        context.Context
		Network, Address string
		Router           miface.IRouter
		Codec            miface.ICodec
		UdpAddr          net.Addr
	}
)

var (
	DefaultOptions = Options{
		ParentCtx: context.Background(),
		Address:   "127.0.0.1:5555",
		Router:    NewDefaultRouter(),
	}
)

func WithContext(context context.Context) Option {
	return func(opts *Options) {
		opts.ParentCtx = context
	}
}

func WithNetwork(network string) Option {
	return func(opts *Options) {
		opts.Network = network
	}
}

func WithAddress(address string) Option {
	return func(opts *Options) {
		opts.Address = address
	}
}

func WithRouter(router miface.IRouter) Option {
	return func(opts *Options) {
		opts.Router = router
	}
}

func WithTcpCodec(codec miface.ICodec) Option {
	return func(opts *Options) {
		opts.Codec = codec
	}
}

func WithUdpRemoteAddr(addr net.Addr) Option {
	return func(opts *Options) {
		opts.UdpAddr = addr
	}
}
