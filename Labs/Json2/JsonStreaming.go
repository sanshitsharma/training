package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

func main() {
	const jsonStream = `
	[
		{"Name": "Robert", "Role": "Bob", "Salary":500000},
		{"Name": "Elliot", "Role": "Ted", "Salary":250000},
		{"Name": "Natalie", "Role": "Carol", "Salary":450000},
		{"Name": "Dyan", "Role": "Alice", "Salary":200000}
	]
`
	type Data struct {
		Name   string
		Role   string
		Salary int
	}
	dec := json.NewDecoder(strings.NewReader(jsonStream))

	dec.Token()

	for dec.More() {
		var data Data

		dec.Decode(&data)

		fmt.Println(data.Name, data.Role, data.Salary)
	}
}
