// https://signal.org/docs/specifications/doubleratchet/#recommended-cryptographic-algorithms
// NON HEADER ENCYPTED implementation of Double Ratched algorithm
package dr

// TODO: header encryption

import (
	"chat/libsignal/dh"
	"chat/libsignal/header"
	"chat/libsignal/hkdf"
	"chat/shared/errs"
)

const MaximumSkip = 10

type skippedMessageKey struct {
	dhReceiving  [32]byte
	numReceiving int
}

type Ratchet struct {
	dhSendingPair      *dh.KeyPair                  // DH Ratchet key pair (the "sending" or "self" ratchet key)
	dhReceivingKey     []byte                       // DH Ratchet public key (the "received" or "remote" key)
	rootKey            []byte                       // 32-byte Root Key
	chainKeySending    []byte                       // 32-byte Chain Keys for sending
	chainKeyReceiving  []byte                       // 32-byte Chain Keys for receiving
	numSending         int                          // Message numbers for sending
	numReceiving       int                          // Message numbers for receiving
	previousNum        int                          // Number of messages in previous sending chain
	skippedMessageKeys map[skippedMessageKey][]byte // Dictionary of skipped-over message keys, indexed by ratchet public key and message number. Raises an exception if too many elements are stored
}

func NewRatchetFromPublicKey(secret []byte, theirDHPublicKey []byte) (*Ratchet, error) {
	eb := errs.B().Msg("failed to generate ratchet from public key")

	ratchet := Ratchet{}

	pair, err := dh.NewKeyPair()
	if err != nil {
		return nil, eb.Cause(err).Err()
	}

	ratchet.dhSendingPair = pair
	ratchet.dhReceivingKey = theirDHPublicKey

	dhOut, err := dh.DH(pair, theirDHPublicKey)
	if err != nil {
		return nil, eb.Cause(err).Err()
	}

	rootKey, chainKeySending, err := hkdf.KDF_RK(secret, dhOut)
	if err != nil {
		return nil, eb.Cause(err).Err()
	}

	ratchet.rootKey = rootKey
	ratchet.chainKeySending = chainKeySending

	ratchet.numSending = 0
	ratchet.numReceiving = 0
	ratchet.previousNum = 0

	ratchet.skippedMessageKeys = make(map[skippedMessageKey][]byte)

	return &ratchet, nil
}

func NewRatchetFromKeyPair(secret []byte, ourKeyPair *dh.KeyPair) *Ratchet {
	ratchet := Ratchet{}

	ratchet.dhSendingPair = ourKeyPair
	ratchet.dhReceivingKey = nil

	ratchet.rootKey = secret

	ratchet.chainKeySending = nil
	ratchet.chainKeyReceiving = nil

	ratchet.numSending = 0
	ratchet.numReceiving = 0
	ratchet.previousNum = 0

	ratchet.skippedMessageKeys = make(map[skippedMessageKey][]byte)

	return &ratchet
}

func (r *Ratchet) updateState(headers map[string]any) error {
	eb := errs.B().Msg("failed to update state")

	pubKey, ok := headers[header.PublicKeyName].([]byte)
	if !ok {
		return eb.Causef("'%s' is missing from header", header.PublicKeyName).Err()
	}

	r.previousNum = r.numSending
	r.numSending = 0
	r.numReceiving = 0

	r.dhReceivingKey = pubKey

	dhOut, err := dh.DH(r.dhSendingPair, r.dhReceivingKey)
	if err != nil {
		return eb.Cause(err).Err()
	}

	rootKey, chainKeyReceiving, err := hkdf.KDF_RK(r.rootKey, dhOut)
	if err != nil {
		return eb.Cause(err).Err()
	}

	r.rootKey = rootKey
	r.chainKeyReceiving = chainKeyReceiving

	r.dhSendingPair, err = dh.NewKeyPair()
	if err != nil {
		return eb.Cause(err).Err()
	}

	dhOut, err = dh.DH(r.dhSendingPair, r.dhReceivingKey)
	if err != nil {
		return eb.Cause(err).Err()
	}

	rootKey, chainKeySending, err := hkdf.KDF_RK(r.rootKey, dhOut)
	if err != nil {
		return eb.Cause(err).Err()
	}

	r.rootKey = rootKey
	r.chainKeySending = chainKeySending

	return nil
}
