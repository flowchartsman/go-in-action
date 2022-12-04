package main

import (
	"fmt"
)

func main() {
	array := [...]int{1, 2, 3, 4}

	fmt.Println("range with index only")
	// Use a range statement to provide indices in order.
	fmt.Println("Starting range loop:")
	for i := range array {
		fmt.Printf("index %d is: %d\n", i, array[i])
	}
	fmt.Println("Range loop complete!")

	fmt.Println("range with index and value")
	// Use a range statement to provide indices and copies of each value.
	for i, item := range array {
		fmt.Printf("index %d is: %d\n", i, item)
	}

	fmt.Println("range value only")
	// Use a range statement to operate on values only.
	for _, item := range array {
		fmt.Println(item)
	}

	fmt.Println("standard for loop with len check")
	// Use a standard for loop with an index variable and a len() check.
	for i := 0; i < len(array); i++ {
		fmt.Printf("index %d is: %d\n", i, array[i])
	}

	fmt.Println("standard loop in reverse order")

	// Iterate the array backwards.
	for i := len(array) - 1; i >= 0; i-- {
		fmt.Printf("index %d is: %d\n", i, array[i])
	}

	fmt.Println("standard loop with alternate index step distance")

	// Skip every other element.
	for i := 0; i < len(array); i += 2 {
		fmt.Printf("index %d is: %d\n", i, array[i])
	}
}
