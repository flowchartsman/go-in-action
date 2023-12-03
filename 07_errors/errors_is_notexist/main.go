package main

import (
	"encoding/json"
	"errors"
	"log"
	"os"
)

type Config struct {
	Environment string `json:"environment"`
	StringVal   string `json:"string_val"`
	IntVal      int    `json:"int_val"`
}

func GetConfig(filename string) (*Config, error) {
	fileContents, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	loadedConfig := &Config{}
	if err := json.Unmarshal(fileContents, loadedConfig); err != nil {
		return nil, err
	}
	return loadedConfig, nil
}

func defaultConfig() *Config {
	return &Config{
		Environment: "default config",
		IntVal:      0,
		StringVal:   "default",
	}
}

// This example shows how you can selectively change what you do, using
// errors.Is.
// Try running this file with `go run main.go <filename>` and try loading
// config_good.json, config_bad.json and config_notexist.json.
// Observe how the error of the file being missing can be handled differently
// than the error where the file itself is invalid.
func main() {
	if len(os.Args) < 2 { // os.Args[0] is always the program name
		log.Fatal("config file name required")
	}
	config, err := GetConfig(os.Args[1])
	if err != nil {
		switch {
		case errors.Is(err, os.ErrNotExist):
			log.Println("file not found, using defaults")
			config = defaultConfig()
		default:
			log.Fatalf("configuration file invalid: %v", err)
		}
	}
	log.Printf("we are using %s -- (%s, %d)", config.Environment, config.StringVal, config.IntVal)
}
