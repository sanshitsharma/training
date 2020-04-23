package main

import "fmt"

type fptr func(int, int) int

type mystruct struct {
	a int
	b int

	fa fptr
	fb fptr
}

func main() {

	s := mystruct{3, 2,
		func(pa int, pb int) int { return pa + pb },
		func(pa int, pb int) int { return pa * pb }}
	fmt.Println(s.fa(s.a, s.b))
	fmt.Println(s.fb(s.a, s.b))
}
