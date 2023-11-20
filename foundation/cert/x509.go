package cert

// CreateCACert given a CA, keyFilePath, caCertFilePath create a CACert.
func CreateCACert(ca *CACert, keyFilePath, caCertFilePath string) error {
	return nil
}

// CreateCert given a cert, caKey byte slice and caCert byte slice creates a cert.
func CreateCert(cert *Cert, caKey []byte, caCert []byte,
	keyFilePath, certFilePath string) error {
	return nil
}
