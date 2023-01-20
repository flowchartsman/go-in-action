package main

import "fmt"

type Describer interface {
	Describe() string
}

func PrintDescription(d Describer) {
	fmt.Println(d.Describe())
}

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func (p Person) Name() string {
	return p.FirstName + " " + p.LastName
}

func (p Person) Describe() string {
	return fmt.Sprintf("%s is a %d year-old human", p.Name(), p.Age)
}

type Cat struct {
	Name     string
	Attitude string
	Color    string
}

func (c Cat) Describe() string {
	return fmt.Sprintf("%s is a %s cat (%s)", c.Name, c.Color, c.Attitude)
}

func main() {
	andy := Person{
		FirstName: "Andy",
		LastName:  "Walker",
		Age:       43,
	}
	barry := Cat{
		Name:     "Barry",
		Attitude: "the best lil' duder",
		Color:    "black and white",
	}
	benny := Cat{
		Name:     "Benny",
		Attitude: "kind of a jerk",
		Color:    "orange",
	}
	PrintDescription(andy)
	PrintDescription(barry)
	PrintDescription(benny)
}
