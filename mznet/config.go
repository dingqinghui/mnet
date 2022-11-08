/**
 * @Author: dingQingHui
 * @Description:
 * @File: server_options
 * @Version: 1.0.0
 * @Date: 2022/10/28 10:54
 */

package mznet

type ClientConfig struct {
	Address       string
	Network       string
	EventListener IEventListener
}

type ServerConfig struct {
	ListenAddress string
	Network       string
	EventListener IEventListener
}

var defaultClientConfig = &ClientConfig{
	EventListener: DefaultEventListener{},
	Network:       "tcp",
	Address:       "127.0.0.1:9843",
}

var defaultServerConfig = &ServerConfig{
	EventListener: DefaultEventListener{},
	Network:       "tcp",
	ListenAddress: "127.0.0.1:9843",
}
