package sdkConfig

import (
	"crypto/rsa"
	"github.com/go-tron/crypto/rsaUtil"
)

type SDKConfig struct {
	SignPublicKeyPem string
	SignPublicKey    *rsa.PublicKey
	GateWay          string
	ResponseProperty map[string]string
}

func New(config *SDKConfig) *SDKConfig {
	signPublicKey, err := rsaUtil.GetPublicKeyPem([]byte(config.SignPublicKeyPem))
	if err != nil {
		panic(err)
	}
	config.SignPublicKey = signPublicKey

	return config
}
