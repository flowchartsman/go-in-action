package main

import "fmt"

func main() {
	// Declare an array of five integers.
	var intArray [5]int

	// Declare and initialize an array with values.
	intArrayShort := [5]int{1, 2, 3, 4, 5}

	// Declare and initialize an array with only certain values.
	intArraySparse := [5]int{0: 1, 3: 4}
	// Declare and initialize an array starting from index 2.
	intArrayLast3 := [5]int{2: 3, 4, 5}

	// Declare and initialize an array with automatic length.
	intArrayAuto := [...]int{5, 4, 3, 2, 1}

	// Declare and initialize an array with automatic length and
	// only certain values.
	intArrayAutoSparse := [...]int{0: 5, 3: 2}

	fmt.Println(intArray)
	fmt.Println(intArrayShort)
	fmt.Println(intArraySparse)
	fmt.Println(intArrayLast3)
	fmt.Println(intArrayAuto)
	fmt.Println(intArrayAutoSparse)
}
