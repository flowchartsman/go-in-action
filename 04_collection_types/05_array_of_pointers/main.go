package main

import "fmt"

func main() {
	// Declare an integer pointer array of five elements.
	// Initialize index 0 and 1 of the array with integer pointers.
	ptrArray := [5]*int{0: new(int), 1: new(int)}
	// Assign values to index 0 and 1.
	*ptrArray[0] = 10
	*ptrArray[1] = 20

	valOfPtr := *ptrArray[0] // 10
	fmt.Println(valOfPtr)
}
