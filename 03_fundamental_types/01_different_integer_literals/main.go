package main

import (
	"fmt"
)

func main() {
	// These are different ways of writing the same value.
	decInt := 1000         // Base 10 Notation          (no prefix)
	hexInt := 0x3E8        // Hexidecimal Notation      (Prefix: "0x")
	octInt := 0o1750       // Octal Notation            (Prefix: "0")
	octInt = 0o1750        // Alternate Octal Notation  (Prefix: "0o")
	binInt := 0b1111101000 // Binary Notation           (Prefix: "0b")

	// Optional underscores can be used to separate digits for readability
	withUnderscores := 1_000
	underscoreHex := 0x3_E_8

	// Notation is purely cosmetic, and does not affect value:
	fmt.Println(decInt, hexInt, octInt, binInt, withUnderscores, underscoreHex)
	// Output: 1000 1000 1000 1000 1000 1000

	// Negative integer literals:
	negativeInt := -10
	negativeHexInt := -0xa
	fmt.Println(negativeInt, negativeHexInt)
	// Output: -10 -10
}
