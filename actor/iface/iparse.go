/**
 * @Author: dingQingHui
 * @Description:
 * @File: iparese
 * @Version: 1.0.0
 * @Date: 2022/7/11 14:30
 */

package iface

type (
	IParse interface {
		Register(msgId uint32, msg interface{}) error
		UnMarshal(data []byte) ([]interface{}, error)
		Marshal(...interface{}) ([]byte, error)
	}
)
