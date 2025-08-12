package dr

import (
	"chat/libsignal/aecd"
	"chat/libsignal/header"
	"chat/libsignal/hkdf"
	"chat/shared"
)

func (r *Ratchet) Encrypt(plaintext []byte, associatedData []byte) (map[string]any, []byte, error) {
	eb := shared.B().Msg("failed to ratchet encrypt")

	chainKeySending, messageKey := hkdf.KDF_CK(r.chainKeySending)
	r.chainKeySending = chainKeySending
	head := header.New(r.dhSendingPair, r.previousNum, r.numSending)
	r.numSending += 1

	cyphertext, err := aecd.Encrypt(messageKey, plaintext, header.Concat(associatedData, head))
	if err != nil {
		return nil, nil, eb.Cause(err).Err()
	}

	return head, cyphertext, nil
}
