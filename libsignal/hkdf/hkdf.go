package hkdf

import (
	"chat/shared/errs"
	"crypto/hmac"
	"crypto/sha512"
	"io"

	"golang.org/x/crypto/hkdf"
)

var hash = sha512.New

var rootKeyKDFInfo = []byte("KDF_RK")

func KDF_RK(rootKey, dhOut []byte) ([]byte, []byte, error) {
	eb := errs.B().Msg("failed to perform KDF_RK")

	reader := hkdf.New(hash, dhOut, rootKey, rootKeyKDFInfo)

	output := make([]byte, 64)
	_, err := io.ReadFull(reader, output)
	if err != nil {
		return nil, nil, eb.Cause(err).Err()
	}

	newRootKey := output[:32]
	newChainKey := output[32:]
	return newRootKey, newChainKey, nil
}

var chainKeyKDFInfo = []byte("KDF_CK")

var messageKeyConstant = []byte{0x01}
var chainKeyConstant = []byte{0x02}

func KDF_CK(chainKey []byte) ([]byte, []byte) {
	hmac1 := hmac.New(hash, chainKey)
	hmac1.Write(messageKeyConstant)
	messageKey := hmac1.Sum(nil)

	hmac2 := hmac.New(hash, chainKey)
	hmac2.Write(chainKeyConstant)
	nextChainKey := hmac2.Sum(nil)

	return nextChainKey, messageKey
}

var kdfInfo = []byte("x3dh")

func KDF(key []byte) ([]byte, error) {
	eb := errs.B().Msg("failed to perform KDF")

	var bytes []byte = make([]byte, 32)
	for i := range bytes {
		bytes[i] = 0xFF
	}

	bytes = append(bytes, key...)

	salt := make([]byte, 32)

	reader := hkdf.New(hash, bytes, salt, kdfInfo)

	output := make([]byte, 32)
	_, err := io.ReadFull(reader, output)
	if err != nil {
		return nil, eb.Cause(err).Err()
	}

	return output, nil
}
