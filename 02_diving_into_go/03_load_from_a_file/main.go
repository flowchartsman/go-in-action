package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	filename := "words.txt"

	fileContents, err := os.ReadFile(filename)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	words := strings.Fields(string(fileContents))

	fmt.Printf("found %d words", len(words))
}
