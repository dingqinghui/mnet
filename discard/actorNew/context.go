/**
 * @Author: dingQingHui
 * @Description:
 * @File: IActor
 * @Version: 1.0.0
 * @Date: 2022/7/11 16:46
 */

package actorNew

import "time"

type (
	IContext interface {
		iBasePart
		iMessagePart
		iSenderPart
		iStopPart
		iSpawnerPart
	}

	iBasePart interface {
		Self() IPid
		Actor() IActor
		System() IActorSystem
	}

	iMessagePart interface {
		Message() interface{}
	}

	iSenderPart interface {
		Send(pid IPid, message interface{}) error
		Call(pid IPid, message interface{}, timeout time.Duration) (*future, error)
		Request(pid IPid, message interface{}) error
		RequestWithCustomSender(pid IPid, message interface{}, sender IPid) error
	}

	iStopPart interface {
		Stop(pid IPid)
	}

	iSpawnerPart interface {
		Spawn(p *Props) IPid
	}
)

type sender struct{}

func (r *sender) Send(pid IPid, message interface{}) error {
	return pid.SendUserMessage(message)
}

func (r *sender) Call(pid IPid, message interface{}, timeout time.Duration) (*future, error) {
	fur := newFuture(timeout, pid)
	env := &MessageEnvelope{
		Header:  nil,
		Message: message,
		Sender:  fur.Pid(),
	}
	err := pid.SendUserMessage(env)
	if err != nil {
		return nil, err
	}
	return fur, nil
}
func (r *sender) Request(pid IPid, message interface{}) error {
	panic("must implementation request")
	return nil
}
func (r *sender) RequestWithCustomSender(pid IPid, message interface{}, sender IPid) error {
	env := &MessageEnvelope{
		Header:  nil,
		Message: message,
		Sender:  sender,
	}
	return pid.SendUserMessage(env)
}

var defaultSender = new(sender)
