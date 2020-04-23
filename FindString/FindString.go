package main

import (
	"fmt"
	"regexp"
)

func main() {
	re, _ := regexp.Compile(".at")
	fmt.Println(
		re.FindString("Cat Dog Bat"))

}
