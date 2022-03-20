package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncryptDecrypt(t *testing.T) {
	data := []byte("data")
	iv := generateIV()
	key := generateKey()

	encrypted := encrypt(data, iv, key)
	decrypted := decrypt(encrypted, iv, key)

	assert.Equal(t, data, decrypted)
}

func TestEncryptDecryptStr(t *testing.T) {
	data := "data"

	iv := generateIV()
	key := generateKey()

	ivStr := exportB58(iv)
	keyStr := exportB58(key)

	dataBytes := []byte(data)

	assert.Equal(t, data, string(dataBytes))

	encrypted := encrypt(dataBytes, iv, key)
	encryptedStr := exportB58(encrypted)

	assert.Equal(t, iv, importB58(ivStr))
	assert.Equal(t, key, importB58(keyStr))
	assert.Equal(t, encrypted, importB58(encryptedStr))

	decrypted := decrypt(encrypted, iv, key)

	assert.Equal(t, string(decrypted), data)
}
