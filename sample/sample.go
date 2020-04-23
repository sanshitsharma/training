package main

import "fmt"

type data struct {
	a int
	b int
}

func AFunc(data <--chan int)

func main() {

	c := new(data)
	fmt.Printf("%T", c)
	fmt.Printf("%d", c.b)
}
