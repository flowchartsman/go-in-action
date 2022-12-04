package main

import "fmt"

func main() {
	// Declare and initialize an integer slice with length 3
	sliceWithLen := make([]int, 3) // identical to []int{0, 0, 0}
	fmt.Println(len(sliceWithLen)) // 3

	fmt.Println(sliceWithLen[1]) // 0

	// Declare and initialize an integer slice with length 3, and capacity 5
	sliceWithLenAndCap := make([]int, 3, 5)
	fmt.Println(len(sliceWithLenAndCap)) // 3

	// fmt.Println(sliceWithLenAndCap[3])
	// panic: runtime error: index out of range [3] with length 3

	// Declare and initialize an integer slice with length 0, and capacity 5
	sliceWithCap := make([]int, 0, 5)
	fmt.Println(len(sliceWithCap)) // 0

	fmt.Println(sliceWithLen)
	fmt.Println(sliceWithCap)
	fmt.Println(sliceWithLenAndCap)
}
