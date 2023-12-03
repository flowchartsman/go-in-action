package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/flowchartsman/go-in-action/catco/productdb"
)

const namingGuideLines = `- Item Description must contain at least 30 characters
- Item Description must contain the full item name
- Item Description must be enthusiastic (ends with an exclamation point)
- Item Description cannot contain CERTAIN PEOPLE'S NAMES

Please refer to https://wiki.qualityniblets.com/naming`

func main() {
	client := productdb.NewClient("devel-db.qualityniblets.com")
	// Check for connection error.
	if err := client.Connect(); err != nil {
		log.Print(err)
		os.Exit(1)
	}
	defer client.Close()
	item, err := client.GetItem(1)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}
	item.Description = "Scrumpy Numplings: They're Steve approved, lmfao!"

	if err := client.UpdateItem(1, item); err != nil {
		log.Printf("cannot update item: %v", err)
		var descriptionErr *productdb.DescriptionError
		if errors.As(err, &descriptionErr) {
			fmt.Printf("%s\n", namingGuideLines)
		}
	}
}
