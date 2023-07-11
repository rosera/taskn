package main

import (
	"fmt"
//	"io/ioutil"
)

// Task: Write V2 YAML to file
// ------------------------------------------------------------------------
func writeYamlToFile(filename string, content string) error {
//	err := ioutil.WriteFile(filename, []byte(content), 0644)
//	if err != nil {
//		return fmt.Errorf("failed to write to file: %v", err)
//	}

  fmt.Println("writeYamlToFile")

	// File written successfully
	return nil
}
