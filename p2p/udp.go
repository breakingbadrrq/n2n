package p2p

import (
	"fmt"

	"lightningdog-app.com/common"
)

const (
	P2P_UDP_PORT = "6025"
	P2P_UDP_HOST = "172.0.0.1"
)

func Loop() {
	loopAddr := P2P_UDP_HOST + ":" + P2P_UDP_PORT
	udpAddr, err = net.ResolveUDPAddr("udp4", loopAddr)
	if err != nil {
		fmt.Println("fail to listen local address")
	}

	conn, err := net.ListenUDP("udp", loopAddr)
	if err != nil {
		fmt.Println("fail to listen local address")
		os.Exit(1)
	}

	defer conn.Close()
	buf := make([512]byte)
	for {
		_, remoteAddr, err := conn.ReadFromUDP(buf)
		if err != nil {
			continue
		}

		go handleData(&remoteAddr, buf, len)
	}
}

func handleData(remoteAddr *net.UDPAddr, buf []byte) {

}
