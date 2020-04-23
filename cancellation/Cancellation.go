package main

import "fmt"
import "sync"
import "time"

func TaskA() {
	defer wg.Done()

	for {
		select {
		//	case <-doSomething:
		//
		case <-stop:
			//		 cleanup
			fmt.Println("Exiting")
			return
		}
	}
}

func TaskB() {
	defer wg.Done()
}

var wg sync.WaitGroup
var stop = make(chan struct{})

func main() {

	go TaskA()
	go TaskB()
	wg.Add(2)

	time.Sleep(time.Second * 5)
	close(stop)
	time.Sleep(time.Second * 5)
}
