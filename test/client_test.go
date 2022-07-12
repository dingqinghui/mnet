/**
 * @Author: dingQingHui
 * @Description:
 * @File: client_test
 * @Version: 1.0.0
 * @Date: 2022/7/7 18:03
 */

package test

import (
	"github.com/dingqinghui/mz/service"
	"testing"
)

func TestClient(t *testing.T) {

	watchDog := service.NewClient()
	watchDog.Init()
	watchDog.Run()

	select {}
}
