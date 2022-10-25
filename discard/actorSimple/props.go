/**
 * @Author: dingQingHui
 * @Description:
 * @File: props
 * @Version: 1.0.0
 * @Date: 2022/9/27 10:49
 */

package actorNew

var (
	defaultSpawnFun = func(p *Props, system IActorSystem, parent IPid) IPid {
		context := newActorContext(system, parent, p)
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
		strategy iSupervisionStrategy
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

func (p *Props) supervisorStrategy() iSupervisionStrategy {
	if p.strategy == nil {
		return defaultSupervisionStrategy
	}
	return p.strategy
}
