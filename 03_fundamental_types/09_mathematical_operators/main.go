package main

func main() {
	a := 7
	b := 3

	var i int
	i = a + b // 10
	i = a - b // 4
	i = b - a // -4
	i = a * b // 21
	i = a % b // 1

	// Divide by zero can be caught by the compiler if the 0 is a literal.
	//c = a / 0  // error: invalid operation: division by zero

	// Otherwise, divide by zero is a runtime panic.
	zero := 0
	//c = a / zero // panic: runtime error: integer divide by zero

	// Float divide by zero is +Inf
	f := 1.0
	f = f / float64(zero) // +Inf

	u := uint64(1)
	// u = u + i // error: mismatched types uint64 and int
	// Type conversion is needed to allow different types to work together.
	u = u + uint64(i) // 2

	_ = i // keep compiler from complaining

}
