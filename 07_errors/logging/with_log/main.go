package main

import (
	"fmt"
	"log"
)

func WillError() error {
	return fmt.Errorf("an error happened")
}

func main() {
	// no newlines required with log!
	if err := WillError(); err != nil {
		log.Println(err)
		log.Printf("%v", err)
	}

	// or log.Fatal/Fatalf if you want to exit right away
	// Note: do not use this outside of main()
	if err := WillError(); err != nil {
		log.Fatalf("%v", err)
	}
}
