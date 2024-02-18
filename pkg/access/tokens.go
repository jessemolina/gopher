package access

import (
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// GenerateToken provides an rsa token with the given claim set and roles.
func GenerateToken(pk *rsa.PrivateKey, rc jwt.RegisteredClaims, roles []string) ([]byte, error) {
	claims := struct {
		jwt.RegisteredClaims
		Roles []string
	}{
		rc,
		roles,
	}

	method := jwt.GetSigningMethod(jwt.SigningMethodRS256.Name)

	kid, err := PublicKeyID(pk)
	if err != nil {
		return nil, fmt.Errorf("Failed to generate KID: %v", err)
	}

	token := jwt.NewWithClaims(method, claims)
	token.Header["kid"] = kid

	signed, err := token.SignedString(pk)
	if err != nil {
		return nil, fmt.Errorf("Failed to sign token: %v", err)
	}

	return []byte(signed), nil
}

// PubKeyID generates a KID from a private key's public key.
func PublicKeyID(pk *rsa.PrivateKey) (string, error) {
	pub, err := x509.MarshalPKIXPublicKey(&pk.PublicKey)
	if err != nil {
		return "", err
	}

	hash := sha256.Sum256(pub)

	size := len(hash) / 2

	kid := hex.EncodeToString(hash[:size])

	return kid, nil
}

// ClaimSets a JWT registered claim set based with the given lifespan.
func MakeClaimSet(subject string, issuer string, lifespan time.Duration) jwt.RegisteredClaims {
	return jwt.RegisteredClaims{
		Subject:   subject,
		Issuer:    issuer,
		IssuedAt:  jwt.NewNumericDate(time.Now().UTC()),
		ExpiresAt: jwt.NewNumericDate(time.Now().UTC().Add(lifespan)),
	}
}
