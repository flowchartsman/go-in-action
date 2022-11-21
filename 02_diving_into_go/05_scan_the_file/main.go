package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		log.Println("need to provide filename!")
		os.Exit(1)
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	// Ensure the file is closed on return.
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var wordCount int

	for scanner.Scan() {
		words := strings.Fields(scanner.Text())
		wordCount += len(words)
	}

	if scanner.Err() != nil {
		log.Println(scanner.Err())
		// Return so that the file is closed.
		return
	}

	fmt.Println("Found", wordCount, "words")
}
