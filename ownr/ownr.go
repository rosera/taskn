package main

import (
	"fmt"
  "os"
  "regexp"
)

const (
  file    = "QL_OWNER"
  content = "# Lab Owner\nlab-architects@google.com # lab-architects@google.com\n"
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
  regexPattern := `(?:gsp[0-9]{1,4})` 
  re := regexp.MustCompile(regexPattern)
  gitBranchName := re.FindString(filename)

  // If lab id not found - stop processing
  if gitBranchName == "" {
    fmt.Print("gitBranchName is empty")
    return
  }


  // Task: Set Git config
  configs := map[string]string{
      "user.name": "$GIT_NAME",
      "user.email": "$GIT_EMAIL",
  }

	for key, value := range configs {
		err := setGitConfig(key, value)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
  

  // Task: Validate the file exists
	err := fileExists(input)

	if err != nil {
	   fmt.Println(err)
     return
	}   

  // Task: Delete the original file
	err = deleteFile(input)

	if err != nil {
	   fmt.Println(err)
     return
	}   

  // Task: Create a new file
  writeStringToFile(input, content) 

	if err != nil {
	   fmt.Println(err)
     return
	}   

  // Task: git add on the new branch
	// config := "core.editor=vim"
  addCmd := "add " + input

  // Add file to staging 
  err = gitCommandWithConfig("", addCmd) 

	if err != nil {
	   fmt.Println(err)
     return
	}   

  // Task: git commit on the new branch
  commitCmd := "commit -m 'New QL_OWNER'"

  // Add file to staging 
  err = gitCommandWithConfig("", commitCmd) 

	if err != nil {
	   fmt.Println(err)
     return
	}   

  // Task: git push on the new branch
  pushCmd := "push origin " + gitBranchName + "-owner"

  // Add file to staging 
  err = gitCommandWithConfig("", pushCmd) 

	if err != nil {
	   fmt.Println(err)
     return
	}   
}
