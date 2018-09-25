package common

type DeviceServiceDesc struct {
	RemoteDomainIP string
	MacAddr        string
	ExtraMessage   interface{}
}

type DeviceService struct {
	Name    string
	PinCode string

	HashIP string
	Desc   DeviceServiceDesc
	IP     P2P_IP

	publicKey  string
	privateKey string
}

func isCluster(service1 DeviceServiceDesc, service2 DeviceService) bool {
	return service1.Name == service2.Name
}

func genSecretKey(nonce string) (string, string) {

}

func (service DeviceService) genHashIP() string {
	jsonStr := json.Marshal(service.Desc)

	sha := sha256.New()
	sha.Write([]byte(jsonStr))

	return base64.StdEncoding.EncodeToString(sha.Sum(nil))
}
