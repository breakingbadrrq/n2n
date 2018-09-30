package p2p

import (
	"fmt"
	"lightningdog-app.com/common"
	"net"
)

const (
	P2P_UDP_PORT     = "6024"
	P2P_UDP_HOST     = "172.0.0.1"
	P2P_UDP_BUF_SIZE = 1280
)

type (
	Ping struct {
		Version  uint
		From, To EndPoint

		Expiration uint64
		Rest       rlp.RawValue `rpl:"tail"`
	}

	Pong struct {
		To      EndPoint
		ReplyOk []byte

		Expiration uint64
		Rest       rlp.RawValue `rpl:"tail"`
	}

	EndPoint struct {
		HashIP []byte
		Desc   interface{}

		LocalIP uint64
	}
)

func Loop() {
	udpAddr, err = genUDPAddr()
	if err != nil {
		fmt.Println("fail to listen local address")
	}

	conn, err := net.ListenUDP("udp", loopAddr)
	if err != nil {
		fmt.Println("fail to listen local address")
		os.Exit(1)
	}

	defer conn.Close()
	buf := make([P2P_UDP_BUF_SIZE]byte)
	for {
		_, remoteAddr, err := conn.ReadFromUDP(buf)
		if err != nil {
			continue
		}

		go handleData(&remoteAddr, buf, len)
	}
}

func genUDPAddr() (net.UDPAddr, error) {
	loopAddr := P2P_UDP_HOST + ":" + P2P_UDP_PORT
	return net.ResolveUDPAddr("udp4", loopAddr)
}

func handleData(remoteAddr *net.UDPAddr, buf []byte) {

}
