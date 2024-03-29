/**
 * @Author: dingQingHui
 * @Description:
 * @File: receiver
 * @Version: 1.0.0
 * @Date: 2022/10/6 18:34
 */

package actor

type (
	IMessageReceiver interface {
		ReceiveUserMessage(interface{})
		ReceiveSystemMessage(message interface{})
		Panic(reason, message interface{})
	}

	IActor interface {
		Receive(c IContext)
	}

	ReceiveFunc func(c IContext)
)

func (f ReceiveFunc) Receive(c IContext) {
	f(c)
}
