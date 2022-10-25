/**
 * @Author: dingQingHui
 * @Description:
 * @File: supervision
 * @Version: 1.0.0
 * @Date: 2022/10/11 15:00
 */

package actorNew

type (
	iSupervision interface {
		Parent() IPid

		Children() []IPid
		AddChild(pid IPid)
		RemoveChild(pid IPid)

		Resume(children ...IPid)
		Restart(children ...IPid)
		Stop(children ...IPid)
		EscalateFailure(reason, message interface{})
	}

	// 处理子actor异常策略
	iSupervisionStrategy interface {
		Handle(supervision iSupervision, who IPid, reason interface{}, message interface{})
	}
)

var (
	defaultDeciderFun = func(reason interface{}) Directive {
		return RestartDirective
	}
	defaultSupervisionStrategy = NewOneForOneStrategy(defaultDeciderFun)
)

type supervision struct {
	context  IContext
	parent   IPid
	children *PIDSet
}

func newSupervision(parent IPid, context IContext) iSupervision {
	return &supervision{
		parent:   parent,
		context:  context,
		children: NewPIDSet(),
	}
}

func (s *supervision) AddChild(pid IPid) {
	s.children.Add(pid)
}

func (s *supervision) RemoveChild(pid IPid) {
	s.children.Remove(pid)
}

func (s *supervision) Parent() IPid {
	return s.parent
}
func (s *supervision) Children() []IPid {
	return s.children.Values()
}

func (s *supervision) Resume(children ...IPid) {
	for _, child := range children {
		_ = child.SendSystemMessage(mailBoxResumeMessage)
	}
}

func (s *supervision) Restart(children ...IPid) {
	for _, child := range children {
		_ = child.SendSystemMessage(restartMessage)
	}
}

func (s *supervision) Stop(children ...IPid) {
	for _, child := range children {
		_ = child.SendSystemMessage(stopMessage)
	}
}

//
// EscalateFailure
// @Description: 将异常传递给父节点
// @receiver s
// @param reason
// @param msg
//
func (s *supervision) EscalateFailure(reason, msg interface{}) {
	childFailMessage := &Fail{Reason: reason, Pid: s.context.Self(), Message: msg}
	_ = s.context.Self().SendSystemMessage(mailBoxSuspendMessage)
	if s.Parent() == nil {
		s.handleRootPanic(childFailMessage)
	} else {
		_ = s.Parent().SendSystemMessage(childFailMessage)
	}
}
func (s *supervision) handleRootPanic(childFailMessage *Fail) {
	panic("root panic")
}
