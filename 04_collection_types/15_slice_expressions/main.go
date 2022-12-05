package main

import (
	"fmt"
)

// avg returns the average of a []float64.
func avg(fs []float64) float64 {
	var sum float64
	// Add each value to the sum.
	for i := range fs {
		sum += fs[i]
	}
	// Divide the sum by the length of the slice.
	return sum / float64(len(fs))
}

func getMeasurements() []float64 {
	// Simulate a large slice with 10 values and a bunch of zeros.
	return []float64{3.3, 1.5, 3.5, 2.9, 2.7, 3.0, 3.1, 4.5, 2.9, 3.5, 99: 0}
}

// printSliceInfo is a generic function that can pring out any slice and
// some basic information about it.
func printSliceInfo[T any](declaration string, s []T) {
	fmt.Printf("%s\n    %v len: %d cap: %d\n", declaration, s, len(s), cap(s))
}

func main() {
	stringSlice := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I"}

	newSliceA := stringSlice[3:6] // [D E F]
	newSliceB := stringSlice[:2]  // [A B] - "up to" index 2
	newSliceC := stringSlice[7:]  // [H I] - "starting from" index 7

	stringArray := [5]string{"A", "B", "C", "D", "E"}
	newStringSlice := stringArray[1:3] // [B C]

	fmt.Println("--Using Slice Expressions:")
	printSliceInfo(`stringSlice := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I"}`, stringSlice)
	printSliceInfo("newSliceA := stringSlice[3:6]", newSliceA)
	printSliceInfo("newSliceB := stringSlice[:2]", newSliceB)
	printSliceInfo("newSliceC := stringSlice[7:]", newSliceC)

	fmt.Println("\n--Using Slice Expressions On An Array:")
	// stringArray is an array, so the generic function can't print it. You
	// could convert it to a slice with `stringArray[:], but then it would
	// get a capacity, which arrays do not have, so rather than be
	// innacurate, it's just printed manually here.`
	fmt.Printf("stringArray := [5]string{\"A\", \"B\", \"C\", \"D\", \"E\"}\n    %v len: %d\n", stringArray, len(stringArray))
	printSliceInfo("newStringSlice := stringArray[1:3]", newStringSlice)

	intSlice := []int{1, 2, 3, 4, 5}
	newSlice1 := intSlice[:2]   // [1 2]
	newSlice2 := intSlice[:2:3] // [1 2]

	fmt.Println("\n--Using Slice Expressions To Limit Capacity:")
	printSliceInfo("intSlice := []int{1, 2, 3, 4, 5}", intSlice)
	printSliceInfo("newSlice1 := intSlice[:2]", newSlice1)
	printSliceInfo("newSlice2 := intSlice[:2:3]", newSlice2)

	fmt.Println("\n--Passing A Slice Range To A Function:")
	measurements := getMeasurements()
	fmt.Printf("averaging 10 out of of %d measurements\n", len(measurements))
	averageMeasurement := avg(measurements[:10])
	fmt.Printf("average measurement is: %f\n", averageMeasurement)
}
