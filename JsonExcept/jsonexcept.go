package main

import (
	"encoding/json"
	"fmt"
)

type data struct {
	First  string
	Second int
}

func main() {
	dataJson := `{"first": "data","second": "t"}`
	var mydata data
	err := json.Unmarshal([]byte(dataJson), &mydata)
	if err != nil {
		fmt.Print(err)
		return
	}
	fmt.Printf("Species: %s, Description: %d",
		mydata.First, mydata.Second)
}
