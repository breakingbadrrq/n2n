package main

import "fmt"

const (
	UDP string = "udp"
	TCP string = "tcp"
)

type PingPong interface { //物理加密通道
	Ping() error
	Pong() error
	Exchange() error

	WaitLoop() error
}

type PipeProtocal interface { // 建立 P2P 通道（通过交换 NAT信息）
	HandShake() error
	WaitLoop() error
}

type Pipe struct {
	Name string

	SendChan chan []byte
	RecvChan chan []byte

	protocal PipeProtocal
}
