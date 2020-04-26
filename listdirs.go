package main

import (
	"flag"
	"fmt"
)

func main() {

	dirArg := flag.String("dir", ".", "Directory to browse")
	flag.Parse()

	filenames, err := FilePathWalkDir(*dirArg)
	if err != nil {
		fmt.Printf("Invalid directory: %v\n", err)
		return
	}

	// Print every file in dir
	for _, file := range filenames {
		fmt.Println(file)
	}
}
