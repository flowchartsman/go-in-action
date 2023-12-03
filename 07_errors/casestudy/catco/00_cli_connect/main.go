package main

import (
	"fmt"
	"log"

	"github.com/flowchartsman/go-in-action/catco/productdb"
)

func main() {
	client := productdb.NewClient("devel-db.qualityniblets.com")
	err := client.Connect()
	if err != nil {
		// these are all functionally identical
		fmt.Println(err)
		fmt.Printf("%s\n", err)
		fmt.Printf("%v\n", err)
	}

	// simple-statement if can be a bit more concise
	if err := client.Connect(); err != nil {
		fmt.Printf("%v\n", err)
	}

	// no newlines required with log!
	if err := client.Connect(); err != nil {
		log.Println(err)
		log.Printf("%s", err)
	}

	// or log.Fatal/Fatalf if you want to exit right away
	// (to be used with care)
	if err := client.Connect(); err != nil {
		log.Fatalf("%v", err)
	}
}
