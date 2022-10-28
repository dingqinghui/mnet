/**
 * @Author: dingQingHui
 * @Description:
 * @File: close
 * @Version: 1.0.0
 * @Date: 2022/10/28 23:09
 */

package mznet

import (
	"errors"
	"sync/atomic"
)

type ICloser interface {
	Close() error
	Done() <-chan struct{}
	IsClose() bool
}

var defaultCloser ICloser = &closer{
	done: make(chan struct{}),
}

type closer struct {
	close int32
	done  chan struct{}
}

func (c *closer) Close() error {
	if !atomic.CompareAndSwapInt32(&c.close, 0, 1) {
		return errors.New("object repeated close")
	}
	close(c.done)
	return nil
}
func (c *closer) Done() <-chan struct{} {
	return c.done
}
func (c *closer) IsClose() bool {
	return atomic.CompareAndSwapInt32(&c.close, 1, 1)
}
