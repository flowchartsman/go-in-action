package main

import "fmt"

func main() {
	// Declare and initialize a map with make
	intToString := make(map[int]string)
	fmt.Printf("%#v\n", intToString)

	// Declare an empty map with a literal.
	// (this is identical to using make)
	intToString = map[int]string{}
	fmt.Printf("%#v\n", intToString)

	// Declare and populate a map with a literal.
	// Keys are separated from values with a colon, and commas separate
	// key-value pairs.
	intToString = map[int]string{
		1: "one",
		2: "two",
	}
	fmt.Printf("%#v\n", intToString)
}
