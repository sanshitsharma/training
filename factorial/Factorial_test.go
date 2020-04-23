package main

import (
	"fmt"
	"os"
	"testing"
)

var testValues = []struct {
	value  int
	result int
}{
	{1, 1},
	{2, 5},
	{3, 6},
	{5, 120},
	{6, 600},
}

func TestMain(m *testing.M) {
	fmt.Println("Okay")
	os.Exit(m.Run())
}
func TestFactorial(t *testing.T) {

	for _, input := range testValues {
		actual := Factorial(input.value)
		if actual != input.result {
			t.Errorf("Test failed %d", input.value)
		}
	}
}
