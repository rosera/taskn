package main

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestWriteStringToFile(t *testing.T) {

  filename := file 

	// Call the function being tested
	err := writeStringToFile(filename, content)

	// Check for any errors during the write operation
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Read the file to verify the content
	fileContent, readErr := ioutil.ReadFile(filename)
	if readErr != nil {
		t.Errorf("Failed to read the file: %v", readErr)
	}

	// Convert the file content bytes to string
	fileContentStr := string(fileContent)

	// Compare the actual content with the expected content
	if fileContentStr != content {
		t.Errorf("Unexpected file content: got '%s', expected '%s'", fileContentStr, content)
	}

	// Remove the test file
	removeErr := os.Remove(filename)
	if removeErr != nil {
		t.Errorf("Failed to remove the test file: %v", removeErr)
	}
}
