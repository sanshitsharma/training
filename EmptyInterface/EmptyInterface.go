package main

func myFunc(param interface{}) {
	variable := param.(int)
	variable = variable + 5
}
func main() {
	data := "my data"
	myFunc(data)
}
