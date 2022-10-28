/**
 * @Author: dingQingHui
 * @Description:
 * @File: context
 * @Version: 1.0.0
 * @Date: 2022/9/27 11:43
 */

package actorNew

type (
	ActorContext struct {
		actor   IActor
		pid     IPid
		message interface{}
		system  IActorSystem
		props   *Props
		*sender
	}
)

func newActorContext(system IActorSystem) *ActorContext {
	context := &ActorContext{
		system: system,
		sender: defaultSender,
	}
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
	a.actor.Receive(a)
}

func (a *ActorContext) ReceiveSystemMessage(msg interface{}) {
	a.message = msg
	switch msg.(type) {
	case *Stop:
		a.handleStop()
	}
	a.message = msg
	a.actor.Receive(a)
}

func (a *ActorContext) handleStop() {
	_ = a.Send(a.Self(), mailBoxStopMessage)
}

func (a *ActorContext) Message() interface{} {
	return a.message
}

func (a *ActorContext) Spawn(p *Props) IPid {
	pid := p.spawn(a.system, a.Self())
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

func (a *ActorContext) Panic(reason, message interface{}) {
	a.ReceiveSystemMessage(&Panic{Reason: reason, Message: message})
}
