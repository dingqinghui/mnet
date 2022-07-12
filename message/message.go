/**
 * @Author: dingQingHui
 * @Description:
 * @File: message
 * @Version: 1.0.0
 * @Date: 2022/7/11 15:20
 */

package message

type Message struct {
	msgId uint32
	msg   interface{}
}

func NewMessage(msgId uint32, msg interface{}) IMessage {
	return &Message{
		msgId: msgId,
		msg:   msg,
	}
}

func (m *Message) MsgId() uint32 {
	return m.msgId
}

func (m *Message) Msg() interface{} {
	return m.msg
}
