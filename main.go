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
	
	dockerStatus, dockerErr := checkDocker()
	fmt.Println("Docker:", dockerStatus)
	if dockerErr != nil {
		fmt.Println("Docker error:", dockerErr)
	}

	tailscaleStatus, tailscaleErr := checkTailscale()
	fmt.Println("Tailscale:", tailscaleStatus)
	if tailscaleErr != nil {
		fmt.Println("Tailscale error:", tailscaleErr)
	}
}
