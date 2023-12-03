package main

import (
	"fmt"
)

func a() {
	fmt.Println("a() --> b()")
	b()
}

func b() {
	fmt.Println("b() --> c()")
	c()
	fmt.Println("a() <-- b()")
}

func c() {
	fmt.Println("c()")
	panic("Panic! At The Subroutine")
	fmt.Println("b() <-- c()")
}

func main() {
	fmt.Println("main() --> a()")
	a()
	fmt.Println("done!")
}
