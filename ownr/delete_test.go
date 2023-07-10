package main

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestDeleteFile(t *testing.T) {
	// Create a temporary file
	file, err := ioutil.TempFile("", "testfile")
	if err != nil {
		t.Fatalf("Failed to create temporary file: %v", err)
	}
	defer os.Remove(file.Name())

	// Attempt to delete the file
	err = deleteFile(file.Name())

	// Verify that the file is deleted
	_, statErr := os.Stat(file.Name())
	if statErr == nil {
		t.Errorf("File still exists, expected it to be deleted")
	}

	// Verify the error returned by deleteFile
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}
