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
		Pack(con net.Conn, msg IPackage) error
		Unpack(con net.Conn, message IPackage) error
	}
)
