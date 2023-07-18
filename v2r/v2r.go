package main

import (
	"fmt"
	"github.com/rosera/v2-qwiklabs"
	"os"
	"regexp"
)

const (
	file      = "qwiklabs.yaml"
	branch    = "-v2yaml"
	developer = "lab-architects"
	email     = "lab-architects@google.com"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide the folder path as a command-line argument")
		return
	}

	// TODO: Argument is folder/path
	// Task: Read the folder argument
	// ------------------------------------------------------------------------
	folderPath := os.Args[1]

	// var delimit = "/"
	// input := folderPath + delimit + file

	// Task: Set Git config
	// ------------------------------------------------------------------------
	configs := map[string]string{
		"user.name":  developer,
		"user.email": email,
	}

	for key, value := range configs {
		err := setGitConfig(key, value)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	// Task: Regex pattern for lab identifier
	// ------------------------------------------------------------------------
	regexPattern := `(?:gsp[0-9]{1,4}|GSP[0-9]{1,4})`
	re := regexp.MustCompile(regexPattern)
	labId := re.FindString(folderPath)

	// If lab id not found - stop processing
	if labId == "" {
		fmt.Print("gitBranchName is empty")
		return
	}

	// Task: Set a branch name
	// ------------------------------------------------------------------------
	gitBranchName := labId + branch

	// Task: git checkout new branch
	// ------------------------------------------------------------------------
	fmt.Printf("BRANCH: %s\n", gitBranchName)
	fmt.Printf("PATH: %s\n", folderPath)

	// Switch to new branch
	err := gitNewCheckoutCommand(folderPath, gitBranchName)

	if err != nil {
		fmt.Println(err)
		return
	}

	// Task: Set the file names
	// ------------------------------------------------------------------------
	sourceFile := folderPath + "qwiklabs.yaml"
	destinationFile := folderPath + "qwiklabs.backup"

	// Task: Validate the file exists
	// ------------------------------------------------------------------------
	err = fileExists(sourceFile)

	if err != nil {
		fmt.Println(err)
		return
	}

	// Task: Validate the file exists
	// ------------------------------------------------------------------------
	err = copyFile(sourceFile, destinationFile)
	if err != nil {
		fmt.Println(err)
	}

	// Task: Delete the original file
	// ------------------------------------------------------------------------
	err = deleteYamlFile(sourceFile)

	if err != nil {
		fmt.Println(err)
		return
	}

	// Task: Create a new file
	// ------------------------------------------------------------------------
	var v2schema v2.SchemaV2

	// TODO: Use v2-qwiklabs package to read the file qwiklabs.backup
	v2schema.ReadV2Schema(&destinationFile)

	// Replace the allowed_location value
	for i := range v2schema.Environment.Resources {
		// TODO: Set the defaults
		if v2schema.Environment.Resources[i].Type == "gcp_project" {
			v2schema.Environment.Resources[i].Variant = "gcpd"
			v2schema.Environment.Resources[i].AllowedLocations = []string{"us-central1"}
		}
	}

	// TODO: Write the updated V2 content to folder/qwiklabs.yaml
	v2schema.WriteV2Schema(sourceFile, "baseline")

	// Task: git add on the new branch
	// ------------------------------------------------------------------------
	// config := "core.editor=vim"
	// addCmd := "add " + file

	fmt.Printf("File: %s\n", sourceFile)
	fmt.Printf("PATH: %s\n", folderPath)
	//
	// Add file to staging
	err = gitAddCommand(folderPath, file)

	if err != nil {
		fmt.Println(err)
		return
	}

	// Task: git commit on the new branch
	// ------------------------------------------------------------------------
	commitCmd := fmt.Sprintf("%q", "Add: New QL_OWNER")

	fmt.Printf("MSG: %s\n", commitCmd)
	fmt.Printf("PATH: %s\n", folderPath)

	// Add file to staging
	err = gitCommitCommand(folderPath, commitCmd)

	if err != nil {
		fmt.Println(err)
		return
	}

	// Task: git push on the new branch
	// ------------------------------------------------------------------------
	pushCmd := gitBranchName

	fmt.Printf("BRANCH: %s\n", pushCmd)
	fmt.Printf("PATH: %s\n", folderPath)

	// Add file to staging
	err = gitPushCommand(folderPath, pushCmd)

	if err != nil {
		fmt.Println(err)
		return
	}

	// Switch to the main branch
	err = gitMainCheckoutCommand(folderPath, gitBranchName)

	if err != nil {
		fmt.Println(err)
		return
	}
}
