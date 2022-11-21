package main

import "fmt"

func main() {
	// Declare a variable doubleFloat with type float64 (and a default value of 0).
	var doubleFloat float64

	// Declare a variable with type float32.
	var singleFloat float32

	// Use short declaration to create a variable f with type float64 and a
	// value of 12.1.
	f := 12.1

	// Use a type conversion if you want to use short declaration with float32.
	g := float32(12.1)

	// Print these values and their types:
	fmt.Printf("doubleFloat(%T): %f\n", doubleFloat, doubleFloat)
	fmt.Printf("singleFloat(%T): %f\n", singleFloat, singleFloat)
	fmt.Printf("f(%T): %f\n", f, f)
	fmt.Printf("g(%T): %f\n", g, g)
	// Output:
	// doubleFloat(float64): 0.000000
	// singleFloat(float32): 0.000000
	// f(float64): 12.100000
	// g(float32): 12.100000
}
