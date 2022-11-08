/**
 * @Author: dingQingHui
 * @Description:
 * @File: actor_net_test
 * @Version: 1.0.0
 * @Date: 2022/11/7 10:31
 */

package main

import (
	"fmt"
	"github.com/dingqinghui/mz/actor"
	"github.com/dingqinghui/mz/mznet"
	"reflect"
	"testing"
	"time"
)

type ActorClient struct {
	client mznet.IClient
}

func (a *ActorClient) Receive(ctx actor.IContext) {
	fmt.Printf("%v\n", reflect.TypeOf(ctx.Message()).String())
	switch ctx.Message().(type) {
	case string:
		println(ctx.Message().(string))
	case *actor.MessageEnvelope:
		env := ctx.Message().(*actor.MessageEnvelope)
		println(env.Message.(string))
		_ = ctx.Send(env.Sender, "result message")
	case *actor.Started:
		a.client = NewClient(ctx.Self(), &mznet.ClientConfig{Network: "tcp", Address: "127.0.0.1:5642"})
		_ = a.client.Connect()
	case *actor.Panic:
		ctx.Stop(ctx.Self())
	case *NetConnected:
		connectedMessage := ctx.Message().(*NetConnected)
		data := []byte("11111111222222\n")
		connectedMessage.connection.Send(data)
	case *NetClosed:
	case *NetProcess:
		connectedMessage := ctx.Message().(*NetProcess)
		data := []byte("11111111222222\n")
		connectedMessage.connection.Send(data)
	case *NetError:
		nError := ctx.Message().(*NetError)
		nError.connection.Close()
		//fmt.Printf("%v\n", reflect.TypeOf(ctx.Message()).String())
	default:

	}
}

func OnInitClient(ctx actor.IContext) {
	println("OnInit")

}

func TestNewClient(t *testing.T) {

	system := actor.NewActorSystem()
	// 启动
	props := actor.NewPropsWithProducer(func() actor.IActor { return &ActorClient{} }, actor.WithOnInit(OnInitClient))

	pid := system.Root().Spawn(props)
	_ = pid
	_ = system
	time.Sleep(time.Hour)
}

type ActorServer struct {
}

func (a *ActorServer) Receive(ctx actor.IContext) {
	fmt.Printf("%v\n", reflect.TypeOf(ctx.Message()).String())
	switch ctx.Message().(type) {
	case string:
		println(ctx.Message().(string))
	case *actor.MessageEnvelope:
		env := ctx.Message().(*actor.MessageEnvelope)
		println(env.Message.(string))
		_ = ctx.Send(env.Sender, "result message")
	case *actor.Started:
		server := NewServer(ctx.Self(), &mznet.ServerConfig{Network: "tcp", ListenAddress: "127.0.0.1:5642"})
		_ = server.RunEventLoop()
	case *actor.Panic:
		ctx.Stop(ctx.Self())
	case *NetConnected:

	case *NetClosed:
	case *NetProcess:
		connectedMessage := ctx.Message().(*NetProcess)
		data := []byte("11111111222222\n")
		connectedMessage.connection.Send(data)
	case *NetError:
		nError := ctx.Message().(*NetError)
		nError.connection.Close()
		//fmt.Printf("%v\n", reflect.TypeOf(ctx.Message()).String())
	default:

	}
}

func OnInitServer(ctx actor.IContext) {
	println("OnInit")

}

func TestNewServer(t *testing.T) {
	system := actor.NewActorSystem()
	// 启动
	props := actor.NewPropsWithProducer(func() actor.IActor { return &ActorServer{} }, actor.WithOnInit(OnInitServer))
	pid := system.Root().Spawn(props)
	_ = props
	_ = pid
	time.Sleep(time.Hour)
}
