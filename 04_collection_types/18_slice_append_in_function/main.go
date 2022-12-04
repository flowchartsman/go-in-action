package main

import (
	"fmt"
)

// addOneToAll adds 1 to every slice value.
func addOneToAll(s []int) {
	for i := range s {
		s[i] += 1
	}
}

// addOneToAll2 attempts to add one to every slice value, but only ends up
// operating on the new, local slice after append()
func addOneToAll2(s []int) {
	s = append(s, 0)
	for i := range s {
		s[i] += 1
	}
}

// addOneToAll3 returns the modified slice instead
func addOneToAll3(s []int) []int {
	s = append(s, 0)
	for i := range s {
		s[i] += 1
	}
	return s
}

func main() {
	slice := []int{1, 2, 3}
	// addOneToAll will modify the slice in place
	addOneToAll(slice)
	fmt.Println(slice)
	// output: [2 3 4]

	// addOneToAll2 will not modify the slice, because it performs an
	// append, which gives it a new local slice
	slice = []int{1, 2, 3}
	addOneToAll2(slice)
	fmt.Println(slice)
	// output: [1 2 3]

	// Adding capacity to the slice seems to work, but the final value is
	// missing, since the local slice retains its length
	slice = make([]int, 0, 4)
	slice[0] = 1
	slice[1] = 2
	slice[2] = 3

	addOneToAll2(slice)
	fmt.Println(slice)
	// output: [2 3 4]

	// In order to see the newly-appended slice, you need to return it,
	// just as append itself does
	slice = []int{1, 2, 3}
	slice = addOneToAll3(slice)
	fmt.Println(slice)
	// output: [2 3 4 1]
}
