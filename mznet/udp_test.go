/**
 * @Author: dingQingHui
 * @Description:
 * @File: udp_test
 * @Version: 1.0.0
 * @Date: 2022/10/28 18:22
 */

package mznet

import (
	"testing"
	"time"
)

func TestNewUdpServer(t *testing.T) {
	serve := newUdpServer(WithServerNetwork("udp"), WithServerAddress("0.0.0.0:8889"), WithServerEventListener(new(defaultEventListener)))
	serve.RunEventLoop()

	time.Sleep(time.Second * 10000)
}

func TestNewUdpClient(t *testing.T) {
	client := newUdpClient(WithClientNetwork("udp"), WithClientAddress("0.0.0.0:8889"), WithClientEventListener(new(defaultEventListener)))
	client.Connect()
	_ = client
	time.Sleep(time.Second * 10000)
}

//// UDP Server端
//func TestUdpRawServer(t *testing.T) {
//	listen, err := net.ListenUDP("udp", &net.UDPAddr{
//		IP:   net.IPv4(0, 0, 0, 0),
//		Port: 30000,
//	})
//	if err != nil {
//		fmt.Println("Listen failed, err: ", err)
//		return
//	}
//	defer listen.Close()
//	for {
//		var data [1024]byte
//		n, addr, err := listen.ReadFromUDP(data[:]) // 接收数据
//		if err != nil {
//			fmt.Println("read udp failed, err: ", err)
//			continue
//		}
//		fmt.Printf("data:%v addr:%v count:%v\n", string(data[:n]), addr, n)
//		_, err = listen.WriteToUDP(data[:n], addr) // 发送数据
//		if err != nil {
//			fmt.Println("Write to udp failed, err: ", err)
//			continue
//		}
//	}
//}
//
//func TestUdpRawClient(t *testing.T) {
//	addr := &net.UDPAddr{
//		IP:   net.IPv4(0, 0, 0, 0),
//		Port: 8889,
//	}
//	socket, err := net.DialUDP("udp", nil, addr)
//	if err != nil {
//		fmt.Println("连接UDP服务器失败，err: ", err)
//		return
//	}
//	defer socket.Close()
//	sendData := []byte("Hello Server")
//	_, err = socket.Write(sendData) // 发送数据
//	if err != nil {
//		fmt.Println("发送数据失败，err: ", err)
//		return
//	}
//	data := make([]byte, 4096)
//	n, remoteAddr, err := socket.ReadFromUDP(data) // 接收数据
//	if err != nil {
//		fmt.Println("接收数据失败, err: ", err)
//		return
//	}
//	fmt.Printf("recv:%v addr:%v count:%v\n", string(data[:n]), remoteAddr, n)
//}
