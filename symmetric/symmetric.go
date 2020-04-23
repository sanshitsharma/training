package main

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
)

var commonIV = []byte{0x00, 0x01, 0x02, 0x03, 0x04,
	0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c,
	0x0d, 0x0e, 0x0f}

func main() {
	plaintext := []byte("Hello, world")

	keyValue := "12345678901234567890123456789012"

	// 32byte key returns AES-256 block cipher
	cipherBlock, _ := aes.NewCipher([]byte(keyValue))

	cfb := cipher.NewCFBEncrypter(cipherBlock, commonIV)
	ciphertext := make([]byte, len(plaintext))
	cfb.XORKeyStream(ciphertext, plaintext)

	// Decrypt strings
	cfbdec := cipher.NewCFBDecrypter(cipherBlock, commonIV)
	plaintextCopy := make([]byte, len(plaintext))
	cfbdec.XORKeyStream(plaintextCopy, ciphertext)
	fmt.Printf("%x=>%s\n", ciphertext, plaintextCopy)
}
