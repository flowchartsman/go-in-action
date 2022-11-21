package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type counts struct {
	lines int
	words int
}

func countWordsInFile(filename string) (counts, error) {
	file, err := os.Open(filename)
	if err != nil {
		return counts{}, err
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
	return counts{
		lines: lineCount,
		words: wordCount,
	}, nil
}

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("usage: %s <filename>\n", os.Args[0])
		os.Exit(1)
	}

	count, err := countWordsInFile(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("found %d words on %d lines", count.words, count.lines)
}
