package main

import (
	"fmt"
)

func main() {
	var array3D [4][3][2]string

	fmt.Printf("array3D is a %T\n", array3D)
	fmt.Printf("array3D[0] is a %T\n", array3D[0])
	fmt.Printf("array3D[0][0] is a %T\n", array3D[0][0])
	fmt.Printf("array3D[0][0][0] is a %T\n", array3D[0][0][0])
	//output:
	// variable is a [4][3][2]string
	// variable is a [3][2]string
	// variable is a [2]string
	// variable is a string
}
