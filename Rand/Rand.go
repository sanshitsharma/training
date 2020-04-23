package main

import "fmt"

func GlobalPanic() {

	global_err := recover()

	if global_err != nil {
		// cleanup
		fmt.Println(
			"Catastrophic exception. Exiting...")
	}
}

func main() {
	defer GlobalPanic()

	panic("kaboom")
}
