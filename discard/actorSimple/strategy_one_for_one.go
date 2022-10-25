/**
 * @Author: dingQingHui
 * @Description:
 * @File: strategy_one_for_one
 * @Version: 1.0.0
 * @Date: 2022/10/11 15:44
 */

package actorNew

type (
	DeciderFun func(reason interface{}) Directive

	oneForOneStrategy struct {
		decider DeciderFun
	}
)

func NewOneForOneStrategy(decider DeciderFun) *oneForOneStrategy {
	return &oneForOneStrategy{decider: decider}
}

func (s *oneForOneStrategy) Handle(supervision iSupervision, who IPid, reason interface{}, message interface{}) {
	directive := s.decider(reason)
	switch directive {
	case ResumeDirective:
		supervision.Resume(who)
	case RestartDirective:
		supervision.Restart(who)
	case StopDirective:
		supervision.Stop(who)
	case EscalateStopDirective:
		supervision.EscalateFailure(reason, message)
	}
}
