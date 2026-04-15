package main

import (
	"fmt"
	"os/exec"
)

func checkDocker() (string, error) {
	cmd := exec.Command("docker", "info")
	err := cmd.Run()
	if err != nil {
		return "NOT AVAILABLE", err
	}
	return "RUNNING", nil
}

func checkTailscale() (string, error) {
	cmd := exec.Command("tailscale", "status")
	err := cmd.Run()
	if err != nil {
		return "NOT AVAILABLE", err
	}
	return "RUNNING", nil
}

func main() {
	fmt.Println("=== opspeek ===")
	
	status, err := checkDocker()
	fmt.Println("Docker:", status)

	status, err = checkTailscale()
	fmt.Println("Tailscale:", status)

	if err != nil {
		fmt.Println("Error:", err)
	}
}
