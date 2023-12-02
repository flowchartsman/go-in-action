package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Println("need to provide at least one filename!")
		os.Exit(1)
	}

	for _, filename := range os.Args[1:] {
		count, err := CountWordsInFile(filename)
		if err != nil {
			log.Printf("%s - err: %v/n", err)
		} else {
			fmt.Printf("")
		}
	}
}

func CountWordsInFile(filename string) (int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return -1, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	var wordCount int

	for scanner.Scan() {
		wordCount++
	}
	if scanner.Err() != nil {
		return -1, scanner.Err()
	}
	return wordCount, err
}
