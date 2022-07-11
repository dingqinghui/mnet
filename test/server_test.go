/**
 * @Author: dingQingHui
 * @Description:
 * @File: server_test
 * @Version: 1.0.0
 * @Date: 2022/7/7 18:04
 */

package test

import (
	"log"
	"mz/iface"
	"mz/message"
	"mz/mznet"
	"mz/mznet/codec"
	"mz/mznet/core"
	"mz/mznet/miface"
	"mz/parser"
	"mz/service"
	"testing"
	"time"
)

var pActor iface.IActor

type defaultProcessor struct {
}

func (r *defaultProcessor) OnConnected(connection miface.IConnection) {
	log.Printf("OnConnected")
	if connection.GetType() == miface.TypeConnectionConnect {

		p := parser.NewJsonParser()
		msg := message.NewMessage(1, service.MessageEcho{Msg: "echo"})
		data, _ := p.Marshal(msg)
		connection.Send(core.NewPackage(uint32(len(data)), data))
	} else {
		// 创建actor
		pActor = service.NewPlayerActor()
		pActor.Init()
		pActor.Run()
	}
}

func (r *defaultProcessor) OnDisconnect(connection miface.IConnection) {
}

func (r *defaultProcessor) OnProcess(connection miface.IConnection, pack miface.IPackage) {
	pActor.PutMessage(iface.NewMessage(iface.ActorMessageNetWork, nil, connection, pack))
}

func TestServer(t *testing.T) {
	// 创建actor
	s := mznet.NewServer(core.WithAddress("192.168.1.149:2100"),
		core.WithNetwork("tcp"),
		core.WithRouter(&defaultProcessor{}),
		core.WithTcpCodec(codec.NewCommonCodec()))

	s.Run()

	for true {
		time.Sleep(time.Second)
	}
}
