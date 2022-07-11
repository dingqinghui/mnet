/**
 * @Author: dingQingHui
 * @Description:
 * @File: json_parse_test
 * @Version: 1.0.0
 * @Date: 2022/7/11 15:28
 */

package parser

import (
	"mz/message"
	"testing"
)

type (
	Value struct {
		Id uint32
	}
)

func TestJsonParse(t *testing.T) {
	parse := NewJsonParser()
	if err := parse.Register(1, &Value{}); err != nil {
		println(err)
		return
	}
	data, err := parse.Marshal(message.NewMessage(1, &Value{Id: 1}))
	if err != nil {
		println(err)
		return
	}
	msg, err := parse.UnMarshal(data)
	if err != nil {
		println(err)
		return
	}
	println(msg)
}
