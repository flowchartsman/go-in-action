package main

import (
	"fmt"
)

func main() {
	defer fmt.Println("defer 1")
	fmt.Println("main 1")
	defer fmt.Println("defer 2")
	fmt.Println("main 2")
}
