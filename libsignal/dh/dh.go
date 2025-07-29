package dh

import (
	"chat/shared"

	"golang.org/x/crypto/curve25519"
)

func DH(pair *KeyPair, pub []byte) ([]byte, error) {
	eb := shared.NewErrorBuilder().Msg("failed to perform DH")

	secret, err := curve25519.X25519(pair.private, pub)
	if err != nil {
		return nil, eb.Cause(err).Err()
	}

	return secret, nil
}
