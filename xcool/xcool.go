package main

import "fmt"

type XCool struct {
	a      string
	values []int
}

func main() {
	obj := XCool{"My name", []int{1, 2, 3, 4, 5}}
	fmt.Println(obj)
}
