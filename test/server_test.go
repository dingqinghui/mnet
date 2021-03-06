/**
 * @Author: dingQingHui
 * @Description:
 * @File: server_test
 * @Version: 1.0.0
 * @Date: 2022/7/7 18:04
 */

package test

import (
	"github.com/dingqinghui/mz/service"
	"testing"
)

func TestServer(t *testing.T) {
	watchDog := service.NewWatchdog()
	watchDog.Init()
	watchDog.Run()
	select {}
}
