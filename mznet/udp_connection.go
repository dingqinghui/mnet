/**
 * @Author: dingQingHui
 * @Description:
 * @File: udp_cpnnection
 * @Version: 1.0.0
 * @Date: 2022/10/28 15:09
 */

package mznet

import (
	"bytes"
	"errors"
	"net"
)

type udpConnection struct {
	IConnection
	readChan chan []byte
	addr     string
	network  string
}

func newUdpConnection(con *net.UDPConn, cType ConnectionType, addr string, network string, eventListener IEventListener) *udpConnection {
	u := &udpConnection{
		IConnection: newConnection(con, cType, eventListener),
		readChan:    make(chan []byte, 64),
		addr:        addr,
		network:     network,
	}
	go u.read()
	go u.write()
	return u
}

func (u *udpConnection) recUdpStream(data []byte) {
	u.readChan <- data
}

func (u *udpConnection) read() {
	for !u.IsClose() {
		select {
		case data, _ := <-u.readChan:
			msg, err := u.GetCodec().UnPack(bytes.NewReader(data))
			if err != nil {
				if !u.GetEventListener().OnError(u, err) {
					return
				}
			}
			if !u.GetEventListener().OnProcess(u, msg) {
				return
			}
		case <-u.Done():
			return
		}
	}
}

func (u *udpConnection) Write(p []byte) (n int, err error) {
	udpAddress, err := net.ResolveUDPAddr(u.network, u.addr)
	if err != nil {
		return 0, err
	}
	udpCon, ok := u.GetCon().(*net.UDPConn)
	if !ok {
		return 0, errors.New("change type fail")
	}
	if u.GetType() == ConnectConnection {
		return udpCon.Write(p)
	} else {
		return udpCon.WriteToUDP(p, udpAddress)
	}
}

func (u *udpConnection) write() {
	for !u.IsClose() {
		select {
		case msg, _ := <-u.GetWriteChan():
			if err := u.GetCodec().Pack(u, msg); err != nil {
				if !u.GetEventListener().OnError(u, err) {
					return
				}
			}
		case <-u.Done():
			return
		}
	}
}
