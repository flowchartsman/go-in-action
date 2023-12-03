package main

import "fmt"

func main() {
	var stringMap map[string]string

	stringMap["this"] = "is not going to end well"

	fmt.Println(stringMap["this"])
}
