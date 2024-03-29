/**
 * @Author: dingQingHui
 * @Description:
 * @File: props
 * @Version: 1.0.0
 * @Date: 2022/9/27 10:49
 */

package actor

var (
	defaultSpawnFun = func(p *Props, system IActorSystem, parent IPid) IPid {
		context := newActorContext(system)
		mb := NewMailBoxChan(context, 10)
		var pid IPid = &Pid{
			id:       system.NextId(),
			iProcess: newDefaultProcess(mb),
		}
		context.pid = pid
		context.actor = p.producer()
		for _, init := range p.onInits {
			init(context)
		}
		_ = context.Send(context.Self(), startedMessage)
		return pid
	}
)

type (
	Producer func() IActor
	Spawner  func(p *Props, system IActorSystem, parent IPid) IPid
	InitFunc func(ctx IContext)
	Props    struct {
		producer Producer
		onInits  []InitFunc
		spawner  Spawner
	}
)

func (p *Props) getSpawnFun() Spawner {
	if p.spawner == nil {
		return defaultSpawnFun
	}
	return p.spawner
}

func (p *Props) spawn(system IActorSystem, parent IPid) IPid {
	return p.getSpawnFun()(p, system, parent)
}
