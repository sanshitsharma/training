package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
)

func main() {

	privateKey, _ := rsa.GenerateKey(rand.Reader, 2048)
	publicKey := &privateKey.PublicKey

	plaintext := []byte("Hello, world")
	label := []byte("")
	hash := sha256.New()

	// Encrypt

	ciphertext, _ := rsa.EncryptOAEP(
		hash,
		rand.Reader,
		publicKey,
		plaintext,
		label)

	// Decrypt

	plainText, _ := rsa.DecryptOAEP(
		hash,
		rand.Reader,
		privateKey,
		ciphertext,
		label)

	fmt.Printf("Decrypted Text\n[%s]\n", plainText)

}
