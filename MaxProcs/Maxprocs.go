package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Printf("Number of CPUs: %d\n",
		runtime.NumCPU())

	fmt.Printf("Number of Goroutines: %d\n",
		runtime.NumGoroutine())

	fmt.Printf("Maximum number of CPUs: %d\n",
		runtime.GOMAXPROCS(0))
}
