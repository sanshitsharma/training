package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
)

func main() {
	key := []byte("1234123412341234")
	plaintext := []byte("Hello, world")

	block, err := aes.NewCipher(key)

	if err != nil {
		fmt.Println("error")
		return
	}
	nonce := make([]byte, 12)
	io.ReadFull(rand.Reader, nonce)

	fmt.Printf("nonce: %x\n", nonce)

	aesgcm, _ := cipher.NewGCM(block)

	ciphertext := aesgcm.Seal(nil, nonce, plaintext, nil)
	fmt.Printf("cipher:%xn", ciphertext)

	// Decrypt

	dplaintext, err2 := aesgcm.Open(nil, nonce, ciphertext, nil)

	if err2 == nil {
		fmt.Printf("%s\n", dplaintext)
	}
}
