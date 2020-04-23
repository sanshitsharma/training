package main

import "fmt"

import "regexp"

func main() {

	re, _ := regexp.Compile("ted|carol")
	fmt.Println(re.FindString("bob ted"))

	re, _ = regexp.Compile("Bob (Jones|Smith)")
	fmt.Println(re.FindAllString(
		"Bob Jones and Bob Wilson and Bob Smith", -1))

}
