package main

import (
	"fmt"
	"os"
	"regexp"
)

const (
	file      = "QL_OWNER"
	content   = "# Lab Owner\nlab-architects@google.com # lab-architects@google.com\n"
	branch    = "-lab-owner"
	developer = "lab-architects"
	email     = "lab-architects@google.com"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide the folder path as a command-line argument")
		return
	}

	filename := os.Args[1]

	var delimit = "/"
	input := filename + delimit + file

	// Task: Regex pattern for lab identifier
	// Match GSPXXXX | gspXXXX
	// ------------------------------------------------------------------------
	regexPattern := `(?:gsp|GSP[0-9]{1,4})`
	re := regexp.MustCompile(regexPattern)
	labId := re.FindString(filename)

	// If lab id not found - stop processing
	if labId == "" {
		fmt.Print("gitBranchName is empty")
		return
	}

	// Task: Set a branch name
	gitBranchName := labId + branch

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

	// Task: git checkout new branch
	// ------------------------------------------------------------------------
	fmt.Printf("BRANCH: %s\n", gitBranchName)
	fmt.Printf("PATH: %s\n", filename)

	// Add file to staging
	err := gitCheckoutCommand(filename, gitBranchName)

	if err != nil {
		fmt.Println(err)
		return
	}

	// Task: Validate the file exists
	// ------------------------------------------------------------------------
	err = fileExists(input)

	if err != nil {
		fmt.Println(err)
		return
	}

	// Task: Delete the original file
	// ------------------------------------------------------------------------
	err = deleteFile(input)

	if err != nil {
		fmt.Println(err)
		return
	}

	// Task: Create a new file
	// ------------------------------------------------------------------------
	writeStringToFile(input, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	// Task: git add on the new branch
	// config := "core.editor=vim"
	// addCmd := "add " + file
	// ------------------------------------------------------------------------
	fmt.Printf("File: %s\n", file)
	fmt.Printf("PATH: %s\n", filename)

	// Add file to staging
	err = gitAddCommand(filename, file)

	if err != nil {
		fmt.Println(err)
		return
	}

	// Task: git commit on the new branch
	// ------------------------------------------------------------------------
	commitCmd := fmt.Sprintf("%q", "Add: New QL_OWNER")

	fmt.Printf("MSG: %s\n", commitCmd)
	fmt.Printf("PATH: %s\n", filename)

	// Add file to staging
	err = gitCommitCommand(filename, commitCmd)

	if err != nil {
		fmt.Println(err)
		return
	}

	// Task: git push on the new branch
	// ------------------------------------------------------------------------
	pushCmd := gitBranchName

	fmt.Printf("BRANCH: %s\n", pushCmd)
	fmt.Printf("PATH: %s\n", filename)

	// Add file to staging
	err = gitPushCommand(filename, pushCmd)

	if err != nil {
		fmt.Println(err)
		return
	}
}
