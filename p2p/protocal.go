package main

import "fmt"

const (
	UDP string = "udp"
	TCP string = "tcp"
)

type protocal struct {
	Name string

	PeerInfo  func() interface{}
	HandShake func() interface{}

	Send    func(val interface{}) (bool, error)
	Receive func(val interface{}) (bool, error)
}
