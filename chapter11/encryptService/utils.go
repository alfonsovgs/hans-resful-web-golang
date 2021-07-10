package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

// Initialization vector fo the AES algorithm
var initVector = []byte{35, 46, 57, 24, 85, 35, 24, 74, 87, 35, 88, 98, 66, 32, 14, 05}

// EncriptString encrypts the string withh given key
func EncryptString(key, text string) string {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		panic(err)
	}

	plainText := []byte(text)
	cfb := cipher.NewCFBEncrypter(block, initVector)
	cipherText := make([]byte, len(plainText))
	cfb.XORKeyStream(cipherText, plainText)
	return base64.StdEncoding.EncodeToString(cipherText)
}

// DecryptString decrypts the encrypted string to original
func DecryptString(key, text string) string {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		panic(err)
	}

	cipherText, _ := base64.StdEncoding.DecodeString(text)
	cfb := cipher.NewCFBEncrypter(block, initVector)
	plainText := make([]byte, len(cipherText))
	cfb.XORKeyStream(plainText, cipherText)

	return string(plainText)
}
