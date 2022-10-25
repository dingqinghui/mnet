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

//
// Wait
// @Description: 等待返回结果,返回结果 超时 接受者关闭都会出发wait返回
// @receiver f
// @return interface{}
// @return bool
//
func (f *future) Wait() (interface{}, bool) {
	after := time.After(f.timeout)
	t := time.NewTicker(1 * time.Second)
	defer t.Stop()
	ok := false
loop:
	select {
	case <-after:
		ok = false
	case _, ok = <-f.waitCh:
		ok = true
	case <-t.C:
		if !f.pid.IsStop() {
			goto loop
		} else {
			ok = false
		}
	}
	f.Release()
	return f.result, ok
}

func (f *future) Release() {
	f.pid.Stop()
}

func (f *future) Result() interface{} {
	return f.result
}
func (f *future) EscalateFailure(reason, msg interface{}) {

}
func (f *future) ReceiveUserMessage(msg interface{}) {
	f.result = msg
	f.waitCh <- 1
}
func (f *future) ReceiveSystemMessage(message interface{}) {

}
