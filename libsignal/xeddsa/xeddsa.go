package xeddsa

import (
	"crypto/ed25519"
	"crypto/sha512"
	"fmt"
)

func clampScalar(secret *[32]byte) [32]byte {
	var s [32]byte
	copy(s[:], secret[:])
	s[0] &= 248
	s[31] &= 127
	s[31] |= 64
	return s
}

func Sign(x25519PrivateKey, message []byte) ([]byte, error) {
	if len(x25519PrivateKey) != 32 {
		return nil, fmt.Errorf("invalid private key length")
	}

	clamped := clampScalar((*[32]byte)(x25519PrivateKey))

	h := sha512.New()
	h.Write(clamped[:])
	digest := h.Sum(nil)

	edPrivateKey := ed25519.NewKeyFromSeed(digest[:32])
	sig := ed25519.Sign(edPrivateKey, message)

	return sig, nil
}
