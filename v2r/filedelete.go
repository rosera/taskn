package main

import (
	"fmt"
	"os"
)

// Task: Delete V2 YAML to file
// ------------------------------------------------------------------------
func deleteYamlFile(filename string) error {
	// File exists, delete it
	err := os.Remove(filename)
	if err != nil {
		return fmt.Errorf("failed to delete file: %v", err)
	}

 // fmt.Println("deleteYamlFile")

	// File deleted successfully
	return nil
}
