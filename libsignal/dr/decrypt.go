package dr

import (
	"chat/libsignal/aecd"
	"chat/libsignal/header"
	"chat/libsignal/hkdf"
	"chat/shared"
	"slices"
)

func (r *Ratchet) Decrypt(headers map[string]any, cyphertext []byte, associatedData []byte) ([]byte, error) {
	eb := shared.NewErrorBuilder().Msg("failed to ratched decrypt")

	plaintext, err := r.trySkippedMessageKeys(headers, cyphertext, associatedData)
	if err != nil {
		return nil, eb.Cause(err).Err()
	}

	if plaintext != nil {
		return plaintext, nil
	}

	pubKey, ok := headers[header.PublicKeyName].([]byte)
	if !ok {
		return nil, eb.Causef("'%s' is missing from header", header.PublicKeyName).Err()
	}

	pn, ok := headers[header.PreviousChainLengthName].(int)
	if !ok {
		return nil, eb.Causef("'%s' is missing from header", header.PreviousChainLengthName).Err()
	}

	n, ok := headers[header.MessageNumber].(int)
	if !ok {
		return nil, eb.Causef("'%s' is missing from header", header.MessageNumber).Err()
	}

	if !slices.Equal(pubKey, r.dhReceivingKey) {
		err := r.skipMessageKeys(pn)
		if err != nil {
			return nil, eb.Cause(err).Err()
		}

		err = r.updateState(headers)
		if err != nil {
			return nil, eb.Cause(err).Err()
		}
	}

	err = r.skipMessageKeys(n)
	if err != nil {
		return nil, eb.Cause(err).Err()
	}

	chainKeyReceiving, messageKey := hkdf.KDF_CK(r.chainKeyReceiving)
	r.chainKeyReceiving = chainKeyReceiving
	r.numReceiving += 1

	plaintext, err = aecd.Decrypt(messageKey, cyphertext, header.Concat(associatedData, headers))
	if err != nil {
		return nil, eb.Cause(err).Err()
	}

	return plaintext, nil
}

func (r *Ratchet) trySkippedMessageKeys(headers map[string]any, cyphertext []byte, associatedData []byte) ([]byte, error) {
	eb := shared.NewErrorBuilder().Msg("failed to try skipped message keys")

	pubKey, ok := headers[header.PublicKeyName].([]byte)
	if !ok {
		return nil, eb.Causef("'%s' is missing from header", header.PublicKeyName).Err()
	}

	num, ok := headers[header.MessageNumber].(int)
	if !ok {
		return nil, eb.Causef("'%s' is missing from header", header.MessageNumber).Err()
	}

	skipped := skippedMessageKey{
		dhReceiving:  [32]byte(pubKey),
		numReceiving: num,
	}

	messageKey, ok := r.skippedMessageKeys[skipped]
	if !ok {
		return nil, nil
	}

	plaintext, err := aecd.Decrypt(messageKey, cyphertext, associatedData)
	if err != nil {
		return nil, eb.Cause(err).Err()
	}

	return plaintext, nil
}

func (r *Ratchet) skipMessageKeys(until int) error {
	eb := shared.NewErrorBuilder().Msg("failed to skip message keys")

	if r.numReceiving+MaximumSkip < until {
		return eb.Causef("tried skip too much").Err()
	}

	if r.chainKeyReceiving != nil {
		for r.numReceiving < until {
			chainKeyReceiving, messageKey := hkdf.KDF_CK(r.chainKeyReceiving)
			r.chainKeyReceiving = chainKeyReceiving

			skipped := skippedMessageKey{
				dhReceiving:  [32]byte(r.dhReceivingKey),
				numReceiving: r.numReceiving,
			}
			r.skippedMessageKeys[skipped] = messageKey

			r.numReceiving += 1
		}
	}

	return nil
}
