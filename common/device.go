package common

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"fmt"
)

type DevDesc struct {
	Name string

	RemoteDomainIP   string
	MacAddr          string
	Producer         string
	ProduceAddr      string
	SerialNumber     string
	PartNumber       string
	MadeIn           string
	ProductCatagory  string
	ProductionDate   string
	DurabilityPeriod string
	ExtraMessage     interface{}
}

type Device struct {
	HashIP string

	Desc    DevDesc
	PinCode string
	IP      P2P_IP

	publicKey  string
	privateKey string
}

func (dev Device) genHashIP() string {
	jsonStr := json.Marshal(dev.Desc)
	sha := sha256.New()
	sha.Write([]byte(jsonStr))
	sha.Write([]byte(dev.PinCode))

	return base64.StdEncoding.EncodeToString(sha.Sum(nil))
}

func (dev Device) genSecretKey(nonce string) (string, string) { //private key, public key
}
