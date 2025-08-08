package main

import (
	"os"
	"os/exec"
)

func runCmd(cmd []string) error {
	instance := exec.Command(cmd[0], cmd[1:]...)

	instance.Stdin = os.Stdin
	instance.Stdout = os.Stdout
	instance.Stderr = os.Stderr

	return instance.Run()
}
