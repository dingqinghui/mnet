/**
 * @Author: dingQingHui
 * @Description:
 * @File: idatapack
 * @Version: 1.0.0
 * @Date: 2022/7/7 16:23
 */

package miface

import (
	"net"
)

type (
	ICodec interface {
		Pack(con net.Conn, msg IMessage) error
		Unpack(con net.Conn, message IMessage) error
	}
)
