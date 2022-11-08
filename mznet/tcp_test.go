/**
 * @Author: dingQingHui
 * @Description:
 * @File: tcp_test
 * @Version: 1.0.0
 * @Date: 2022/10/28 14:39
 */

package mznet

type defaultEventListener struct{}

func (defaultEventListener) OnConnected(connection IConnection) bool {
	connection.SetEventListener(new(defaultEventListener))

	if connection.GetType() == ConnectConnection {
		data := []byte("11111111222222\n")
		connection.Send(data)
	} else {
		//_ = connection.Close()
	}
	println("implement me OnConnected")
	return true
}

func (defaultEventListener) OnProcess(connection IConnection, msg interface{}) bool {
	println("implement me OnProcess", msg)
	data := msg.(string)
	connection.Send([]byte(data))
	return true
}

func (defaultEventListener) OnClosed(IConnection) bool {
	println("implement me OnClosed")
	return true
}

func (defaultEventListener) OnError(connection IConnection, err error, users ...interface{}) bool {
	println("implement me OnError", err.Error())
	if connection != nil {
		connection.Close()
	}
	return false
}

//func TestNewTcpServer(t *testing.T) {
//	serve := NewTcpServer(WithServerNetwork("tcp"), WithServerAddress("127.0.0.1:50000"), WithServerEventListener(new(defaultEventListener)))
//	serve.RunEventLoop()
//
//	time.Sleep(time.Second * 10000)
//}
//
//func TestNewTcpClient(t *testing.T) {
//	client := newTcpClient(WithClientAddress("tcp"), WithClientAddress("127.0.0.1:50000"), WithClientEventListener(new(defaultEventListener)))
//	client.Connect()
//	_ = client
//	time.Sleep(time.Second * 10000)
//}
