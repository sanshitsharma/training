package main

import "fmt"

func Test() {
	recover()
}

func FuncA() {
	defer fmt.Println("FuncA")
	FuncB()
}

func FuncB() {
	defer Test()
	FuncC()
	fmt.Println("Inside FuncB")
}

func FuncC() {
	defer fmt.Println("FuncC")
	panic("Kaboom!")
}

func main() {
	defer fmt.Println("Main")
	FuncA()
}
