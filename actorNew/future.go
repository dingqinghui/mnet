/**
 * @Author: dingQingHui
 * @Description:
 * @File: future
 * @Version: 1.0.0
 * @Date: 2022/10/6 11:35
 */

package actorNew

import "time"

type future struct {
	pid     IPid
	result  interface{}
	timeout time.Duration
	waitCh  chan interface{}
	waitPid IPid
	err     error
}

func newFuture(timeout time.Duration, waitPid IPid) *future {
	f := &future{
		waitCh:  make(chan interface{}),
		timeout: timeout,
		waitPid: waitPid,
	}
	mb := NewMailBoxChan(f, 10)
	var pid IPid = &Pid{
		iProcess: newDefaultProcess(mb),
	}
	f.pid = pid
	return f
}

func (f *future) Pid() IPid {
	return f.pid
}

func (f *future) Wait() (interface{}, bool) {
	after := time.After(f.timeout)
	ok := false
	select {
	case <-after:
		ok = false
	case _, ok = <-f.waitCh:
		ok = true
	}
	f.Release()
	return f.result, ok
}

func (f *future) Release() {
	f.pid.Stop()
}
func (f *future) Result() (interface{}, error) {
	return f.result, f.err
}
func (f *future) Panic(reason, msg interface{}) {
	f.result = nil
	f.err = reason.(error)
	f.waitCh <- struct{}{}
}

func (f *future) ReceiveUserMessage(message interface{}) {
	f.result = message
	f.err = nil
	f.waitCh <- struct{}{}
}

func (f *future) ReceiveSystemMessage(message interface{}) {
	f.result = message
	f.err = nil
	f.waitCh <- struct{}{}
}
