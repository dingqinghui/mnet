/**
 * @Author: dingQingHui
 * @Description:
 * @File: client_test
 * @Version: 1.0.0
 * @Date: 2022/7/7 18:03
 */

package test

import (
	"github.com/dingqinghui/mz/actor"
	"net"
	"testing"
)

func TestClient(t *testing.T) {
	c, err := net.Dial("tcp", "192.168.1.149:22000")

	println(c, err)
	actor.NewService("client")

	select {}
}
