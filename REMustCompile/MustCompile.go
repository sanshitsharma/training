package main

import (
	"fmt"
	"regexp"
)

func expressionFailed() {
	if nil != recover() {
		fmt.Println("RE failed")
	}
}

func main() {
	defer expressionFailed()
	re := regexp.MustCompile("([a-z]+)ing")
	fmt.Println(re.MatchString("hacking"))
	fmt.Println(re.MatchString("not"))
	fmt.Println(re.MatchString("working"))
}
