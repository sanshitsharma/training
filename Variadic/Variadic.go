package main

import "fmt"

func getstuff(anything ...interface{}) {

	for _, something := range anything {
		fmt.Println(something)
	}
}

func main() {

	getstuff(1, 3, "test", "four", func() {})

}
