package auth

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"

	"github.com/kakkky/app/domain/errors"
)

func parsePrivateKey(pemData []byte) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode(pemData)
	key, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	privateKey, ok := key.(*rsa.PrivateKey)
	if !ok {
		return nil, errors.New("unknown key type")
	}
	return privateKey, nil
}
func parsePublicKey(pemData []byte) (*rsa.PublicKey, error) {
	block, _ := pem.Decode(pemData)
	key, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	publicKey, _ := key.(*rsa.PublicKey)
	return publicKey, nil
}
