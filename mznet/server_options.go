/**
 * @Author: dingQingHui
 * @Description:
 * @File: server_options
 * @Version: 1.0.0
 * @Date: 2022/10/28 10:54
 */

package mznet

type ServerOptionFun func(o *serverOptions)

type serverOptions struct {
	listenAddress string
	network       string
	eventListener IEventListener
}

var defaultServerOption = &serverOptions{
	eventListener: DefaultEventListener{},
	network:       "tcp",
	listenAddress: "127.0.0.1:9843",
}

func WithServerAddress(address string) ServerOptionFun {
	return func(o *serverOptions) {
		o.listenAddress = address
	}
}
func WithServerNetwork(network string) ServerOptionFun {
	return func(o *serverOptions) {
		o.network = network
	}
}

func WithServerEventListener(eventListener IEventListener) ServerOptionFun {
	return func(o *serverOptions) {
		o.eventListener = eventListener
	}
}
