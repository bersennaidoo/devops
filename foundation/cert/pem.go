package cert

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
)

// PemToX509 given input byte slice representing a PEM block creates a x509 certificate.
func PemToX509(input []byte) (*x509.Certificate, error) {
	block, _ := pem.Decode(input)
	if block == nil {
		return nil, fmt.Errorf("failed to parse certificate PEM")
	}
	return x509.ParseCertificate(block.Bytes)
}
