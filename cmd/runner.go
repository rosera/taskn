/*
Copyright © 2023 Rich Rose <richardrose@google.com>

*/
package cmd

import (
	"encoding/csv"
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"os"
  "os/exec"
)

const (
	emailToken = "lab-architects@google.com"
	nameToken  = "lab-architects@google.com"
)

var	cmdToken string 

func apiOwnerUpdate(emailToken, nameToken, filename string) {
	fmt.Println("# Lab Owner")
	if nameToken == "" {
		fmt.Println(emailToken)
	} else {
		fmt.Println(emailToken, " # ", nameToken)
	}
}

func executeCommand(command string, args ...string) (string, error) {
	cmd := exec.Command(command, args...)

	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("failed to execute command: %s", err)
	}

	return string(output), nil
}

func commandExists(cmd string) bool {
	// Run the `which` command with the given command name
	cmdPath, err := exec.LookPath(cmd)
	if err != nil {
		return false
	}

	// Command exists if LookPath does not return an error
	return cmdPath != ""
}

func writeFile(filename string) error {
	qlOwnerFile, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("error creating the QL_OWNER file: %s", err)
	}
	defer qlOwnerFile.Close()

	// Redirect the standard output to the QL_OWNER file
	origStdout := os.Stdout
	os.Stdout = qlOwnerFile
	defer func() { os.Stdout = origStdout }()

	apiOwnerUpdate(emailToken, nameToken, filename)
	return nil
}


func deleteFile(filename string) error {
	err := os.Remove(filename)
	if err != nil {
		return fmt.Errorf("error deleting file: %s", err)
	}
	return nil
}

var taskCmd = &cobra.Command{
	Use:   "runner",
	Short: "Read the input file and perform a command",
	Long:  "Read an input CSV container filenames. Repeat the named command per filename",

	Run: func(cmd *cobra.Command, args []string) {
		inputFile, _ := cmd.Flags().GetString("input")
		if inputFile == "" {
			fmt.Println("Please specify an input CSV file using --input option.")
			return
		}

		taskCommand, _ := cmd.Flags().GetString("command")
		if taskCommand == "" {
			fmt.Println("Please specify a command to use --command option.")
			return
		}

    exists := commandExists(taskCommand)

    if !exists {
			fmt.Printf("%s not found on path\n", taskCommand)
			return
    }

		file, err := os.Open(inputFile)
		if err != nil {
			fmt.Printf("Error opening the input file: %s\n", err)
			return
		}
		defer file.Close()

		csvReader := csv.NewReader(file)

		for {
			record, err := csvReader.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				fmt.Printf("Error reading CSV record: %s\n", err)
				return
			}

			for _, filename := range record {
//				if err := writeFile(filename + "_QL_OWNER"); err != nil {
//					fmt.Printf("Error writing QL_OWNER file: %s\n", err)
//				}
//          fmt.Printf("Execute: %s on %s\n",taskCommand, filename)
          output, err := executeCommand(taskCommand, filename) 
          if err != nil {
             fmt.Printf("Execution failed: %s\n", err)
          } else {
            fmt.Printf("Output: %s", output)
          }	
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(taskCmd)
	taskCmd.PersistentFlags().StringP("input", "i", "", "Input CSV file")
	taskCmd.PersistentFlags().StringP("command", "c", "", "Command to be run")
	_ = taskCmd.MarkFlagRequired("input")
}
