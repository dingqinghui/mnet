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
)

var (
	ErrMailboxFull        = errors.New("put message to full mailbox")
	SystemMessageQueueLen = 1024
)

type (
	IMailBox interface {
		PostUserMessage(message interface{}) error
		PostSystemMessage(message interface{}) error
	}

	mailboxChan struct {
		userMessageQueue   chan interface{}
		systemMessageQueue chan interface{}
		receiver           IMessageReceiver
		stop               bool
		queueSize          int
		stopChan           chan struct{}
		message            interface{}
	}
)

func NewMailBoxChan(consumer IMessageReceiver, size int) IMailBox {
	mb := &mailboxChan{
		queueSize: size,
		receiver:  consumer,
	}
	mb.construction()
	mb.run()
	return mb
}

func (d *mailboxChan) construction() {
	d.stop = false
	d.userMessageQueue = make(chan interface{}, d.queueSize)
	d.systemMessageQueue = make(chan interface{}, SystemMessageQueueLen)
	d.stopChan = make(chan struct{}, SystemMessageQueueLen)
}

func (d *mailboxChan) PostUserMessage(message interface{}) error {
	select {
	case d.userMessageQueue <- message:
	default:
		return ErrMailboxFull
	}
	return nil
}

func (d *mailboxChan) PostSystemMessage(message interface{}) error {
	select {
	case d.systemMessageQueue <- message:
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
			d.receiver.Panic(e, d.message)
			d.run()
		}
	}()
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
}

func (d *mailboxChan) handleSystemMessageMsg(msg interface{}) {
	switch msg.(type) {
	case *MailBoxStop:
		d.handleStop()
	default:
	}
	d.message = msg
	d.receiver.ReceiveSystemMessage(msg)
}
func (d *mailboxChan) handleStop() {
	d.stopChan <- struct{}{}
}

func (d *mailboxChan) handleUserMessage(msg interface{}) {
	d.message = msg
	d.receiver.ReceiveUserMessage(msg)
}
