/**
 * @Author: dingQingHui
 * @Description:
 * @File: main
 * @Version: 1.0.0
 * @Date: 2022/7/7 14:56
 */

package main

import (
	"log"
	"mz/mznet"
	"mz/mznet/codec"
	"mz/mznet/core"
	miface2 "mz/mznet/miface"
	"time"
)

type defaultRouter struct {
}

func (r *defaultRouter) OnConnected(connection miface2.IConnection) {
	log.Printf("OnConnected")
}
func (r *defaultRouter) OnDisconnect(connection miface2.IConnection) {

}
func (r *defaultRouter) OnProcess(connection miface2.IConnection, message miface2.IMessage) {
	//log.Printf("recv msg len:%d data:%s", message.GetDataLen(), message.GetData())
	connection.Send(message)
}

func main() {
	s := mznet.NewServer(core.WithAddress("192.168.1.170:2100"),
		core.WithNetwork("tcp"),
		core.WithRouter(&defaultRouter{}),
		core.WithTcpCodec(codec.NewCommonCodec()))

	s.Run()

	for true {
		time.Sleep(time.Second)
	}
}
