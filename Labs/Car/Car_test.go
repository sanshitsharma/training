package car

import "testing"

func TestIncrease(t *testing.T) {
	test := Increase(100, 200)
	if test != 300 {
		t.Error("Increase failed")
	}
}

func TestDecrease(t *testing.T) {
	test := Decrease(100, 200)
	if test != 100 {
		t.Error("Decrease failed")
	}
}
