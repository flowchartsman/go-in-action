package main

import (
	"fmt"
)

func main() {
	var intArray [3]int
	myArray := [3]int{1, 2, 3}

	fmt.Println(len(intArray)) //output: 3
	fmt.Println(len(myArray))  //output: 3
}
