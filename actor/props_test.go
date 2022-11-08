/**
 * @Author: dingQingHui
 * @Description:
 * @File: props_test
 * @Version: 1.0.0
 * @Date: 2022/9/27 11:25
 */

package actor

import (
	"testing"
	"time"
)

type childActor struct {
}

func (ch *childActor) Receive(ctx IContext) {
	switch ctx.Message().(type) {
	case string:
		println("child message", ctx.Message().(string))
	case *MessageEnvelope:
		env := ctx.Message().(*MessageEnvelope)
		println("child message", env.Message.(string))
	default:
		println("child message")
	}
}

func Receive(ctx IContext) {
	switch ctx.Message().(type) {
	case string:
		println(ctx.Message().(string))
	case *MessageEnvelope:
		env := ctx.Message().(*MessageEnvelope)
		println(env.Message.(string))
		_ = ctx.Send(env.Sender, "result message")
	case *Panic:
		ctx.Stop(ctx.Self())
	default:
		panic("test panic")
	}
}

func OnInit(ctx IContext) {
	println("OnInit")
}

func TestNewPropsWithFunc(t *testing.T) {
	system := NewActorSystem()
	// 启动
	props := NewPropsWithFunc(Receive, WithOnInit(OnInit))
	pid := system.Root().Spawn(props)
	// 异步消息
	err := system.Root().Send(pid, "2222222222222222222")
	if err != nil {
		println(err)
	}
	// 同步消息
	fut, err := system.Root().Call(pid, "3333333333333333", 5*time.Second)
	if err != nil {
		println(err)
		return
	}
	result, ok := fut.Wait()
	result, err = fut.Result()
	// 关闭
	system.Root().Stop(pid)
	err = system.Root().Send(pid, "2222222222222222222")

	time.Sleep(time.Second * 10)
	_, _ = ok, result
}
