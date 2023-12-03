package main

import (
	"log"
	"os"

	"github.com/flowchartsman/go-in-action/catco/productdb"
)

func main() {
	log.SetFlags(0) // Don't log any timestamps.

	// Create a client and connect to the database.
	client := productdb.NewClient("devel-db.qualityniblets.com")
	// Check for connection error.
	if err := client.Connect(); err != nil {
		log.Printf("couldn't connect to db: %s", err)
		os.Exit(1)
	}
	// Get Item 1 so that we can update its description and its price.
	item, err := client.GetItem(1)
	// Check for error retrieving item.
	if err != nil {
		log.Printf("couldn't get item: %s", err)
		os.Exit(1)
	}
	// Set the sale price and description.
	item.Price /= 2
	item.Description += " Currently half off!"
	// Store the item again, and check for an error.
	if err := client.UpdateItem(1, item); err != nil {
		log.Printf("couldn't update item: %s", err)
		os.Exit(1)
	}
}
