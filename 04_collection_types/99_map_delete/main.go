package main

import (
	"fmt"
)

func main() {
	countryCodes := map[string]string{
		"FR": "France",
		"CN": "China",
		"ES": "Spain",
		"TX": "Texas",
	}
	fmt.Println(countryCodes)
	// output: map[CN:China ES:Spain FR:France TX:Texas]

	fmt.Println("Delete key \"TX\"")
	// Delete the value associated with key "TX" from the map
	delete(countryCodes, "TX")
	fmt.Println(countryCodes)
	// output: map[CN:China ES:Spain FR:France]

	fmt.Println("Delete non-existent key \"XX\"")
	// The "XX" key does not exist, so no changes will be made to the map, and
	// no error will result.
	delete(countryCodes, "XX") // Key "XX" does not exist, so no action will be
	fmt.Println(countryCodes)
	// output: map[CN:China ES:Spain FR:France]
}
