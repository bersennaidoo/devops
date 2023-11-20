// Package key Creates, saves RSA private key, maps to and from RSA to PEM.
package key

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

// CreateRSAPrivateKey given n bits creates a random rsa private key.
func CreateRSAPrivateKey(n int) (*rsa.PrivateKey, error) {
	return rsa.GenerateKey(rand.Reader, n)
}

// RSAPrivateKeyToPEM given a rsa private key, returns a PEM block with a private key encoded as PKCS.
func RSAPrivateKeyToPEM(privateKey *rsa.PrivateKey) *pem.Block {
	return &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	}
}

// CreateRSAPrivateKeyAndSave given a path and number of bits, creates and
// saves a PEM encoded Block
func CreateRSAPrivateKeyAndSave(path string, n int) error {
	privateKey, err := CreateRSAPrivateKey(n)
	if err != nil {
		return err
	}
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0600)
	if err != nil {
		return err
	}
	if err := pem.Encode(f, RSAPrivateKeyToPEM(privateKey)); err != nil {
		return err
	}
	if err := f.Close(); err != nil {
		return err
	}
	return nil
}

// PrivateKeyPemToRSA given a input byte slice representing PEM Block
// Creates RSA key.
func PrivateKeyPemToRSA(input []byte) (*rsa.PrivateKey, error) {
	var parsedKey *rsa.PrivateKey
	var err error

	privPem, _ := pem.Decode(input)

	if privPem.Type != "RSA PRIVATE KEY" {
		return nil, fmt.Errorf("RSA private key is of the wrong type: %s",
			privPem.Type)
	}
	if parsedKey, err = x509.ParsePKCS1PrivateKey(privPem.Bytes); err != nil {
		return nil, fmt.Errorf("Unable to parse RSA private key: %v", err)
	}
	return parsedKey, nil
}
