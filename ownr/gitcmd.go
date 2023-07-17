package main

import (
	"fmt"
	"os"
	"os/exec"
)

func setGitConfig(key, value string) error {
	cmd := exec.Command("git", "config", "--global", key, value)

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to set Git config: %v", err)
	}

	return nil
}

// Task: git checkout - switch to a new branch
// git checkout -b [branch]
// ------------------------------------------------------------------------
func gitCheckoutCommand(directory string, branch string) error {
	cmd := exec.Command("git", "-C", directory, "checkout", "-b", branch)
	fmt.Printf("Cmd: %s\n", cmd)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Printf("Stdout: %s\n", cmd.Stderr)
		fmt.Printf("Stderr: %s\n", cmd.Stderr)
		return fmt.Errorf("Git failed to execute Add command %s : %v", cmd, err)
	}

	return nil
}

// Task: git add - add file to staging
// git add QL_OWNER
// ------------------------------------------------------------------------
func gitAddCommand(directory string, file string) error {
	cmd := exec.Command("git", "-C", directory, "add", file)
	fmt.Printf("Cmd: %s\n", cmd)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Printf("Stdout: %s\n", cmd.Stderr)
		fmt.Printf("Stderr: %s\n", cmd.Stderr)
		return fmt.Errorf("Git failed to execute Add command %s : %v", cmd, err)
	}

	return nil
}

// Task: git add - add file to staging
// git commit -m "Add: QL_OWNER"
// ------------------------------------------------------------------------
func gitCommitCommand(directory string, msg string) error {
	cmd := exec.Command("git", "-C", directory, "commit", "-m", msg)
	fmt.Printf("Cmd: %s\n", cmd)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Printf("Stdout: %s\n", cmd.Stderr)
		fmt.Printf("Stderr: %s\n", cmd.Stderr)
		return fmt.Errorf("Git failed to execute Commit command %s : %v", cmd, err)
	}

	return nil
}

// Task: git push - add file to staging
// git push origin [branch]
// ------------------------------------------------------------------------
func gitPushCommand(directory string, branch string) error {
	cmd := exec.Command("git", "-C", directory, "push", "origin", branch)

	fmt.Printf("Cmd: %s\n", cmd)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Printf("Stdout: %s\n", cmd.Stderr)
		fmt.Printf("Stderr: %s\n", cmd.Stderr)
		return fmt.Errorf("Git failed to execute Push command %s : %v", cmd, err)
	}

	// gitExecuteCommand(cmd)

	return nil
}

func gitExecuteCommand(cmd *exec.Cmd) error {
	fmt.Printf("Cmd: %s\n", cmd)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Printf("Stdout: %s\n", cmd.Stderr)
		fmt.Printf("Stderr: %s\n", cmd.Stderr)
		return fmt.Errorf("Git failed to execute Push command %s : %v", cmd, err)
	}

	return nil
}

// func gitCommandWithConfig(directory string, command string) error {
// 	// cmd := exec.Command("git", "-C", config, command)
// 	// cmd := exec.Command("git", "-C", directory, command)
// 	cmd := exec.Command("git", "-C", directory, "add", "QL_OWNER")
// 	// cmd := exec.Command("git", "--git-dir="+directory, "--work-tree="+directory,  command)
// 	// cmd := exec.Command("sh", "-c", "git", "-C", directory, command)
//   // test := "git -C " + directory + " " + command
//   // test := "git -C " + "gcp-spl-content" + " " + "labs/gsp001-creating-a-virtual-machine/QL_OWNER"
//   fmt.Printf("Cmd: %s\n", cmd)
// 	// cmd := exec.Command(test)
// 	cmd.Stdout = os.Stdout
// 	cmd.Stderr = os.Stderr
//
// 	err := cmd.Run()
// 	if err != nil {
//     fmt.Printf("Stdout: %s\n",cmd.Stderr)
//     fmt.Printf("Stderr: %s\n",cmd.Stderr)
// 		return fmt.Errorf("failed to execute git command %s : %v", command, err)
// 	}
//
// 	return nil
// }
