package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Data struct {
	Name string
	Age  int
}

func main() {

	file, err := os.Open("C:\\goprojects\\Json3\\sample2.json")

	if err != nil {
		fmt.Println("test")
		return
	}

	decoder := json.NewDecoder(file)

	//for {
	var data Data
	decoder.Decode(&data)
	fmt.Println(data)

	_, status := decoder.Token()

	if status == io.EOF {
		fmt.Println("Done")
	}
}
