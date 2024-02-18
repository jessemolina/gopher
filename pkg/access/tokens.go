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
func GenerateToken(pk *rsa.PrivateKey, rc jwt.RegisteredClaims, roles []string) (string, error) {
	claims := struct {
		jwt.RegisteredClaims
		Roles []string
	}{
		rc,
		roles,
	}

	method := jwt.GetSigningMethod(jwt.SigningMethodRS256.Name)

	kid, err := makePublicKid(pk)
	if err != nil {
		return "", fmt.Errorf("failed to generate KID: %v", err)
	}

	token := jwt.NewWithClaims(method, claims)
	token.Header["kid"] = kid

	signed, err := token.SignedString(pk)
	if err != nil {
		return "", fmt.Errorf("failed to sign token: %v", err)
	}

	return signed, nil
}

// ValidateToken verifies a token's signature via an RSA public key.
func ValidateToken(token string, pub *rsa.PublicKey) error {
	parser := jwt.NewParser(jwt.WithValidMethods([]string{jwt.SigningMethodRS256.Name}))
	keyFunc := func(t *jwt.Token) (interface{}, error) {
		return pub, nil
	}

	var claims struct {
		jwt.RegisteredClaims
		Roles []string
	}

	tk, err := parser.ParseWithClaims(token, &claims, keyFunc)
	if err != nil {
		return fmt.Errorf("failed to parse token with claims: %v", err)
	}

	if !tk.Valid {
		return fmt.Errorf("failed to validate token: %v", err)
	}

	return nil
}

// PubKeyID generates a KID from a private key's public key.
func makePublicKid(pk *rsa.PrivateKey) (string, error) {
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
