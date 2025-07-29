package main

import (
	"chat/libsignal/dh"
	"chat/libsignal/dr"
	"crypto/rand"
	"fmt"
	"os"
)

func main() {
	secret := make([]byte, 32)
	_, err := rand.Read(secret)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	ourKeyPair, err := dh.NewKeyPair()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	ourRatchet := dr.NewRatchetFromKeyPair(secret, ourKeyPair)

	theirRatchet, err := dr.NewRatchetFromPublicKey(secret, ourKeyPair.Public())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	plaintext := []byte("plaintext")
	associatedData := []byte("")
	headers, cyphertext, err := ourRatchet.Encrypt(plaintext, associatedData)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	plaintext2, err := theirRatchet.Decrypt(headers, cyphertext, associatedData)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(string(plaintext2))
}
