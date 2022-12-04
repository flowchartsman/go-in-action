package main

import "fmt"

func main() {
	// Declare a two dimensional integer array of two elements.
	var array2D [2][2]int
	// Set integer values to each individual element.
	array2D[0][0] = 10
	array2D[0][1] = 20
	array2D[1][0] = 30
	array2D[1][1] = 40

	fmt.Println(array2D)

	// Declare two different two dimensional integer arrays.
	var array1 [2][2]int
	var array2 [2][2]int
	// Add integer values to each individual element.
	array2[0][0] = 10
	array2[0][1] = 20
	array2[1][0] = 30
	array2[1][1] = 40
	// Copy the values from array2 into array1.
	array1 = array2

	fmt.Println(array1)
	// [[10 20] [30 40]]

	// Copy index 1 of array1 into a new array of the same type.
	var array3 [2]int
	array3 = array1[1]
	fmt.Println(array3)

	// Copy the integer found in index 1 of the outer array
	// and index 0 of the interior array into a new variable of
	// type integer.
	var value int = array1[1][0]
	fmt.Println(value)
}
