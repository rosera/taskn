package main

import (
  "os"
	"io"
	"strings"
	"testing"
)

func captureOutput(f func()) string {
	// Redirect output to a pipe
	r, w, _ := os.Pipe()
	old := os.Stdout
	os.Stdout = w

	// Restore the original output after the function finishes
	defer func() {
		os.Stdout = old
	}()

	// Capture the output
	var capturedOutput strings.Builder
	done := make(chan bool)

	go func() {
		defer close(done)
		io.Copy(&capturedOutput, r)
	}()

	f()

	w.Close()
	<-done

	return capturedOutput.String()
}

func TestApiOwnerUpdate(t *testing.T) {
	expectedOutput := "# Lab Owner\n" + emailToken + " # " + nameToken + "\n"

	// Call the function being tested and capture the output
	actualOutput := captureOutput(func() {
		apiOwnerUpdate(emailToken, nameToken)
	})

	// Compare the expected and actual output
	if !strings.EqualFold(expectedOutput, actualOutput) {
		t.Errorf("Unexpected output:\nExpected: %s\nActual: %s", expectedOutput, actualOutput)
	}
}
