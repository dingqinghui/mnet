/**
 * @Author: dingQingHui
 * @Description:
 * @File: datapack
 * @Version: 1.0.0
 * @Date: 2022/7/7 16:25
 */

package codec

import (
	"bytes"
	"encoding/binary"
	miface2 "mz/mznet/miface"
	"net"
)

// dataLen(4byte)+data

var headSize uint32 = 4

type (
	commonCodec struct {
	}
)

func NewCommonCodec() miface2.ICodec {
	return &commonCodec{}
}

func (d *commonCodec) Unpack(con net.Conn, message miface2.IPackage) error {
	headBuf := make([]byte, headSize)
	if _, err := con.Read(headBuf); err != nil {
		return err
	}
	var dataLen uint32
	reader := bytes.NewReader(headBuf)
	if err := binary.Read(reader, binary.BigEndian, &dataLen); err != nil {
		return err
	}
	dataBuf := make([]byte, dataLen)
	if _, err := con.Read(dataBuf); err != nil {
		return err
	}
	message.SetDataLen(dataLen)
	message.SetData(dataBuf)
	return nil
}
func (d *commonCodec) Pack(con net.Conn, msg miface2.IPackage) error {
	len := msg.GetDataLen() + headSize
	buf := make([]byte, len, len)
	// 写入头
	binary.BigEndian.PutUint32(buf, msg.GetDataLen())
	// 写书数据
	copy(buf[headSize:], msg.GetData())

	if _, err := con.Write(buf); err != nil {
		return err
	}
	return nil
}
