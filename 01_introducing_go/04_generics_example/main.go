package main

import "fmt"

func SumNumbers[N int | float64](numberSlice ...N) N {
	var total N
	for i := range numberSlice {
		total += numberSlice[i]
	}
	return total
}

func main() {
	fmt.Println(SumNumbers(1, 2, 3))
	fmt.Println(SumNumbers(1.1, 2.2, 3.3))
}
