package main

func main() {
	var (
		fourArray [4]int
		fiveArray [5]int
	)
	fiveArray = fourArray
	// error: cannot use fourArray (variable of type [4]int)
	// as type [5]int in assignment
}

// Yes, this file is supposed to be red :) This is the type system doing its
// job, so worry not.
