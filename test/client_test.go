/**
 * @Author: dingQingHui
 * @Description:
 * @File: client_test
 * @Version: 1.0.0
 * @Date: 2022/7/7 18:03
 */

package test

import (
	"github.com/dingqinghui/mz/mznet"
	"github.com/dingqinghui/mz/mznet/codec"
	"github.com/dingqinghui/mz/mznet/core"
	"testing"
)

func TestClient(t *testing.T) {

	c := mznet.NewClient(core.WithAddress("192.168.1.149:2100"),
		core.WithNetwork("tcp"),
		core.WithRouter(&defaultProcessor{}),
		core.WithTcpCodec(codec.NewCommonCodec()))

	c.Connect()

	select {}
}
