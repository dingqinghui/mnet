/**
 * @Author: dingQingHui
 * @Description:
 * @File: imessage
 * @Version: 1.0.0
 * @Date: 2022/7/11 15:14
 */

package message

type (
	IMessage interface {
		Msg() interface{}
		MsgId() uint32
	}
)
