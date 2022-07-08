/**
 * @Author: dingQingHui
 * @Description:
 * @File: server_test
 * @Version: 1.0.0
 * @Date: 2022/7/7 18:04
 */

package test

import (
	"mz/mznet"
	"mz/mznet/core"
	"testing"
)

func TestServer(t *testing.T) {
	s := mznet.NewServer(core.WithAddress("192.168.1.170:2100"),
		core.WithNetwork("udp"),
		core.WithRouter(&defaultRouter1{}))

	s.Run()

	select {}
}
