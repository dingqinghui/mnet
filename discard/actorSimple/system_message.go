/**
 * @Author: dingQingHui
 * @Description:
 * @File: system_message
 * @Version: 1.0.0
 * @Date: 2022/10/6 17:49
 */

package actorNew

type iSystemMessage interface {
	SystemMessage()
}

type Started struct{}

func (_ *Started) SystemMessage() {}

type Stop struct{}

func (_ *Stop) SystemMessage() {}

type Restart struct{}

func (_ *Restart) SystemMessage() {}

type ChildStop struct{ Who IPid }

func (_ *ChildStop) SystemMessage() {}

type Fail struct {
	Reason  interface{}
	Pid     IPid
	Message interface{}
}

func (_ *Fail) SystemMessage() {}

type MailBoxSuspend struct{}

func (_ *MailBoxSuspend) SystemMessage() {}

type MailBoxResume struct{}

func (_ *MailBoxResume) SystemMessage() {}

type MailBoxRestart struct{}

func (_ *MailBoxRestart) SystemMessage() {}

type MailBoxStop struct{}

func (_ *MailBoxStop) SystemMessage() {}

var (
	startedMessage iSystemMessage = new(Started)
	stopMessage    iSystemMessage = new(Stop)
	restartMessage iSystemMessage = new(Restart)

	mailBoxSuspendMessage iSystemMessage = new(MailBoxSuspend)
	mailBoxResumeMessage  iSystemMessage = new(MailBoxResume)
	mailBoxReStartMessage iSystemMessage = new(MailBoxRestart)
	mailBoxStopMessage    iSystemMessage = new(MailBoxStop)
)
