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
	numbers := []int{1, 2, 3, 4}
	for _, n := range numbers {
		if IsEven(n) {
			fmt.Printf("%d is even.\n", n)
		}
	}

	if IsWorkHours(time.Now()) {
		fmt.Println("During working hours - sending a push notification.")
		// notification sending code
	} else {
		fmt.Println("Outside of working hours - sending an email.")
		// email sending code
	}
}
