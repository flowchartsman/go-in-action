package main

import (
	"fmt"
)

func main() {
	sliceLen6 := []int{0, 0, 0, 0, 0, 0}
	sliceLen4 := []int{1, 1, 1, 1}
	sliceLen2 := []int{2, 2, 2}
	sliceLen0 := []int{}

	copy(sliceLen6, sliceLen4)
	fmt.Println(sliceLen6) // output: [1 1 1 1 0 0]

	copy(sliceLen2, sliceLen4)
	fmt.Println(sliceLen2) // output: [1 1 1]

	copy(sliceLen0, sliceLen4)
	fmt.Println(sliceLen0) // output: []

	intSlice := []int{1, 2, 3, 4, 5, 6}
	subSlice := intSlice[2:4] //[3 4]
	subSliceCopy := make([]int, len(subSlice))
	copy(subSliceCopy, subSlice)

	// Change a value in the original slice.
	intSlice[2] = 10
	// Change is visible in subSlice, but not in subSliceCopy
	fmt.Println(subSlice)     //output: [10 4]
	fmt.Println(subSliceCopy) //output: [3 4]
}
