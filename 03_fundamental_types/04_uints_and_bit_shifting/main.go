package main

import "fmt"

func main() {
	// Perform bitwise operations on unsigned types.
	var (
		d uint8
		e uint8
	)
	d = 0b00001010 // (10)
	e = 0b00000111 // (7)

	var u uint8

	// Binary representations can be printed using the %04b print verb,
	// which pads to 4 digits if necessary
	// Try it out for each of the following
	u = d & e  // 00000010 (2)
	u = d | e  // 00001111 (15)
	u = d ^ e  // 00001101 (13)
	u = d &^ e // 00001000 (8)

	fmt.Printf("%b (%d)\n", u, u)
}
