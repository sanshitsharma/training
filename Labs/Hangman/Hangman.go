package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

func main() {

	data, _ := ioutil.ReadFile("word.txt")
	sdata := string(data)

	length := len(sdata)

	current := strings.Repeat("_ ", length)

	guess := ""
	answer := []rune{}
	for _, char := range sdata {
		answer = append(answer, rune(char))
		answer = append(answer, rune(32))
	}

	for i := 0; i < 8; i++ {
		fmt.Println("Enter guess")
		fmt.Print(">")
		fmt.Scanln(&guess)
		if len(guess) == 0 {
		}

		re, _ := regexp.Compile(guess[:1])
		locations := re.FindAllStringIndex(string(sdata), -1)
		if len(locations) == 0 {
			fmt.Println("Bad guess")
			continue
		}
		for _, item := range locations {
			temp := []rune(current)
			temp[(item[0] * 2)] = rune(guess[0])
			current = string(temp)
		}
		if current == string(answer) {
			fmt.Println("You win!")
			return
		}
		fmt.Println(current)
	}
	fmt.Println("You lose!")
	fmt.Println("The answer was ", string(answer))
}
