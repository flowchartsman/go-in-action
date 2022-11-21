package main

import (
	"fmt"
)

type person struct {
	name string
	age  int
}

func main() {

	andy := person{
		name: "Andy",
		age:  42,
	}
	james := person{
		name: "James",
		age:  26,
	}
	fmt.Println(andy.name, andy.age)   // output: Andy 42
	fmt.Println(james.name, james.age) // output: James 26
}
