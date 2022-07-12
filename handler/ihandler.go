/**
 * @Author: dingQingHui
 * @Description:
 * @File: ihandler
 * @Version: 1.0.0
 * @Date: 2022/7/11 15:51
 */

package handler

import (
	"github.com/dingqinghui/mz/message"
	"github.com/dingqinghui/mz/mznet/miface"
)

type (
	HandlerFun func(connection miface.IConnection, msg message.IMessage)

	IHandler interface {
		SetHandler(msgId uint32, handler HandlerFun) error
		GetHandler(msgId uint32) (HandlerFun, error)
	}
)
