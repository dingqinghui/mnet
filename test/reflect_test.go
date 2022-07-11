/**
 * @Author: dingQingHui
 * @Description:
 * @File: reflect_test
 * @Version: 1.0.0
 * @Date: 2022/7/11 9:35
 */

package test

import (
	"mz/mznet"
	"reflect"
	"testing"
)

type (
	Value struct {
		id int
	}
)

func TestReflect(t *testing.T) {
	a := mznet.NewClient()
	tt := reflect.TypeOf(&a)
	s := tt.Elem().Name()
	println(s)
	b := reflect.New(tt.Elem())
	c := b.Interface().(*Value)
	c.id = 1
}
