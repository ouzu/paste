package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"

	"github.com/btcsuite/btcutil/base58"
)

func generateIV() []byte {
	iv := make([]byte, 16)

	_, err := rand.Read(iv)
	if err != nil {
		panic(err)
	}

	return iv
}

func exportIV(iv []byte) string {
	return base58.Encode(iv)
}

func generateKey() []byte {
	key := make([]byte, 256)

	_, err := rand.Read(key)
	if err != nil {
		panic(err)
	}

	return key
}

func encrypt(data []byte, iv []byte, key []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err)
	}

	encrypted := gcm.Seal(iv, iv, data, nil)
	return encrypted
}

func decrypt(encrypted []byte, iv []byte, key []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err)
	}

	data, err := gcm.Open(nil, iv, encrypted, nil)
	if err != nil {
		panic(err)
	}

	return data
}
