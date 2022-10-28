/**
 * @Author: dingQingHui
 * @Description:
 * @File: codec
 * @Version: 1.0.0
 * @Date: 2022/10/28 14:14
 */

package mznet

import (
	"bufio"
	"errors"
	"io"
)

type ICodec interface {
	Pack(writer io.Writer, msg interface{}) error
	UnPack(reader io.Reader) (interface{}, error)
}

var DefaultCodec = &defaultCodec{}

type defaultCodec struct{}

func (defaultCodec) Pack(writer io.Writer, msg interface{}) error {
	data, ok := msg.([]byte)
	if !ok {
		return errors.New("change type fail")
	}
	_, err := writer.Write(data)
	return err
}

func (defaultCodec) UnPack(reader io.Reader) (interface{}, error) {
	readerBuf := bufio.NewReader(reader)
	str, err := readerBuf.ReadString('\n')
	if err != nil {
		return nil, err
	}
	return str, nil
}
