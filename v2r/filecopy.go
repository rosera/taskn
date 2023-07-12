package main

import (
	"fmt"
	"io"
	"os"
)

func copyFile(sourceFile, destinationFile string) error {
	// Open the source file for reading
	source, err := os.Open(sourceFile)
	if err != nil {
		return fmt.Errorf("failed to open source file: %v", err)
	}
	defer source.Close()

	// Create the destination file for writing
	destination, err := os.Create(destinationFile)
	if err != nil {
		return fmt.Errorf("failed to create destination file: %v", err)
	}
	defer destination.Close()

	// Copy the contents from source to destination
	_, err = io.Copy(destination, source)
	if err != nil {
		return fmt.Errorf("failed to copy file contents: %v", err)
	}

	fmt.Println("File copied successfully")
	return nil
}
