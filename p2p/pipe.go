package main

import "fmt"

const (
	UDP string = "udp"
	TCP string = "tcp"
)

type PingPong interface { //物理通道, 注册NAT信息
	Ping() error
	Pong() error

	Exchange() error
	WaitLoop() error
}

type PipeProtocol interface { // 建立 P2P 通道
	HandShake() error
	WaitLoop() error
}

type Pipe struct {
	Name string

	SendChan chan []byte
	RecvChan chan []byte

	protocol PipeProtocol
}
