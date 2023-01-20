package main

import "fmt"

type Person struct {
	FirstName string
	LastName string
	Age int
}

func (p Person) Name() string {
	return p.FirstName + " " + p.LastName
}

func (p Person) Describe() string {
	return fmt.Sprintf("%s is %d years old.\n", p.Name(), p.Age)
}

func main(){
	andy := Person{
		FirstName: "Andy",
		LastName: "Walker",
		Age: 43,
	}
	fmt.Println(andy.Describe())
}
