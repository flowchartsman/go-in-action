package wordcount

import (
	"bufio"
	"os"
	"strings"
)

type Counts struct {
	Lines int
	Words int
}

func CountWordsInFile(filename string) (Counts, error) {
	file, err := os.Open(filename)
	if err != nil {
		return Counts{}, err
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
	return Counts{
		Lines: lineCount,
		Words: wordCount,
	}, nil
}
