/**
 * @Author: dingQingHui
 * @Description:
 * @File: irouter
 * @Version: 1.0.0
 * @Date: 2022/7/8 11:28
 */

package miface

type (
	IRouter interface {
		OnConnected(connection IConnection)
		OnDisconnect(connection IConnection)
		OnProcess(connection IConnection, message *Message)
	}
)
