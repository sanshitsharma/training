package main

import "fmt"

func main() {

	var d interface{}
	d = 5

	switch d.(type) {
	case bool:
		fmt.Println("boolean")
	case int, int16, int32:
		fmt.Println("integer")
	case *bool:
		fmt.Println("boolean pointer")
	default:
		fmt.Println("undetermined")
	}
}
