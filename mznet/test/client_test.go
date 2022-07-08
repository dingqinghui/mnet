/**
 * @Author: dingQingHui
 * @Description:
 * @File: client_test
 * @Version: 1.0.0
 * @Date: 2022/7/7 18:03
 */

package test

import (
	"mz/mznet"
	"mz/mznet/codec"
	core2 "mz/mznet/core"
	miface2 "mz/mznet/miface"
	"testing"
)

type defaultRouter1 struct {
}

func (r *defaultRouter1) OnConnected(connection miface2.IConnection) {
	msg := core2.NewMessage(512, make([]byte, 512))
	connection.Send(msg)
}
func (r *defaultRouter1) OnDisconnect(connection miface2.IConnection) {

}
func (r *defaultRouter1) OnProcess(connection miface2.IConnection, message miface2.IMessage) {
	connection.Send(message)
	//log.Printf("recv msg len:%d data:%s", message.GetDataLen(), message.GetData())
}

func TestClient(t *testing.T) {

	c := mznet.NewClient(core2.WithAddress("192.168.1.170:2100"),
		core2.WithNetwork("tcp"),
		core2.WithRouter(&defaultRouter1{}),
		core2.WithTcpCodec(codec.NewCommonCodec()))

	c.Connect()

	select {}
}
