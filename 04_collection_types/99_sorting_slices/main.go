package main

import (
	"fmt"
	"sort"
)

func main() {
	intSlice := []int{5, 2, 1, 4, 3}
	sort.Ints(intSlice)
	fmt.Println(intSlice) // output: [1 2 3 4 5]

	float64Slice := []float64{1.5, 1.2, 1.1, 1.4, 1.3}
	sort.Float64s(float64Slice)
	fmt.Println(float64Slice) // output: [1.1 1.2 1.3 1.4 1.5]

	stringSlice := []string{"Elephant", "Alligator", "Cat", "Bear", "Dog"}
	sort.Strings(stringSlice)
	fmt.Println(stringSlice) // output: [Alligator Bear Cat Dog Elephant]
}
