/**
 * @Author: dingQingHui
 * @Description:
 * @File: manager
 * @Version: 1.0.0
 * @Date: 2022/7/12 16:31
 */

package actor

import (
	"errors"
	"github.com/dingqinghui/mz/actorNew"
	"sync/atomic"
)

type (
	Factory func() actorNew.IActor
)

var (
	actorReg map[string]Factory

	manager map[uint64]actorNew.IActor

	genId uint64
)

func init() {
	actorReg = make(map[string]Factory)
	manager = make(map[uint64]actorNew.IActor)
}

func RegistryActor(name string, factory Factory) {
	actorReg[name] = factory
}

func NewService(name string, args ...interface{}) actorNew.IActor {
	factory, ok := actorReg[name]
	if !ok {
		return nil
	}
	id := atomic.AddUint64(&genId, 1)

	actor := factory()
	actor.SetId(id)

	actor.Init(args)
	actor.Run()

	manager[id] = actor

	return actor
}

func Call(serviceId uint64, methodName string, args ...interface{}) []interface{} {
	actor := manager[serviceId]
	if actor == nil {
		return []interface{}{errors.New("actor not exist")}
	}

	//actor.SendUserMessage()

	return nil
}
