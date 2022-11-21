package main

import (
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

	fileContents, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	words := strings.Fields(string(fileContents))

	fmt.Println("Found", len(words), "words")
}
