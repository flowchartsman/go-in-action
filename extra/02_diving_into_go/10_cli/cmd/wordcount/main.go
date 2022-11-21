package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/flowchartsman/go-in-action/extra/02_diving_into_go/10_cli/wordcount"
)

func main() {
	var (
		doWords bool
		doLines bool
	)
	flag.BoolVar(&doWords, "w", false, "count words in a file (default)")
	flag.BoolVar(&doLines, "l", false, "count lines in a file")
	flag.Parse()

	// default to words if neither was specified
	if !doWords && !doLines {
		doWords = true
	}
	filename := flag.Args()[0]

	counts, err := wordcount.CountWordsInFile(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if doWords {
		fmt.Printf("word count: %d\n", counts.Words)
	}
	if doLines {
		fmt.Printf("line count: %d\n", counts.Lines)
	}
}
