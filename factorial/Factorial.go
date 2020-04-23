package main

func Factorial(value int) int {
	var result int
	for result = 1; value > 1; value-- {
		result = result * value
	}

	return result
}
func main() {
	Factorial(5)
}
