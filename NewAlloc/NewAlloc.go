package main

import "fmt"

type Data struct {
	a int
	b int
	c float32
}

func main() {

	a := new(Data)
	fmt.Println(a.c)
	fmt.Printf("%T", a)
}
