package car

var speed = 100

func Increase(accelerate int, current int) int {
	return current + accelerate
}

func Decrease(deaccelerate int, current int) int {

	speed := current - deaccelerate
	if speed < 0 {
		return 0
	}

	if speed > 100 {
		return 100
	}

	return speed
}
