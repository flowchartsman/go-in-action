package wordcount

import (
	"bufio"
	"os"
	"strings"
)

// CountWordsInFile counts the number of words in a file.
func CountWordsInFile(filename string) (int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	// ensure the file is closed on return
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var wordCount int

	for scanner.Scan() {
		words := strings.Fields(scanner.Text())
		wordCount += len(words)
	}

	return wordCount, scanner.Err()
}
