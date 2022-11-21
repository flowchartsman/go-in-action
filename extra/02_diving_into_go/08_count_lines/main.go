package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func countWordsInFile(filename string) (int, int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return -1, -1, err
	}
	defer file.Close()

	var wordCount int
	var lineCount int

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lineCount++
		words := strings.Fields(scanner.Text())
		wordCount += len(words)
	}
	return wordCount, lineCount, nil
}

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("usage: %s <filename>\n", os.Args[0])
		os.Exit(1)
	}

	wordCount, lineCount, err := countWordsInFile(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	//fmt.Println("found", wordCount, "words on", lineCount, "lines")
	fmt.Printf("found %d words on %d lines", wordCount, lineCount)
}
