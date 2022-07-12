/**
 * @Author: dingQingHui
 * @Description:
 * @File: json
 * @Version: 1.0.0
 * @Date: 2022/7/11 11:51
 */

package parser

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dingqinghui/mz/actor/iface"
	"reflect"
	"sync"
)

type (
	JsonParser struct {
		sync.RWMutex
		m map[uint32]reflect.Type
	}
)

func NewJsonParser() iface.IParse {
	return &JsonParser{
		m: make(map[uint32]reflect.Type),
	}
}

func (j *JsonParser) Register(msgId uint32, msg interface{}) error {
	t := reflect.TypeOf(msg)
	if t == nil || t.Kind() != reflect.Ptr {
		return errors.New("json message pointer required")
	}

	name := t.Elem().Name()
	j.Lock()
	defer j.Unlock()
	if _, ok := j.m[msgId]; ok {
		return fmt.Errorf("json register fail  name:%s", name)
	}
	j.m[msgId] = t
	return nil
}

func (j *JsonParser) Marshal(args ...interface{}) ([]byte, error) {
	msgId := args[0].(uint32)
	msg := args[1]
	body, err := json.Marshal(msg)
	if err != nil {
		return nil, err
	}
	msgIdLen := binary.Size(msgId)

	length := len(body) + msgIdLen

	buf := make([]byte, length, length)

	binary.BigEndian.PutUint32(buf, msgId)

	copy(buf[msgIdLen:], body)

	return buf, nil
}

func (j *JsonParser) UnMarshal(data []byte) ([]interface{}, error) {
	if len(data) <= 0 {
		return nil, nil
	}
	msgId := binary.BigEndian.Uint32(data)
	t, ok := j.m[msgId]
	if !ok {
		return nil, fmt.Errorf("json register fail  msgId:%d", msgId)
	}
	msg := reflect.New(t.Elem()).Interface()
	msgIdLen := binary.Size(msgId)
	if err := json.Unmarshal(data[msgIdLen:], msg); err != nil {
		return nil, err
	}
	return []interface{}{msgId, msg}, nil
}
