package main

import "fmt"

func main() {
	// "こんにちは"

	hellostr := "Hello, world!"

	hellostrBytes := []byte(hellostr)

	hellostrRunes := []rune(hellostr)

	fmt.Println(hellostr)
	fmt.Println(hellostrBytes)
	fmt.Println(hellostrRunes)
}
