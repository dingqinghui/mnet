/**
 * @Author: dingQingHui
 * @Description:
 * @File: Package
 * @Version: 1.0.0
 * @Date: 2022/7/7 16:33
 */

package core

import (
	"github.com/dingqinghui/mz/mznet/miface"
)

type (
	Package struct {
		data    []byte
		dataLen uint32
	}
)

func NewPackage(dataLen uint32, data []byte) miface.IPackage {
	return &Package{
		data:    data,
		dataLen: dataLen,
	}
}

func (m *Package) GetDataLen() uint32 {
	return m.dataLen
}

func (m *Package) GetData() []byte {
	return m.data
}

func (m *Package) SetDataLen(dataLen uint32) {
	m.dataLen = dataLen
}
func (m *Package) SetData(data []byte) {
	m.data = data
}
