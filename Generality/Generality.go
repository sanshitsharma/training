package main

type Data struct {
	a int
	b int
	c float32
}

type MyInterface interface {
	FuncA()
	FuncB()
}

func (Data) FuncA() {

}

func (Data) FuncB() {

}

func MyFunc() MyInterface {
	d := new(Data)
	return d
}
