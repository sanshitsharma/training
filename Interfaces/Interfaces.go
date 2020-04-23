package main

type bird interface {
    fly()
    breath()
}

type fish interface {
    swim()
}

type flyingFish struct {
}

func (flyingFish) fly() {
}

func (flyingFish) breath() {
}

func (flyingFish) swim() {
}

func FA(my fish) {
    my.swim()
}

func main() {
    value := flyingFish{}
    FA(value)
}


