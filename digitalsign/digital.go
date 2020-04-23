package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
)

func main() {

	priv, _ := rsa.GenerateKey(rand.Reader, 1024)
	//pub:=priv.Public()

	plaintext := []byte("Hello, world")

	hashed := sha256.Sum256(plaintext)

	signature, _ := rsa.SignPKCS1v15(rand.Reader, priv, crypto.SHA256,
		hashed[:])

	// verify

	okay := rsa.VerifyPKCS1v15(&priv.PublicKey, crypto.SHA256, hashed[:], signature)

	if okay == nil {
		fmt.Println("Signature verified!")
	} else {
		fmt.Println("Signature failed!")
	}
}
