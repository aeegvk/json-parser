package main

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

func main() {
	// Check if a filename was provided
	if len(os.Args) < 2 {
		log.Fatal("Please provide a filename as a command-line argument")
	}

	// Get the filename from the command-line arguments
	filename := os.Args[1]

	// Read the JSON file
	jsonFile, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)

	// Unmarshal the JSON data
	var result map[string]interface{}
	json.Unmarshal([]byte(byteValue), &result)

	// Marshal the data with indentation
	prettyJSON, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		log.Fatal(err)
	}

	// Write the prettified JSON to another file
	err = os.WriteFile("output.json", prettyJSON, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
