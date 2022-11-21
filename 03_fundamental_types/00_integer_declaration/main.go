package main

import "fmt"

func main() {
	// Declare a variable with a type of int (and a default value of 0).
	var myInt int

	// Declare variable with type int64
	var largeInt int64

	// Use short declaration to create variable with type int and a value
	// of 3.
	i := 3

	// Use a type conversion if you want to use short declaration with other
	// integer types.
	u := uint64(4)

	// Print these values and their types:
	fmt.Printf("myInt(%T): %d\n", myInt, myInt)
	fmt.Printf("largeInt(%T): %d\n", largeInt, largeInt)
	fmt.Printf("i(%T): %d\n", i, i)
	fmt.Printf("u(%T): %d\n", u, u)
	// Output:
	// 	myInt(int): 0
	// largeInt(int64): 0
	// i(int): 3
	// u(uint64): 4
}
