package utils

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"io/ioutil"

	"golang.org/x/crypto/ssh"
)

type keyPair struct {
	key        *rsa.PrivateKey
	PrivateKey []byte
	PublicKey  []byte
}

// NewKeyPair return a new keyPair
func NewKeyPair(bitSize int) (*keyPair, error) {
	var kp keyPair
	var err error

	kp.key, err = rsa.GenerateKey(rand.Reader, bitSize)
	if err != nil {
		return nil, err
	}

	err = kp.key.Validate()
	if err != nil {
		return nil, err
	}

	kp.PrivateKey = kp.marshalPrivateKey()
	kp.PublicKey, err = kp.marshalPublicKey()
	if err != nil {
		return nil, err
	}

	return &kp, nil
}

// WriteToFile save key pair to files
func (kp *keyPair) WriteToFile(filename string) (err error) {
	err = kp.writeFile(filename, kp.PrivateKey)
	if err != nil {
		return
	}
	err = kp.writeFile(filename+".pub", kp.PublicKey)
	if err != nil {
		return
	}
	return nil
}

func (kp *keyPair) marshalPrivateKey() []byte {
	privDer := x509.MarshalPKCS1PrivateKey(kp.key)

	privBlock := pem.Block{
		Type:    "RSA PRIVATE KEY",
		Headers: nil,
		Bytes:   privDer,
	}

	return pem.EncodeToMemory(&privBlock)
}

func (kp *keyPair) marshalPublicKey() ([]byte, error) {
	publicKey, err := ssh.NewPublicKey(&kp.key.PublicKey)
	if err != nil {
		return nil, err
	}

	return ssh.MarshalAuthorizedKey(publicKey), nil
}

func (kp *keyPair) writeFile(filename string, content []byte) error {
	err := ioutil.WriteFile(filename, content, 0600)
	if err != nil {
		return err
	}
	return nil
}
