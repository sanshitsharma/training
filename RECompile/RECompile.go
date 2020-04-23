package main

import (
	"fmt"
	"regexp"
)

func main() {
	re, _ := regexp.Compile("([a-z]+)ing")
	fmt.Println(re.MatchString("hacking"))
	fmt.Println(re.MatchString("not"))
	fmt.Println(re.MatchString("working"))
}
