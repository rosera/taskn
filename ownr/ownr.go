package main

import (
	"fmt"
  "os"
)

const (
  file    = "QL_OWNER"
  content = "# Lab Owner\nlab-architects@google.com # lab-architects@google.com\n"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide the folder path as a command-line argument")
		return
	}

	filename := os.Args[1]

  var delimit = "/"
  input := filename + delimit + file
  
	// Delete the original file
	err := deleteFileIfExists(input)

	if err != nil {
	   fmt.Println(err)
     return
	}   

  // Create a new file
  writeStringToFile(input, content) 

	if err != nil {
	   fmt.Println(err)
     return
	}   
}
