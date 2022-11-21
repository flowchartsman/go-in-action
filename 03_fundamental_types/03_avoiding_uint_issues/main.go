package main

import (
	"fmt"
	"log"

	_ "github.com/flowchartsman/go-in-action/internal/forbook"
)

// OnlyPositive only accepts positive values.
func OnlyPositive(u uint) {
	fmt.Println("positive value is", u)
}

// BetterOnlyPositive only accepts positive values. An error will be
// returned if a negative value is passed.
func BetterOnlyPositive(i int) error {
	if i < 0 {
		return fmt.Errorf("value must be positive!")
	}
	fmt.Println("positive value is", i)
	return nil
}

func GetConfigValue() int {
	return -1
}

func main() {
	// User has a signed integer value that might come from somewhere else
	// which is unintentionally negative.
	configVal := GetConfigValue() // returns int(-1)

	// They need to convert it to call your function, but wait...
	OnlyPositive(uint(configVal))
	// output: positive value is 1844674407370955161

	// Better to take an int and specify up front that negative values are not
	// valid.
	if err := BetterOnlyPositive(configVal); err != nil {
		// Now you can detect the error!
		log.Fatal(err)
	}
}
