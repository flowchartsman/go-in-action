package main

import (
	"fmt"
)

func main() {
	// fruitWeight stores the average weight of different kinds of fruit
	// in grams
	fruitWeight := map[string]int{
		"orange":     130,
		"apple":      195,
		"watermelon": 10000,
	}

	// About how heavy are 10 apples and 30 oranges?
	fmt.Println(10*fruitWeight["apple"] + 30*fruitWeight["orange"])
	// output:
	//
}
