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


func gitCommandWithConfig(config string, command string) error {
	// cmd := exec.Command("git", "-C", config, command)
	cmd := exec.Command("git", "-C", command)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to execute git command %s : %v", command, err)
	}

	return nil
}
