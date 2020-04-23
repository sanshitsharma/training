package main

import "fmt"
import "os"
import "regexp"

func main() {

	match := false
	if len(os.Args) == 2 {
		match, _ = regexp.MatchString("Hello", os.Args[1])
	}

	if match {
		fmt.Println("Thanks.  Hello back!")
	} else {
		fmt.Println("You did not say hello :-(")
	}
}
