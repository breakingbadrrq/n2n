package p2p

const (
	PEER_N2N = iota
	PEER_N2D
	PEER_D2D
)

type Peer struct {
	HashIPSrc string
	HashIPDes string

	PeerType    uint8
	MessagePipe map[string]Pipe
}
