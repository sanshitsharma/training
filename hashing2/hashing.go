package main

import "fmt"
import "crypto/md5"
import "crypto/sha1"
import "crypto/sha256"

func main() {
	s := "Foo"

	md5 := md5.Sum([]byte(s))
	sha1 := sha1.Sum([]byte(s))
	sha256 := sha256.Sum256([]byte(s))

	fmt.Printf("%x\n", md5)
	fmt.Printf("%x\n", sha1)
	fmt.Printf("%x\n", sha256)
}
