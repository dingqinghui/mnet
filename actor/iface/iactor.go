/**
 * @Author: dingQingHui
 * @Description:
 * @File: IActor
 * @Version: 1.0.0
 * @Date: 2022/7/11 16:46
 */

package iface

var (
	ActorMessageSocket ActorMessageType = 1
	ActorMessageRpc    ActorMessageType = 2
)

type (
	DispatchFun func(msg IActorMessage, args ...interface{})

	ActorMessageType int

	IActor interface {
		GetId() uint64
		Init(args ...interface{})
		Run()
		PutMessage(msg IActorMessage)
		RegistryProtocol(msgType ActorMessageType, parse IParse, dispatch DispatchFun)
	}

	IActorMessage interface {
		GetType() ActorMessageType
		GetData() []byte
	}
)
