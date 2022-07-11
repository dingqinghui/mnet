/**
 * @Author: dingQingHui
 * @Description:
 * @File: ihandler
 * @Version: 1.0.0
 * @Date: 2022/7/11 15:51
 */

package iface

import "mz/mznet/miface"

type (
	HandlerFun func(server miface.IServer, connection miface.IConnection, msg IMessage)

	IHandler interface {
		SetHandler(msgId uint32, handler HandlerFun) error
		GetHandler(msgId uint32) (HandlerFun, error)
	}
)
