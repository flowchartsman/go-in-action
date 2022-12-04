package main

import (
	"fmt"
	"log"
)

func printMedalists(raceName string, contestants []string) {
	fmt.Printf("Results for %s:\n", raceName)
	fmt.Printf("Gold Medal: %s\n", contestants[0])
	fmt.Printf("Silver Medal: %s\n", contestants[1])
	fmt.Printf("Bronze Medal: %s\n", contestants[3])
}

func printMedalists2(raceName string, contestants []string) error {
	if len(contestants) < 3 {
		return fmt.Errorf("not enough contestants for %s. Want: 3 Got: %d", raceName, len(contestants))
	}
	fmt.Printf("Results for %s:\n", raceName)
	fmt.Printf("Gold Medal: %s\n", contestants[0])
	fmt.Printf("Silver Medal: %s\n", contestants[1])
	fmt.Printf("Bronze Medal: %s\n", contestants[3])
	return nil
}

func printMedalists3(raceName string, contestants []string) {
	fmt.Printf("Results for %s:\n", raceName)
	for place, contestant := range contestants {
		switch place {
		case 0:
			fmt.Printf("Gold Medal: %s\n", contestant)
		case 1:
			fmt.Printf("Silver Medal: %s\n", contestant)
		case 2:
			fmt.Printf("Bronze Medal: %s\n", contestant)
		default:
			fmt.Printf("Finisher: %s\n", contestant)
			// or just return
		}
	}
}

func main() {
	log.SetFlags(0)
	race1 := []string{"Sally", "Steve", "Jo", "Adam"}
	printMedalists("Race 1", race1)

	race2 := []string{"Jon", "Jen"}

	// The following call will panic because printMedalists does not
	// properly handle out of bounds errors.
	// printMedalists("Race 2", race2)

	// Test an invalid race with new error handling.
	err := printMedalists2("Race 2", race2)
	if err != nil {
		log.Println(err)
	}

	race3 := []string{}

	printMedalists3("Race 1", race1)
	// Race 2 only has two contestants, but we still need to list them!
	printMedalists3("Race 2", race2)
	// Throw in a race with no contestants at all.
	printMedalists3("Race 3", race3)
}
