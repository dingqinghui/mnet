/**
 * @Author: dingQingHui
 * @Description:
 * @File: actor_system
 * @Version: 1.0.0
 * @Date: 2022/10/10 14:19
 */

package actorNew

import (
	"sync/atomic"
)

type (
	IActorSystem interface {
		Root() IContext
		NextId() int64
	}
)

type ActorSystem struct {
	id   int64
	root IContext
}

func NewActorSystem() IActorSystem {
	system := &ActorSystem{}
	system.root = newRootContext(system)
	return system
}

func (a *ActorSystem) NextId() int64 {
	return atomic.AddInt64(&a.id, 1)
}

func (a *ActorSystem) Root() IContext {
	return a.root
}

func newRootContext(system *ActorSystem) IContext {
	context := newActorContext(system, nil, nil)
	mb := NewMailBoxChan(context, 10)
	var pid IPid = &Pid{
		id:       system.NextId(),
		iProcess: newDefaultProcess(mb),
	}
	context.pid = pid
	var f ReceiveFunc = func(c IContext) {

	}
	context.actor = f
	return context
}
