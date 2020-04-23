package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
)

func main() {

	priv, err := rsa.GenerateKey(rand.Reader, 1024)

	pub := priv.PublicKey
	message := []byte("hello, world")

	hashed := sha256.Sum256(message)

	signature, err := rsa.SignPKCS1v15(rand.Reader,
		priv, crypto.SHA256, hashed[:])
	if err != nil {
		fmt.Printf("Error from signing: %s\n", err)
		return
	}

	err = rsa.VerifyPKCS1v15(&pub, crypto.SHA256,
		hashed[:], signature)
	if err != nil {
		fmt.Printf("Error from verification: %s\n", err)
		return
	} else {
		fmt.Printf("signature is verified\n")
	}

}
