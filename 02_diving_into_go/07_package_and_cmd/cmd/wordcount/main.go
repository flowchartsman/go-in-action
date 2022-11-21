package main

import (
	"fmt"
	"log"
	"os"

	"github.com/flowchartsman/go-in-action/02_diving_into_go/07_package_and_cmd/wordcount"
)

func main() {
	if len(os.Args) < 2 {
		log.Println("need to provide filename!")
		os.Exit(1)
	}

	wordCount, err := wordcount.CountWordsInFile(os.Args[1])
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	fmt.Println("Found", wordCount, "words")
}
