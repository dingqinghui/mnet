/**
 * @Author: dingQingHui
 * @Description:
 * @File: iconnection_mgr
 * @Version: 1.0.0
 * @Date: 2022/7/8 18:42
 */

package miface

type (
	IConnectionMgr interface {
		Add(connection IConnection)
		Delete(connection IConnection)
	}
)
