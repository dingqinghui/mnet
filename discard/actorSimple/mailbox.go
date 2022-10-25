/**
 * @Author: dingQingHui
 * @Description:
 * @File: mailbox
 * @Version: 1.0.0
 * @Date: 2022/10/5 11:02
 */

package actorNew

import (
	"errors"
	"reflect"
)

var (
	ErrMailboxFull        = errors.New("put message to full mailbox")
	SystemMessageQueueLen = 10
)

type (
	IMailBox interface {
		PostUserMessage(message interface{}) error
		PostSystemMessage(message interface{}) error
	}

	mailboxChan struct {
		userMessageQueue    chan interface{}
		systemMessageQueue  chan interface{}
		receiver            IMessageReceiver
		stop                bool
		suspend             bool
		userTmpMessageQueue chan interface{}
		userTmpMessage      interface{}
		queueSize           int
		stopChan            chan struct{}
		message             interface{}
	}
)

func NewMailBoxChan(consumer IMessageReceiver, size int) IMailBox {
	mb := &mailboxChan{
		queueSize: size,
		receiver:  consumer,
	}
	mb.construction()
	mb.run()

	mb.systemMessageQueue <- startedMessage
	return mb
}

func (d *mailboxChan) construction() {
	d.stop = false
	d.suspend = false
	d.userMessageQueue = make(chan interface{}, d.queueSize)
	d.systemMessageQueue = make(chan interface{}, SystemMessageQueueLen)
	d.stopChan = make(chan struct{}, SystemMessageQueueLen)
}

func (d *mailboxChan) PostUserMessage(message interface{}) error {
	select {
	case d.userMessageQueue <- message:
		DebugLog("mailbox post user message %+v %s\n", message, reflect.TypeOf(message).String())
	default:
		DebugLog("put message to full mailbox %+v %s\n", message, reflect.TypeOf(message).String())
		return ErrMailboxFull
	}
	return nil
}

func (d *mailboxChan) PostSystemMessage(message interface{}) error {

	select {
	case d.systemMessageQueue <- message:
		DebugLog("mailbox post system message %+v %s\n", message, reflect.TypeOf(message).String())
	default:
		return ErrMailboxFull
	}
	return nil
}

func (d *mailboxChan) run() {
	go d.schedule()
}

func (d *mailboxChan) schedule() {
	defer func() {
		if e := recover(); e != nil {
			DebugLog("mailbox recover err:%v\n", e)
			d.receiver.EscalateFailure(e, d.message)
			d.run()
		}
	}()

	DebugLog("start mailbox\n")
Exit:
	for {
		select {
		case <-d.stopChan:
			break Exit
		case systemMsg, _ := <-d.systemMessageQueue:
			d.handleSystemMessageMsg(systemMsg)
		case userMsg, _ := <-d.userMessageQueue:
		priority:
			for {
				select {
				case systemMsg, _ := <-d.systemMessageQueue:
					d.handleSystemMessageMsg(systemMsg)
				default:
					break priority
				}
			}
			d.handleUserMessage(userMsg)
		}
	}
	DebugLog("exit mailbox\n")
}

func (d *mailboxChan) handleSystemMessageMsg(msg interface{}) {
	switch msg.(type) {
	case *MailBoxStop:
		d.handleStop()
	case *MailBoxSuspend:
		d.handleSuspend()
	case *MailBoxResume:
		d.handleResume()
	case *MailBoxRestart:
		d.handleRestart()
	default:

	}

	d.message = msg
	d.receiver.ReceiveSystemMessage(msg)
}

func (d *mailboxChan) handleStop() {

}

func (d *mailboxChan) handleSuspend() {
	d.suspend = true
	d.userTmpMessageQueue = d.userMessageQueue
	d.userMessageQueue = make(chan interface{}, 0)

	d.userTmpMessage = nil
}

func (d *mailboxChan) handleResume() {
	d.suspend = true

	d.receiver.ReceiveUserMessage(d.userTmpMessage)
	d.userTmpMessage = nil

	d.userMessageQueue = d.userTmpMessageQueue

}

func (d *mailboxChan) handleRestart() {
	d.stop = false
	d.suspend = false
	d.systemMessageQueue <- startedMessage
}

func (d *mailboxChan) handleUserMessage(msg interface{}) {
	if d.suspend {
		d.userTmpMessage = msg
		return
	}
	d.message = msg
	d.receiver.ReceiveUserMessage(msg)
}
