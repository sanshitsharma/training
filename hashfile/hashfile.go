package main

import (
	"crypto/sha256"
	"fmt"
	"io"
	"os"
)

func main() {
	f, err := os.Open("file.txt")
	if err != nil {
		fmt.Println("error")
		return
	}
	defer f.Close()
	h := sha256.New()

	io.Copy(h, f)
	fmt.Printf("%x", h.Sum(nil))

}
