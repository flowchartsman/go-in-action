package main

import (
	"fmt"
)

func main() {
	var floatVal float64
	// The following are different ways of writing the same value.
	floatVal = 12.         // fractional part is optional
	floatVal = 12e0        // decimal point not required with exponent
	floatVal = 012.        // leading zero is fine
	floatVal = .12e+2      // integer part is optional
	floatVal = 1_2.        // underscores allowed in floats, too
	floatVal = float64(12) // explicit type for otherwise integer literal
	// explicit type converts otherwise integer literal
	myFloat := float32(12)

	fmt.Println(floatVal, myFloat)
	// Output: 12 12

	// Floats can be negative.
	negativeFloat := -12.0
	negativeHexFloat := -0xCp0

	fmt.Println(negativeFloat, negativeHexFloat)
	// Output: -12 -12

	// Here's another easter egg for you, dear reader!
	// The following are some hexidecimal float literals for the above.
	// Note the funny exponent on the third declaration! This is because
	// hexidecimal notation for floats allows you to set the binary values
	// for the exponent and significand directory, and it's necessary to
	// account for bias.
	// ref: https://en.wikipedia.org/wiki/Double-precision_floating-point_format#Exponent_encoding
	floatHex := 0xC.0p0 // exponent required for hex floats
	floatHex = 0xCp0    // fractional part optional, but exponent required
	floatHex = 0x.Cp+4  // integer part is optional

	fmt.Println(floatVal, floatHex, myFloat)
	// Output: 12 12 12
}
