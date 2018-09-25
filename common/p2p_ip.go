package common

const (
	OTHER_CONE_NAT = iota
	FULL_CONE_NAT
	RESTRICTED_CONE_NAT
	PORT_RESTRICTED_CONE_NAT
	SYMMETRIC_NAT
)

type P2P_IP struct {
	LocalIP     uint64
	PublicIP    uint64
	PublicPort  uint64
	CoreNATType uint8
}

func (ip P2P_IP) isPublic() bool {
	if ip.LocalIP == ip.PublicIP {
		return true
	}

	return false
}
