package main

import (
	"fmt"
	"regexp"
)

func main() {

	match1, _ := regexp.MatchString(`^\d{5}$`, "54211")

	match2, _ := regexp.MatchString(`^\d{3}-\d{2}-\d{4}$`, "155-67-1234")

	match3, _ := regexp.MatchString(`^[0-9]{3}-[0-9]{3}-[0-9]{4}$`, "543211")

	if match1 {
		fmt.Println("Valid zipcode")
	} else {
		fmt.Println("Invalid zipcode")
	}

	if match2 {
		fmt.Println("Valid SSN#")
	} else {
		fmt.Println("Invalid SSN#")
	}

	if match3 {
		fmt.Println("Valid phone #")
	} else {
		fmt.Println("Invalid phone #")
	}
}
