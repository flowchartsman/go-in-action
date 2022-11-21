package main

import (
	"fmt"
	"os"
	"time"
)

func removeFile(filename string) {
	fmt.Println("removing file:", filename)
	err := os.Remove(filename)
	if err != nil {
		fmt.Println("couldn't remove file:", err)
		return
	}
	fmt.Println(filename, "removed")
}

func main() {
	go removeFile("somefile.txt")
	fmt.Println("doing other stuff")
	time.Sleep(1 * time.Second)
	fmt.Println("finished other stuff")
}
