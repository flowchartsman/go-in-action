package main

import (
	"fmt"
)

type Driver interface {
	Drive()
}

func DoDrive(d Driver) {
	d.Drive()
}

type Truck struct {
	Model    string
	Capacity int
}

func (t Truck) Drive() {
	fmt.Println("Heavy load coming through!")
}

func main() {
	var t Truck
	DoDrive(t)
}
