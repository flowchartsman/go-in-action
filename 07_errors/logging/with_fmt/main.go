package main

import (
	"fmt"
)

func WillError() error {
	return fmt.Errorf("an error happened")
	// or:
	// return errors.New("an error happened")
}

func main() {
	err := WillError()
	if err != nil {
		// all of these basically do the same thing:
		fmt.Println(err)
		// fmt.Printf("%s\n", err.Error())
		// fmt.Printf("%s\n", err)
		// fmt.Printf("%v\n", err)
	}

	// simple-statement if can be a bit more concise
	if err := WillError(); err != nil {
		fmt.Printf("%v\n", err)
	}
}
