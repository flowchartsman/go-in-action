package main

import "fmt"

func main() {
	// Use hard-coded text to start.
	text := "let's count some words!"

	var numSpaces int

	// Look at the text one byte at a time.
	for i := 0; i < len(text); i++ {
		if text[i] == ' ' {
			numSpaces++
		}
	}

	fmt.Println("Found", numSpaces+1, "words")
}
