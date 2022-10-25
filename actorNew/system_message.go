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

type Fail struct {
	Reason  interface{}
	Pid     IPid
	Message interface{}
}

type Panic struct {
	Reason  interface{}
	Message interface{}
}

func (_ *Panic) SystemMessage() {
}

type MailBoxStop struct{}

func (_ *MailBoxStop) SystemMessage() {}

type MailBoxStared struct{}

func (_ *MailBoxStared) SystemMessage() {}

var (
	startedMessage iSystemMessage = new(Started)
	stopMessage    iSystemMessage = new(Stop)

	mailBoxStaredMessage iSystemMessage = new(MailBoxStared)
	mailBoxStopMessage   iSystemMessage = new(MailBoxStop)
)
