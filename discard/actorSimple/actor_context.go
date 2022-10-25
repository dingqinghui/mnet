/**
 * @Author: dingQingHui
 * @Description:
 * @File: context
 * @Version: 1.0.0
 * @Date: 2022/9/27 11:43
 */

package actorNew

import (
	"reflect"
)

type (
	ActorContext struct {
		actor   IActor
		pid     IPid
		message interface{}
		system  IActorSystem
		props   *Props
		iSupervision
		*sender
	}
)

func newActorContext(system IActorSystem, parent IPid, props *Props) *ActorContext {
	context := &ActorContext{
		system: system,
		sender: defaultSender,
		props:  props,
	}
	context.iSupervision = newSupervision(parent, context)
	return context
}

func (a *ActorContext) System() IActorSystem {
	return a.system
}

func (a *ActorContext) Actor() IActor {
	return a.actor
}

func (a *ActorContext) Self() IPid {
	return a.pid
}

func (a *ActorContext) ReceiveUserMessage(msg interface{}) {
	a.message = msg
	DebugLog("actor:%d receive user message %v %v\n", a.Self().Id(), msg, reflect.TypeOf(msg))
	a.actor.Receive(a)
}

func (a *ActorContext) ReceiveSystemMessage(msg interface{}) {
	a.message = msg

	DebugLog("actor:%d receive system message %v %v\n", a.Self().Id(), msg, reflect.TypeOf(msg))

	switch msg.(type) {
	case *Fail:
		a.handleFail(msg)
	case *Started:
		a.handleStarted(msg)
	case *Stop:
		a.handleStop(msg)
	case *Restart:
		a.handleRestart(msg)
	case *ChildStop:
		a.handleChildStop(msg)
	}
	a.message = msg
	a.actor.Receive(a)
}

func (a *ActorContext) handleFail(msg interface{}) {
	fail := msg.(*Fail)
	if a.props == nil {
		defaultSupervisionStrategy.Handle(a.iSupervision, fail.Pid, fail.Reason, fail.Message)
	} else {
		a.props.supervisorStrategy().Handle(a.iSupervision, fail.Pid, fail.Reason, fail.Message)
	}

	DebugLog("actor:%d child fail  %v \n", a.Self().Id(), fail)
}

func (a *ActorContext) handleRestart(msg interface{}) {
	_ = a.Self().SendSystemMessage(mailBoxReStartMessage)
	// product actor
	a.actor = a.props.producer()

	DebugLog("restart actor:%d\n", a.Self().Id())
}
func (a *ActorContext) handleStarted(msg interface{}) {

}

func (a *ActorContext) handleStop(msg interface{}) {
	_ = a.Self().SendSystemMessage(mailBoxStopMessage)
	// stop all child
	children := a.iSupervision.Children()
	a.iSupervision.Stop(children...)
	// notify parent
	_ = a.iSupervision.Parent().SendSystemMessage(&ChildStop{Who: a.Self()})
	DebugLog("stop actor %d \n", a.Self().Id())
}

func (a *ActorContext) handleChildStop(msg interface{}) {
	childStopMessage := msg.(ChildStop)
	a.iSupervision.RemoveChild(childStopMessage.Who)
}

func (a *ActorContext) Message() interface{} {
	return a.message
}

func (a *ActorContext) Spawn(p *Props) IPid {
	pid := p.spawn(a.system, a.Self())
	a.iSupervision.AddChild(pid)
	DebugLog("%d spawn child %d\n", a.Self().Id(), pid.Id())
	return pid
}

func (a *ActorContext) Request(pid IPid, message interface{}) error {
	env := &MessageEnvelope{
		Header:  nil,
		Message: message,
		Sender:  a.Self(),
	}
	return pid.SendUserMessage(env)
}

func (a *ActorContext) Stop(pid IPid) {
	pid.Stop()
}
