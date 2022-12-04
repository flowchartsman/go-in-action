package main

import (
	"fmt"
)

func main() {
	groupNouns := map[string]string{
		"eagle":    "convocation",
		"cat":      "clowder",
		"kangaroo": "mob",
		"dog":      "pack",
	}

	fmt.Println("Iterating map with keys and values")
	for k, v := range groupNouns {
		fmt.Printf("A group of %ss is called a %s.\n", k, v)
	}

	familyAges := map[string]int{
		"Grandpa": 83,
		"Homer":   36,
		"Marge":   34,
		"Bart":    10,
		"Lisa":    8,
		"Maggie":  1,
	}

	for person, age := range familyAges {
		if age < 18 {
			delete(familyAges, person)
		}
	}
	fmt.Println(familyAges)
	// output: map[Grandpa:83 Homer:36 Marge:34]
}
