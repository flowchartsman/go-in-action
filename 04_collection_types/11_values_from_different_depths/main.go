package main

import (
	"fmt"
)

func main() {
	array2D := [2][2]int{{10, 11}, {20, 21}}

	array1D := array2D[1] // [20 21]

	intValue := array1D[0] // 20

	fmt.Println(array1D)

	fmt.Println(intValue)
}
