package main

import (
	"errors"
	"fmt"
)

type KindOfEmployee int

const (
	salary = iota
	hourly
)

type employee struct {
	kind KindOfEmployee
}

type hourlyEmployee struct {
	employee
	rate int
	hpw  int
}

type salariedEmployee struct {
	employee
	salary int
}

type iEmployee interface {
	Calculate() int
}

func (obj hourlyEmployee) Calculate() int {
	return obj.rate * obj.hpw * 2
}

func (obj salariedEmployee) Calculate() int {
	return obj.salary / 26
}

func (obj *salariedEmployee) SetEmployeeType() {
	obj.kind = salary
}

func (obj *hourlyEmployee) SetEmployeeType() {
	obj.kind = hourly
}

func AddToPayroll(payperiod *[]iEmployee, data ...int) error {
	switch len(data) {
	case 1:
		empl := salariedEmployee{salary: data[0]}
		empl.SetEmployeeType()
		*payperiod = append(*payperiod, empl)
	case 2:
		empl := hourlyEmployee{rate: data[0], hpw: data[1]}
		empl.SetEmployeeType()
		*payperiod = append(*payperiod, empl)
	default:
		return errors.New("Invalid employee")
	}
	return nil
}

func main() {

	var employees = []iEmployee{}
	AddToPayroll(&employees, 30000)
	AddToPayroll(&employees, 40, 40)
	AddToPayroll(&employees, 30, 35)
	//result := employees[0]
	// pay := result.(salariedEmployee).Calculate()
	// fmt.Println(pay)
	for index, obj := range employees {
		fmt.Println(index, obj, obj.Calculate())
	}

}
