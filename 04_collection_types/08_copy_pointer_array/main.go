package main

import "fmt"

func main() {
	// Declare an integer pointer array of five elements.
	// Initialize index 0 and 1 of the array with integer pointers.
	ptrArray := [5]*int{0: new(int), 1: new(int)}
	// Assign values to index 0 and 1.
	*ptrArray[0] = 10
	*ptrArray[1] = 20

	ptrArray2 := ptrArray

	// Pointer values are equal (point to the same location).
	fmt.Println(ptrArray[0] == ptrArray2[0]) // true

	// Which of course means the indirected values are the same.
	fmt.Println(*ptrArray[0] == *ptrArray2[0]) // true

	// Location is NOT the same, because they are separate arrays and
	// separate values now.
	fmt.Println(&ptrArray[0] == &ptrArray2[0]) // false
}
