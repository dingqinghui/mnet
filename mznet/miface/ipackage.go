/**
 * @Author: dingQingHui
 * @Description:
 * @File: imessage
 * @Version: 1.0.0
 * @Date: 2022/7/8 11:24
 */

package miface

type IPackage interface {
	GetDataLen() uint32
	GetData() []byte

	SetDataLen(uint32)
	SetData([]byte)
}
