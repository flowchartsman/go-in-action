package main

import (
	"fmt"
	"runtime/debug"
)

func main() {
	defer func() {
		if p := recover(); p != nil {
			stacktrace := debug.Stack()
			fmt.Printf("caught the panic value %v at:\n%s", p, string(stacktrace))
			fmt.Println("but I'm still in control!")
		}
	}()
	panic("I panicked!")
}
