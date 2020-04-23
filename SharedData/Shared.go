package main

import "fmt"

type test struct {
	a int
}

func main() {

	s1 := [...]int{1, 2, 3, 4, 5}
	s2 := s1[1:4]

	s2[1] = 50

	fmt.Println(s1)
	fmt.Println(s2)

	s3 := make([]int, 3)
	copy(s3, s1[1:4])

	s3[1] = 1000

	fmt.Println(s1)
	fmt.Println(s3)
}
