package main

import "fmt"

type Data struct {
	a int
	b int
	c float32
}

func DoSomething(d *Data) {
	fmt.Println(d.a)
	fmt.Println()
}
func main() {

	a := new(Data)
	var b Data
	b.a = 10
	DoSomething(&b)
	fmt.Println(a.c)
}
