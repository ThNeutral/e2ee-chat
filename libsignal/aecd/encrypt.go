package aecd

import (
	"bytes"
	"chat/shared"
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/sha512"
	"io"

	"golang.org/x/crypto/hkdf"
)

var hash = sha512.New

var encryptInfo = []byte("encrypt")

func Encrypt(messageKey, plaintext, associatedData []byte) ([]byte, error) {
	eb := shared.NewErrorBuilder().Msg("failed to aecd encrypt")

	salt := make([]byte, 80)
	reader := hkdf.New(hash, messageKey, salt, encryptInfo)

	derivedKey := make([]byte, 80)
	_, err := io.ReadFull(reader, derivedKey)
	if err != nil {
		return nil, eb.Cause(err).Err()
	}

	encyptionKey := derivedKey[:32]
	authenticationKey := derivedKey[32:64]
	iv := derivedKey[64:]

	block, err := aes.NewCipher(encyptionKey)
	if err != nil {
		return nil, eb.Cause(err).Err()
	}

	padding := aes.BlockSize - len(plaintext)%aes.BlockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	paddedPlaintext := append(plaintext, padtext...)

	ciphertext := make([]byte, len(paddedPlaintext))
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext, paddedPlaintext)

	mac := hmac.New(hash, authenticationKey)
	mac.Write(associatedData)
	mac.Write(iv)
	mac.Write(ciphertext)
	tag := mac.Sum(nil)

	final := append(iv, ciphertext...)
	final = append(final, tag...)

	return final, nil
}
