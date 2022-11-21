package main

import (
	"fmt"
	"net/http"
)

func incrementInteger(i int) {
	i = i + 1
}

func incrementIntegerByPointer(p *int) {
	*p = *p + 1
}

func main() {
	// Taking the address of an existing variable with the reference
	// operator
	intValue := 0
	intPtr := &intValue

	// Creating a pointer directly using a reference to a type literal
	webClient := &http.Client{}

	// Initializing a new reference to an anonymous integer variable
	newIntPtr := new(int)

	// Accessing the value requires the dereference operator
	*intPtr = 1

	// passing the integer by value copies it, so the original is not
	// modified
	incrementInteger(intValue)
	fmt.Println(intValue) // output: 1

	// passing the integer by reference allows it to be modified by the
	// function
	incrementIntegerByPointer(intPtr)
	fmt.Println(intValue) // output: 2

	// Printing a pointer variable shows the address
	fmt.Println(intPtr) //output: 0x14000122008 (or similar)

	// another pointer to the same integer
	secondPtr := &intValue
	fmt.Println(secondPtr) // output: 0x14000122008 (the same as intPtr)

	// pointer math is not allowed
	//secondPtr += 1 // error: cannot convert 1 (untyped int constant) to *int

	_ = newIntPtr
	_ = webClient
}
