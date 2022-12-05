package main

import (
	"fmt"
)

type person struct {
	firstName string
	lastName  string
	age       int
}

func main() {
	favoriteFoods := map[person][]string{}

	andy := person{
		firstName: "Andy",
		lastName:  "Walker",
		age:       42,
	}

	john := person{
		firstName: "John",
		lastName:  "Smith",
		age:       53,
	}

	favoriteFoods[andy] = []string{"jambalaya", "bulgogi"}
	favoriteFoods[john] = []string{"crab cakes", "licorice"}

	johnFoods, found := favoriteFoods[person{
		firstName: "John",
		lastName:  "Smith",
		age:       29,
	}]
	// The "firstName" and "lastName" fields are the same, but the "age"
	// field is different, so there is no match.
	fmt.Println(johnFoods, found) // output: [], false

	// All fields match a key value in the map, so the element value is
	// returned.
	andyFoods, found := favoriteFoods[person{
		firstName: "Andy",
		lastName:  "Walker",
		age:       42,
	}]
	fmt.Println(andyFoods, found) // output: [jambalaya bulgogi] true
}
