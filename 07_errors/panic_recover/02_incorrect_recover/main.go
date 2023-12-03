package main

import "fmt"

func main() {
	if p := recover(); p != nil {
		fmt.Println("caught the panic?", p)
	}
	panic("nope, doesn't work")
}
