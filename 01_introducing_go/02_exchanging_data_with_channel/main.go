package main

import (
	"fmt"
)

func main() {
	channel := make(chan string)
	go func() {
		s := <-channel
		s = s + "World!"
		channel <- s
	}()

	s := "Hello, "
	channel <- s
	s = <-channel
	s = s + " This string was built concurrently!"
	fmt.Println(s)
}
