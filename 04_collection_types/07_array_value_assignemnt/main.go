package main

import (
	"fmt"
)

func main() {
	// Declare a string array of length five.
	var colors [5]string

	// Initialize another five element array with colors.
	favoriteColors := [5]string{"Red", "Blue", "Green", "Yellow", "Pink"}

	colors = favoriteColors

	fmt.Println(colors)
	fmt.Println(favoriteColors)
	// output:
	// [Red Blue Green Yellow Pink]
	// [Red Blue Green Yellow Pink]
}
