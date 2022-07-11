/**
 * @Author: dingQingHui
 * @Description:
 * @File: handler
 * @Version: 1.0.0
 * @Date: 2022/7/11 14:03
 */

package handler

import (
	"fmt"
	"mz/iface"
	"sync"
)

type (
	Handler struct {
		sync.RWMutex
		m map[uint32]iface.HandlerFun
	}
)

func NewHandler() iface.IHandler {
	return &Handler{
		m: make(map[uint32]iface.HandlerFun),
	}
}

func (h *Handler) SetHandler(msgId uint32, handler iface.HandlerFun) error {
	h.Lock()
	defer h.Unlock()
	if _, ok := h.m[msgId]; ok {
		return fmt.Errorf("set handler is exist msgId:%d", msgId)
	}
	h.m[msgId] = handler
	return nil
}

func (h *Handler) GetHandler(msgId uint32) (iface.HandlerFun, error) {
	h.RLock()
	defer h.RUnlock()
	if handler, ok := h.m[msgId]; !ok {
		return nil, fmt.Errorf("set handler is not exist msgId:%d", msgId)
	} else {
		return handler, nil
	}
}
