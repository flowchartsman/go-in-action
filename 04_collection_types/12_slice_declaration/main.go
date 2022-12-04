package main

import "fmt"

func main() {
	// Declare an integer slice variable.
	var sliceVar []int
	// Declare and initialize an integer slice with a literal.
	sliceLit := []int{10, 20, 30}
	// Get the length and capacity of newly-created slice.
	sliceLitLen := len(sliceLit) // int(3)
	sliceLitCap := cap(sliceLit) // int(3)

	fmt.Println(sliceLitLen, sliceLitCap) //output: "3 3"

	// Check if the unintialized slice is equal to nil.
	fmt.Println(sliceVar == nil) //output: "true"
	// Get the length and capacity of an uninitialized slice.
	fmt.Println(len(sliceVar)) //output: "0"
	fmt.Println(cap(sliceVar)) //output: "0"

	fmt.Println(sliceVar)
	fmt.Println(sliceLit)
}
