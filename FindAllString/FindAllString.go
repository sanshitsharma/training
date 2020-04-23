package main

import (
	"fmt"
	"regexp"
)

func main() {
	re, _ := regexp.Compile(".at")

	fmt.Println(re.FindAllString("Cat Dog Bat", -1))

	fmt.Println(re.FindAllString("Cat Dog Bat", 0))

	fmt.Println(re.FindAllString("Cat Dog Bat", 1))

	fmt.Println(re.FindAllString("Cat Dog Bat", 2))

}
