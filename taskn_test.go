package main

import (
	"bytes"
	"os"
	"testing"
)

// Separate function that takes input CSV data as a parameter
func processCSVData(csvData string) error {
	// Create a buffer and write the CSV data to it
	buffer := bytes.NewBufferString(csvData)

	// Replace os.Stdin with the buffer
	origStdin := rootCmd.InOrStdin()
	rootCmd.SetIn(buffer)
	defer func() { rootCmd.SetIn(origStdin) }()

	// Execute the command
	return rootCmd.Execute()
}

func TestMainFunction(t *testing.T) {
	// Create a sample CSV data to use for testing
	csvData := `file1.txt
    file2.txt
    file3.txt`

	// Call the separate function with the input CSV data
	err := processCSVData(csvData)
	if err != nil {
		t.Fatalf("Error processing CSV data: %v", err)
	}
}

func TestCommandWithInputFlag(t *testing.T) {
	// Set the command-line arguments for the test
	os.Args = []string{"test", "--input", "input.csv"}

	// Execute the main function
	main()
}
