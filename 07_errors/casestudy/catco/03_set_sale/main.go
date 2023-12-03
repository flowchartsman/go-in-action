package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/flowchartsman/go-in-action/catco/productdb"
)

func main() {
	var (
		itemID         int
		newDescription string
	)
	flag.IntVar(&itemID, "id", 1, "the item ID to update")
	flag.StringVar(&newDescription, "desc", "", "the updated item description")
	flag.Parse()
	if itemID == 0 || newDescription == "" {
		flag.Usage()
		os.Exit(1)
	}

	client := productdb.NewClient("devel-db.qualityniblets.com")
	if err := client.Connect(); err != nil {
		log.Fatalf("couldn't connect to db: %s", err)
	}
	defer client.Close()
	err := client.UpdateItemDescription(itemID, newDescription)
	if err != nil {
		log.Printf("couldn't update item description: %s", err)
		// don't exit just yet! give the user another chance if it's a bad description
		var (
			descriptionErr *productdb.DescriptionError
		)
		if errors.As(err, &descriptionErr) {
			fmt.Print("Try a new description: ")
			reader := bufio.NewReader(os.Stdin)
			newDescription, _ := reader.ReadString('\n')
			// trim off the newline
			newDescription = strings.TrimSpace(newDescription)
			err = client.UpdateItemDescription(itemID, newDescription)
			if err != nil {
				log.Printf("still couldn't update item description: %s", err)
				return
			}
		}
	}
}
