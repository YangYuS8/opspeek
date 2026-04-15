package main

import (
	"fmt"
	"os/exec"
)

func checkDocker() string {
	cmd := exec.Command("docker", "info")
	err := cmd.Run()
	if err != nil {
		return "NOT AVAILABLE"
	}
	return "RUNNING"
}

func main() {
	fmt.Println("=== opspeek ===")
	fmt.Println("Docker:", checkDocker())
}
