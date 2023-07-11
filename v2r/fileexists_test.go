package main

import (
	"testing"
)

func TestFileExists_FileDoesNotExist(t *testing.T) {
	// Attempt to delete a non-existent file
	err := fileExists("nonexistentfile.txt")

	// Verify that the function returns without error
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}
