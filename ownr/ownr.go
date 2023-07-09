package main


import (
	"fmt"
  "os"
  "strings"
)

const (
  file    = "QL_OWNER"
  content = "# Lab Owner\nlab-architects@google.com # lab-architects@google.com"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide the filename as a command-line argument")
		return
	}

	filename := os.Args[1]

	// Check if the filename contains "QL_OWNER"
	if strings.Contains(filename, file) {
	  err := deleteFileIfExists(filename)
	  if err != nil {
	  	fmt.Println(err)
      return
	  } else {
      // Create a new file
      writeStringToFile(filename, content) 
	  }   
  }
}
