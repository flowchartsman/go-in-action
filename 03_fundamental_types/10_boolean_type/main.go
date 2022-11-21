package main

import (
	"fmt"
	"time"
)

func IsEven(i int) bool {
	return i%2 == 0
}

func IsWorkHours(t time.Time) bool {
	switch t.Weekday() {
	case time.Saturday, time.Sunday:
		return false
	}
	if t.Hour() < 9 || t.Hour() > 17 {
		return false
	}
	// TODO: Add holiday tracking
	return true
}

func main() {
	var boolVar bool // Defaults to false

	// Boolean variables can be used for conditionals, either
	// on their own, or as part of a boolean expression. In this case,
	// boolVar has not been explicitly set to anything yet, so its value is
	// false, and the if block will not run.
	if boolVar {
		fmt.Println("boolVal is true.")
	}

	// Once boolVar is set to true, the same if condition will pass and the
	// block will run.
	boolVar = true
	if boolVar {
		fmt.Println("NOW boolVal is true!")
	}

	// Short declaration variables assigned to boolean expressions will be
	// given the bool type as well.
	x := 0
	y := 1
	v := x < y // Type: bool, value: true.
	fmt.Printf("Type: %T value: %v.\n", v, v)
}
