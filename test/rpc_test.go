/**
 * @Author: dingQingHui
 * @Description:
 * @File: rpc_test
 * @Version: 1.0.0
 * @Date: 2022/7/12 17:14
 */

package test

import (
	"fmt"
	"net/rpc"
	"testing"
)

type HelloService struct{}

func (h *HelloService) Hello(request string, reply *string) error {
	*reply = "hello:" + request
	return nil
}

func main() {
	//然后就可以将HelloService类型的对象注册为一个RPC服务：
	//其中rpc.Register函数调用会将对象类型中所有满足RPC规则的对象方法注册为RPC函数，
	//所有注册的方法会放在“HelloService”服务空间之下
	err := rpc.RegisterName("HelloService", new(HelloService))

	println(err)
}

func TestRpc(t *testing.T) {
	cli, err := rpc.Dial("tcp", ":1234")
	if err != nil {
		panic(err)
	}

	var reply string
	//在调用client.Call时，
	//第一个参数是用点号链接的RPC服务名字和方法名字，
	//第二和第三个参数分别我们定义RPC方法的两个参数。
	err = cli.Call("HelloService.Hello", "你好", &reply)
	if err != nil {
		panic(err)
	}
	fmt.Println(reply)
}
