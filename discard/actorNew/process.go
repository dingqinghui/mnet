/**
 * @Author: dingQingHui
 * @Description:
 * @File: process
 * @Version: 1.0.0
 * @Date: 2022/10/6 18:35
 */

package actorNew

import (
	"errors"
	"sync/atomic"
)

type (
	iProcess interface {
		SendUserMessage(interface{}) error
		SendSystemMessage(interface{}) error
		Stop()
		IsStop() bool
	}
)

type defaultProcess struct {
	dead    int32
	mailbox IMailBox
}

func newDefaultProcess(mailbox IMailBox) iProcess {
	return &defaultProcess{
		mailbox: mailbox,
		dead:    0,
	}
}
func (s *defaultProcess) SendUserMessage(msg interface{}) error {
	if s.IsStop() {
		return errors.New("mailbox dead")
	}
	return s.mailbox.PostUserMessage(msg)
}

func (s *defaultProcess) SendSystemMessage(msg interface{}) error {
	if s.IsStop() {
		return errors.New("mailbox dead")
	}
	return s.mailbox.PostSystemMessage(msg)
}

func (s *defaultProcess) Stop() {
	if atomic.CompareAndSwapInt32(&s.dead, 0, 1) {
		_ = s.mailbox.PostSystemMessage(stopMessage)
	}
}
func (s *defaultProcess) IsStop() bool {
	return atomic.CompareAndSwapInt32(&s.dead, 1, 1)
}
