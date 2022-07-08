/**
 * @Author: dingQingHui
 * @Description:
 * @File: core
 * @Version: 1.0.0
 * @Date: 2022/7/8 16:58
 */

package core

import "sync/atomic"

var (
	gId int64
)

func GenId() int64 {
	return atomic.AddInt64(&gId, 1)
}
