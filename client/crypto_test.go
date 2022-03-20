package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExportIV(t *testing.T) {
	assert.Equal(t, "3yZe7d", exportIV([]byte("test")))
}

func TextEncryptDecrypt(t *testing.T) {
	data := []byte("data")
	iv := generateIV()
	key := generateKey()

	encrypted := encrypt(data, iv, key)
	decrypted := decrypt(encrypted, iv, key)

	assert.Equal(t, data, decrypted)
}
