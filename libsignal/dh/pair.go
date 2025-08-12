package dh

import (
	"chat/shared"
	"crypto/ecdh"
	"crypto/rand"
)

type KeyPair struct {
	public  []byte
	private []byte
}

func NewKeyPair() (*KeyPair, error) {
	eb := shared.B().Msg("failed to generate dh pair")
	curve := ecdh.X25519()

	privateKey, err := curve.GenerateKey(rand.Reader)
	if err != nil {
		return nil, eb.Cause(err).Err()
	}

	return &KeyPair{
		public:  privateKey.PublicKey().Bytes(),
		private: privateKey.Bytes(),
	}, nil
}

func (pair KeyPair) Public() []byte {
	return pair.public
}

func (pair KeyPair) Private() []byte {
	return pair.private
}
