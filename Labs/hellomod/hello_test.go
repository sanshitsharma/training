package hellomod

import "testing"
import "fmt"
import "rsc.io/quote"

func Hello() {
    fmt.Println("Hello, world");
}

func TestHello(t *testing.T) {
    fmt.Println(quote.Hello())
//    Hello()
}
