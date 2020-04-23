package main

import (
	"fmt"
	"regexp"
)

func main() {

	re, _ := regexp.Compile(`\w+`)

	//       01234567890123456789012
	source := "ted carol bob and alice"

	fmt.Println(re.FindAllString(source, -1))
	locations := re.FindAllStringIndex(source, -1)

	for _, item := range locations {
		fmt.Println(source[item[0]:item[1]])
	}
}
