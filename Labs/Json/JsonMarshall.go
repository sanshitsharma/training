package main

import (
	"encoding/json"
	"fmt"
)

const (
	Manager = iota
	VicePresident
	Warehouse
)

type Employee struct {
	Name      string
	Role      int  `json:"Type"`
	Age       int  `json:",omitempty"`
	Insurance bool `json:",omitempty"`
}

func main() {

	employees := []Employee{
		{Name: "Bob Jones", Role: Manager, Age: 55},
		{Name: "Sally Wilson", Role: VicePresident, Age: 45, Insurance: false},
		{Name: "Sam Smith", Role: Warehouse, Insurance: true},
		{Name: "Adam Freeman", Role: Warehouse, Age: 34, Insurance: true},
	}

	for _, item := range employees {
		fmt.Println(item)
		bdata, _ := json.Marshal(item)
		fmt.Printf("%s\n", bdata)
	}
}
