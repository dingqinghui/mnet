/**
 * @Author: dingQingHui
 * @Description:
 * @File: connection_mgr
 * @Version: 1.0.0
 * @Date: 2022/7/8 18:38
 */

package core

import (
	"mz/mznet/miface"
	"sync"
)

type (
	ConnectionMgr struct {
		sync.Mutex
		cMap map[int64]miface.IConnection
	}
)

func NewConnectionMgr() miface.IConnectionMgr {
	return &ConnectionMgr{
		cMap: make(map[int64]miface.IConnection),
	}
}

func (c *ConnectionMgr) Add(connection miface.IConnection) {
	c.Lock()
	defer c.Unlock()

	c.cMap[connection.GetId()] = connection
}
func (c *ConnectionMgr) Delete(connection miface.IConnection) {
	c.Lock()
	defer c.Unlock()
	delete(c.cMap, connection.GetId())
}
