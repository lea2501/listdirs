package main

import (
	"flag"
	"fmt"
)

func main() {

	// Get command arguments
	dirArg := flag.String("dir", ".", "Directory to browse")
	flag.Parse()

	// Get directory files
	filenames, err := FilePathWalkDir(*dirArg)
	// Check for errors
	if err != nil {
		fmt.Printf("Invalid directory: %v\n", err)
		return
	}

	// Print every file in dir
	for _, file := range filenames {
		fmt.Println(file)
	}
}
