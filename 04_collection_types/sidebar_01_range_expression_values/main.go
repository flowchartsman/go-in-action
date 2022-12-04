package main

import (
	"fmt"
)

func main() {
	strings := [...]string{"hello", "gophers"}

	// This range loop will not actually modify anything.
	// Get each element of the array as variable 's'.
	for _, s := range strings {
		// Add an exclamation point.
		s = s + "!"
	}
	fmt.Println(strings)
	// output: [hello gophers]

	// To actually modify the results, you would need to access them by index.
	for i := range strings {
		strings[i] = strings[i] + "!"
	}
	fmt.Println(strings)
	// output: [hello! gophers!]
}
