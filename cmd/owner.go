/*
Copyright © 2023 Rich Rose <richardrose@google.com>

*/
package cmd

import (
	"encoding/csv"
  "fmt"
	"github.com/spf13/cobra"
	"os"
  "io"
)

var inputFile string
var emailToken string = "lab-architects@google.com"
var nameToken  string = "lab-architects@google.com"

func apiOwnerUpdate(emailToken string, nameToken string) {
	fmt.Println("# Lab Owner")
	if nameToken == "" {
		fmt.Println(emailToken)
	} else {
		fmt.Println(emailToken, " # ", nameToken)
	}
}

var taskCmd = &cobra.Command{
	Use:   "owner",
	Short: "Replace owner file with default QL_OWNER",
	Long:  "Read an input CSV container filenames. Replace existing QL_OWNER file for each of the named files",

	Run: func(cmd *cobra.Command, args []string) {

		if inputFile == "" {
			fmt.Println("Please specify an input CSV file using --input option.")
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

			// for _, filename := range record {
			for range record {
				// fmt.Println(filename)
        apiOwnerUpdate(emailToken, nameToken)
			}
		}
	},
}


func init() {
  rootCmd.AddCommand(taskCmd)
	taskCmd.PersistentFlags().StringVarP(&inputFile, "input", "i", "", "Input CSV file")
}
