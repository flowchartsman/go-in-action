package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

func main() {
	// Complex numbers need to be declared with the built-in complex
	// function or with complex expressions.
	complexVal1 := complex(1.1, 2.2)                   // complex128
	complexVal2 := complex(float32(1.1), float32(2.2)) // complex64

	complexVal3 := complex(1, 2)       // complex128, equivalent to 1+2i
	complexVal4 := complex(0, -3)      // complex128, equivalent to 0-3i
	complexVal5 := complex(-3.3, -4.4) // complex128, equivalent to -3.3-4.4i

	// Complex Expression Equivalents to the above.
	complexVal1 = 1.1 + 2.2i  // equivalent to complex(1.1, 2.2)
	complexVal3 = 1 + 2i      // equivalent to complex(1,2)
	complexVal4 = -3i         // equivalent to complex(0, -3)
	complexVal5 = -3.3 - 4.4i // equivalent to complex(-3.3,-4.4)

	fmt.Println(complexVal1)
	fmt.Println(complexVal2)
	fmt.Println(complexVal3)
	fmt.Println(complexVal4)
	fmt.Println(complexVal5)

	// Hello, dear reader! We knew you wouldn't be completely satisfied if
	// you didn't see at least one example of complex numbers in use, so we
	// scoured our brains and about the simplest use we could come up with
	// was calculating the vertices of a regular polygon using roots of
	// unity.
	//
	// Basically, you can input any integer number >=3, and get the vertices
	// of a regular polygon with that many sides. Even this "simple" example
	// needs to dip into both the math and math/cmplx packages to get useful
	// constants and to represent complex numbers in rectangular form, so
	// it's not exactly the most intuitive thing in the world. But, hey,
	// there it is.
	vertices, _ := nPolygon(7)
	for _, v := range vertices {
		fmt.Printf("%f,%f\n", v.x, v.y)
	}
}

type vertex struct {
	x float64
	y float64
}

func nPolygon(n int) ([]vertex, error) {
	if n <= 3 {
		return nil, fmt.Errorf("n must be > 3")
	}
	output := make([]vertex, n)

	var c complex128
	for j := 0; j < n; j++ {
		c = cmplx.Rect(1, 2*math.Pi*float64(j)/float64(n))
		output[j] = vertex{x: real(c), y: imag(c)}
	}
	return output, nil
}
