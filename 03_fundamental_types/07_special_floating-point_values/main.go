package main

import (
	"fmt"
	"math"
)

func main() {
	f := 2.0
	posInf := math.Pow(f, 10_000)
	fmt.Println(posInf) // output: +Inf
	negInf := f - math.Pow(f, 10_000)
	fmt.Println(negInf) // output: -Inf

	notANumber := math.Log(-f)
	fmt.Println(notANumber) // output: NaN

	// second argument determines the infinity to check for
	// <0: negative infinity
	// >0: positive infinity
	// 0: either infinity
	fmt.Println(math.IsInf(posInf, 0))  // true
	fmt.Println(math.IsInf(negInf, 1))  // false
	fmt.Println(math.IsInf(negInf, -1)) // true

	fmt.Println(math.IsNaN(notANumber)) // true
}
