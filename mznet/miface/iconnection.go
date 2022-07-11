/**
 * @Author: dingQingHui
 * @Description:
 * @File: iconnection
 * @Version: 1.0.0
 * @Date: 2022/7/8 11:28
 */

package miface

import (
	"net"
)

type (
	IConnection interface {
		GetId() int64
		GetType() TypeConnection

		Send(message IPackage) bool

		GetLocalAddr() net.Addr
		GetRemoteAddr() net.Addr

		Close() bool
		IsClose() bool
	}
)
