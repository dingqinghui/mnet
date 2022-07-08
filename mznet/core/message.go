/**
 * @Author: dingQingHui
 * @Description:
 * @File: message
 * @Version: 1.0.0
 * @Date: 2022/7/7 16:33
 */

package core

import (
	"mz/mznet/miface"
)

type (
	message struct {
		data    []byte
		dataLen uint32
	}
)

func NewMessage(dataLen uint32, data []byte) miface.IMessage {
	return &message{
		data:    data,
		dataLen: dataLen,
	}
}

func (m *message) GetDataLen() uint32 {
	return m.dataLen
}

func (m *message) GetData() []byte {
	return m.data
}

func (m *message) SetDataLen(dataLen uint32) {
	m.dataLen = dataLen
}
func (m *message) SetData(data []byte) {
	m.data = data
}
