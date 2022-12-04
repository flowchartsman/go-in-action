package main

import (
	"fmt"
)

func main() {
	slice := []int{10, 20, 30}
	fmt.Printf("len: %d, cap: %d\n", len(slice), cap(slice))

	// Appending single values.
	slice = append(slice, 40)
	fmt.Println(slice)
	fmt.Printf("len: %d, cap: %d\n", len(slice), cap(slice))

	slice = append(slice, 50)
	fmt.Println(slice)
	fmt.Printf("len: %d, cap: %d\n", len(slice), cap(slice))

	// Appending multiple values.
	slice = append(slice, 60, 70)
	fmt.Println(slice)
	fmt.Printf("len: %d, cap: %d\n", len(slice), cap(slice))

	// Appending multiple values with array expansion.
	slice2 := []int{80, 90}
	slice = append(slice, slice2...)
	fmt.Println(slice)
	fmt.Printf("len: %d, cap: %d\n", len(slice), cap(slice))
}
