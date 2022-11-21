package main

import (
	"fmt"
)

func main() {
	var person struct {
		name string
		age  int
	}

	person.name = "Andy"
	person.age = 42
	fmt.Println(person.name, person.age) // output: Andy 42

	james := struct {
		name string
		age  int
	}{
		name: "James",
		age:  25,
	}
	fmt.Println(james.name, james.age) // output: James 26
}
