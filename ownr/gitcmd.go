package main

import (
	"fmt"
	"os"
	"os/exec"
)

func gitCommandWithConfig(config, command string) error {
	cmd := exec.Command("git", "-c", config, command)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to execute git command: %v", err)
	}

	return nil
}
