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

	scanner := bufio.NewScanner(file)

	// Counter to track the running total.
	var wordCount int

	for scanner.Scan() {
		// This is done for each line of the file.
		words := strings.Fields(scanner.Text())
		wordCount += len(words)
	}

	if scanner.Err() != nil {
		log.Println(scanner.Err())
		os.Exit(1)
	}

	fmt.Println("Found", wordCount, "words")
}
