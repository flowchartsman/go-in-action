package main

import (
	"fmt"
)

func main() {
	intArray := [5]int{1, 2, 3, 4, 5}

	firstValue := intArray[0] // 1
	fmt.Println(firstValue)

	// Access the item using a variable value as the index.
	index := 0
	integerValue := intArray[index] // 1
	fmt.Println(integerValue)

	// Other integer types work as well.
	index32 := int32(1)
	index64 := int64(2)
	indexUint := uint64(3)

	integerValue = intArray[index32] // 2
	fmt.Println(integerValue)
	integerValue = intArray[index64] // 3
	fmt.Println(integerValue)
	integerValue = intArray[indexUint] // 4
	fmt.Println(integerValue)
}
