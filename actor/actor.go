/**
 * @Author: dingQingHui
 * @Description:
 * @File: actor
 * @Version: 1.0.0
 * @Date: 2022/7/11 16:49
 */

package actor

import (
	"github.com/dingqinghui/mz/actor/iface"
)

type (
	BaseActor struct {
		id        uint64
		revChan   chan iface.IActorMessage
		protocols map[iface.ActorMessageType]*Protocol
	}
)

func NewBase() *BaseActor {
	return &BaseActor{
		revChan:   make(chan iface.IActorMessage, 1),
		protocols: make(map[iface.ActorMessageType]*Protocol),
	}
}

func (a *BaseActor) SetId(id uint64) {
	a.id = id
}

func (a *BaseActor) Init(_ ...interface{}) {

}
func (a *BaseActor) GetId() uint64 {
	return a.id
}

func (a *BaseActor) Run() {
	go func() {
		for true {
			select {
			case msg := <-a.revChan:
				a.Dispatch(msg)
			}
		}
	}()
}

func (a *BaseActor) RegistryProtocol(msgType iface.ActorMessageType, parse iface.IParse, dispatch iface.DispatchFun) {
	a.protocols[msgType] = &Protocol{
		parse:    parse,
		dispatch: dispatch,
	}
}

func (a *BaseActor) PutMessage(msg iface.IActorMessage) {
	a.revChan <- msg
}

func (a *BaseActor) Destroy() {

}
