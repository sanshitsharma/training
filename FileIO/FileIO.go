package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	filename1 := os.Args[1]
	filename2 := os.Args[2]

	fileBytes1, err1 := ioutil.ReadFile(filename1)
	if err1 != nil {
		panic("File1: problem")
	}

	fileBytes2, err2 := ioutil.ReadFile(filename2)
	if err2 != nil {
		panic("File2: problem")
	}

	fileString := string(fileBytes1) + string(fileBytes2)

	fmt.Println(fileString)

	err3 := ioutil.WriteFile("combined.txt",
		[]byte(fileString), 0644)
	if err3 != nil {
		panic("Combined: write failed.")
	}
}
