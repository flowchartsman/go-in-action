package main

import "fmt"

func main() {
	channel := make(chan string)
	output := make(chan string)

	go func() {
		// First goroutine waits for a message.
		s := <-channel
		s = s + "World!"
		output <- s
	}()

	go func() {
		// Second goroutine sends the initial message.
		s := "Hello, "
		channel <- s
	}()

	// Main goroutine waits for the final output.
	finalString := <-output
	fmt.Println("This string was built concurrently:", finalString)
}
