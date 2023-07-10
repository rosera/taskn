package main

import (
	"fmt"
  "os"
)

func deleteFile(filename string) error {
	// File exists, delete it
  err := os.Remove(filename)
	if err != nil {
		return fmt.Errorf("failed to delete file: %v", err)
	}

	// File deleted successfully
	return nil
}
