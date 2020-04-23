package main

type car interface {
	Drive()
	Turn()
	Brake()
}

type boat interface {
	Steer()
}

type amphibious interface {
	car
	boat
}

func main() {

}
