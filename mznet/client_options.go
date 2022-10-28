/**
 * @Author: dingQingHui
 * @Description:
 * @File: server_options
 * @Version: 1.0.0
 * @Date: 2022/10/28 10:54
 */

package mznet

type ClientOptionFun func(o *clientOptions)

type clientOptions struct {
	address       string
	network       string
	eventListener IEventListener
}

var defaultClientOption = &clientOptions{
	eventListener: DefaultEventListener{},
	network:       "tcp",
	address:       "127.0.0.1:9843",
}

func WithClientAddress(address string) ClientOptionFun {
	return func(o *clientOptions) {
		o.address = address
	}
}
func WithClientNetwork(network string) ClientOptionFun {
	return func(o *clientOptions) {
		o.network = network
	}
}

func WithClientEventListener(eventListener IEventListener) ClientOptionFun {
	return func(o *clientOptions) {
		o.eventListener = eventListener
	}
}
