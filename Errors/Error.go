package main

import "errors"
import "fmt"

func main() {

	specificErr := errors.New("specific error")

	// much later
	fmt.Println(specificErr.Error())
}
