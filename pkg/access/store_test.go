package access

import (
	"testing"
	"time"
)

// TODO Fix test for Writing keys/tokens to file.
func TestWriteToFile(t *testing.T) {
	pk, _ := GeneratePrivateKey()

	pvt, _ := EncodePrivatePEM(pk)
	pub, _ := EncodePublicPEM(pk)

	rc := MakeClaimSet("Tester", "Gopher Access", time.Hour)
	roles := []string{"ADMIN"}
	tk, _ := GenerateToken(pk, rc, roles)

	WriteToFile(pvt, "private.pem")
	WriteToFile(pub, "public.pem")
	WriteToFile([]byte(tk), "token")
}
