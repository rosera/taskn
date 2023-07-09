package main


import (
	"fmt"
  "os"
)

func deleteFileIfExists(filename string) error {
	// Check if the file exists
	_, err := os.Stat(filename)
	if err != nil {
		if os.IsNotExist(err) {
			// File does not exist, return without error
			return nil
		}
		// Error occurred while checking file existence
		return fmt.Errorf("error checking file existence: %v", err)
	}

	// File exists, delete it
	err = os.Remove(filename)
	if err != nil {
		return fmt.Errorf("failed to delete file: %v", err)
	}

	// File deleted successfully
	return nil
}
