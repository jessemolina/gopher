package access

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
)

// GeneratePrivateKey creates a random 2048-bit RSA Key.
func GeneratePrivateKey() (*rsa.PrivateKey, error) {
	pk, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, fmt.Errorf("error generating key: %v", err)
	}
	return pk, nil
}

// EncodePrivatePEM encodes a Private Key PEM block into bytes.
func EncodePrivatePEM(pk *rsa.PrivateKey) ([]byte, error) {
	b := x509.MarshalPKCS1PrivateKey(pk)

	pem, err := encodePEM(b, "PRIVATE KEY")
	if err != nil {
		return nil, err
	}

	return pem, nil
}

// EncodePublicKey encodes a Public Key PEM block into bytes from a private key.
func EncodePublicPEM(pk *rsa.PrivateKey) ([]byte, error) {
	b, err := x509.MarshalPKIXPublicKey(&pk.PublicKey)
	if err != nil {
		return nil, fmt.Errorf("Failed to decode pub key: %v", err)
	}

	pem, err := encodePEM(b, "PUBLIC KEY")
	if err != nil {
		return nil, err
	}

	return pem, nil
}

// encodePEM encodes a PEM block into bytes.
func encodePEM(b []byte, t string) ([]byte, error) {
	buf := new(bytes.Buffer)
	block := pem.Block{
		Type:  t,
		Bytes: b,
	}

	if err := pem.Encode(buf, &block); err != nil {
		return nil, fmt.Errorf("error formatting pem block: %v", err)
	}

	return buf.Bytes(), nil
}
