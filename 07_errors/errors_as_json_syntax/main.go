package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"

	_ "github.com/flowchartsman/go-in-action/internal/forbook"
)

type Message struct {
	Msg string `json:"msg"`
}

// invalid JSON: provides a number when a string is expected
const badJson = `{"msg":5}`

func main() {
	var m Message
	err := json.Unmarshal([]byte(badJson), &m)
	if err != nil {

		log.Println("just logging the error:")
		// just logging the error
		log.Println(err)

		log.Println("getting extra data with errors.As:")
		// localize type errors if we find them
		var typeErr *json.UnmarshalTypeError
		// check to see if this is a JSON type error, and add the offset
		// where it happened if so
		if errors.As(err, &typeErr) {
			err = fmt.Errorf("%s (offset: %d)", err, typeErr.Offset)
		}
		log.Fatal(err)
	}
	log.Printf("no errors, msg is %s", m.Msg)
}
