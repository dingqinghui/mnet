/**
 * @Author: dingQingHui
 * @Description:
 * @File: cmd_codec
 * @Version: 1.0.0
 * @Date: 2022/7/8 11:47
 */

package codec

import (
	"bufio"
	miface2 "mz/mznet/miface"
	"net"
)

type (
	cmdCodec struct {
	}
)

func NewCmdCodec() miface2.ICodec {
	return &cmdCodec{}
}

func (d *cmdCodec) Unpack(con net.Conn, message miface2.IMessage) error {
	reader := bufio.NewReader(con)
	str, err := reader.ReadString('\n')
	if err != nil {
		return err
	}
	message.SetDataLen(uint32(len(str)))
	message.SetData([]byte(str))
	return nil
}

func (d *cmdCodec) Pack(con net.Conn, msg miface2.IMessage) error {
	if _, err := con.Write(msg.GetData()); err != nil {
		return err
	}
	return nil
}
