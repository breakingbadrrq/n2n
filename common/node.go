package common

type NodeDesc struct {
	RemoteDomainIP string
	MacAddr        string
	ExtraMessage   interface{}
}

type Node struct {
	Name string

	HashIP string
	Desc   NodeDesc
	IP     P2P_IP

	publicKey  string
	privateKey string
}

func isCluster(node1 Node, node2 Node) bool {
	return node1.Name == node2.Name
}

func genSecretKey(nonce string) (string, string) {

}

func (node Node) genHashIP() string {
	jsonStr := json.Marshal(service.Desc)

	sha := sha256.New()
	sha.Write([]byte(jsonStr))

	return base64.StdEncoding.EncodeToString(sha.Sum(nil))
}
