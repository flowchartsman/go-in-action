package main

import (
	"fmt"
)

func main() {
	var slice []int
	fmt.Println(slice == nil) // true
	fmt.Println(len(slice))   // 0

	// A slice with an explicit len of zero is NOT nil.
	sliceLenZero := make([]int, 0)
	fmt.Println(sliceLenZero == nil) // false
	fmt.Println(len(sliceLenZero))   // 0

	// A slice made from an empty literal is also NOT nil.
	sliceEmptyLit := []int{}
	fmt.Println(sliceEmptyLit == nil) // false
	fmt.Println(len(sliceEmptyLit))   // 0

	// A slice of a slice that has a length of zero is also NOT nil.
	sliceWithValues := []int{1, 2, 3}
	slicedSlice := sliceWithValues[0:0]
	fmt.Println(slicedSlice == nil) // false
	fmt.Println(len(slicedSlice))   // 0

	// The only way to return to a nil slice is to set it to nil explicitly.
	sliceLenZero = nil
	fmt.Println(sliceLenZero == nil) // true
}
