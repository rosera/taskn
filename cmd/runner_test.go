package cmd

import (
	"encoding/csv"
	"os"
	"testing"
)

func TestTaskCmd(t *testing.T) {
	// Prepare test input
	inputFile := "test.csv"
	testData := [][]string{
		{"file1", "file2", "file3"},
	}

	// Create a test CSV file
	file, err := os.Create(inputFile)
	if err != nil {
		t.Fatalf("Failed to create test CSV file: %s", err)
	}
	defer os.Remove(inputFile)

	// Write test data to the CSV file
	csvWriter := csv.NewWriter(file)
	for _, record := range testData {
		if err := csvWriter.Write(record); err != nil {
			t.Fatalf("Failed to write to test CSV file: %s", err)
		}
	}
	csvWriter.Flush()
	if err := csvWriter.Error(); err != nil {
		t.Fatalf("Error flushing CSV writer: %s", err)
	}

	// Run the task command
	rootCmd.SetArgs([]string{"runner", "--input", inputFile, "--command", "echo"})
	err = rootCmd.Execute()
	if err != nil {
		t.Fatalf("Task command failed: %s", err)
	}

// 	// Assert the existence of QL_OWNER files for each input filename
// 	for _, record := range testData {
// 		for _, filename := range record {
// 			qlOwnerFilename := filename + "_QL_OWNER"
// 			_, err := os.Stat(qlOwnerFilename)
// 			if err != nil {
// 				t.Errorf("QL_OWNER file %s was not created: %s", qlOwnerFilename, err)
// 			}
// 			defer os.Remove(qlOwnerFilename)
// 		}
// 	}
}
