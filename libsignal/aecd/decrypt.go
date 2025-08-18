package aecd

import (
	"chat/shared/errs"
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"io"

	"golang.org/x/crypto/hkdf"
)

func Decrypt(messageKey, ciphertextWithTag, associatedData []byte) ([]byte, error) {
	eb := errs.B().Msg("failed to aecd decrypt")

	if len(ciphertextWithTag) < 96 {
		return nil, eb.Causef("ciphertext too short").Err()
	}

	salt := make([]byte, 80)
	reader := hkdf.New(hash, messageKey, salt, encryptInfo)

	derivedKey := make([]byte, 80)
	_, err := io.ReadFull(reader, derivedKey)
	if err != nil {
		return nil, eb.Cause(err).Err()
	}

	encryptionKey := derivedKey[:32]
	authenticationKey := derivedKey[32:64]
	iv := derivedKey[64:80]

	tagStart := len(ciphertextWithTag) - hash().Size()
	tag := ciphertextWithTag[tagStart:]
	ivFromMessage := ciphertextWithTag[:len(iv)]
	ciphertext := ciphertextWithTag[len(iv):tagStart]

	// Verify MAC
	mac := hmac.New(hash, authenticationKey)
	mac.Write(associatedData)
	mac.Write(ivFromMessage)
	mac.Write(ciphertext)
	expectedTag := mac.Sum(nil)

	if !hmac.Equal(tag, expectedTag) {
		return nil, eb.Causef("authentication failed").Err()
	}

	// Decrypt
	block, err := aes.NewCipher(encryptionKey)
	if err != nil {
		return nil, eb.Cause(err).Err()
	}

	if len(ciphertext)%aes.BlockSize != 0 {
		return nil, eb.Causef("ciphertext is not a multiple of the block size").Err()
	}

	mode := cipher.NewCBCDecrypter(block, ivFromMessage)
	paddedPlaintext := make([]byte, len(ciphertext))
	mode.CryptBlocks(paddedPlaintext, ciphertext)

	// Remove padding
	paddingLen := int(paddedPlaintext[len(paddedPlaintext)-1])
	if paddingLen > aes.BlockSize || paddingLen == 0 {
		return nil, eb.Causef("invalid padding").Err()
	}

	for _, b := range paddedPlaintext[len(paddedPlaintext)-paddingLen:] {
		if int(b) != paddingLen {
			return nil, eb.Causef("invalid padding").Err()
		}
	}

	plaintext := paddedPlaintext[:len(paddedPlaintext)-paddingLen]
	return plaintext, nil
}
