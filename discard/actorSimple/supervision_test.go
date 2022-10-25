/**
 * @Author: dingQingHui
 * @Description:
 * @File: props_test
 * @Version: 1.0.0
 * @Date: 2022/9/27 11:25
 */

package actorNew

import (
	_ "runtime/pprof"
	"testing"
	"time"
)

func ReceiveA(ctx IContext) {
	switch ctx.Message().(type) {
	case string:
		println(ctx.Message().(string))
	case *MessageEnvelope:
		env := ctx.Message().(*MessageEnvelope)
		println(env.Message.(string))
		time.Sleep(time.Second * 2)
		_ = ctx.Send(env.Sender, "result message")
	}

	panic("test panic")
}

func TestSupervision_EscalateFailure(t *testing.T) {
	system := NewActorSystem()

	props := NewPropsWithFunc(ReceiveA, WithOnInit(OnInit))
	pid := system.Root().Spawn(props)
	_ = pid
	//err := system.Root().Send(pid, "2222222222222222222")
	//if err != nil {
	//	println(err)
	//}
	//
	////system.Root().Stop(pid)
	//err = system.Root().Send(pid, "2222222222222222222")

	//fut, err := system.Root().Call(pid, "3333333333333333", 1*time.Second)
	//if err != nil {
	//	println(err)
	//	return
	//}

	//result, ok := fut.Wait()
	//
	//result = fut.Result()

	time.Sleep(time.Second * 1000)
	// _, _ = ok, result
}
