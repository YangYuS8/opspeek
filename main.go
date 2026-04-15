package main

import (
	"fmt"
	"os/exec"
)

func runCheck(command string, args ...string) (string, error) {
	cmd := exec.Command(command, args...)
	err := cmd.Run()
	if err != nil {
		return "NOT AVAILABLE", err
	}
	return "RUNNING", nil
}

func checkDocker() (string, error) {
	return runCheck("docker", "info")
}

func checkKubernetes() (string, error) {
	return runCheck("kubectl", "version", "--client")
}

func checkTailscale() (string, error) {
	return runCheck("tailscale", "status")
}

func main() {
	fmt.Println("=== opspeek ===")
	
	dockerStatus, dockerErr := checkDocker()
	fmt.Println("Docker:", dockerStatus)
	if dockerErr != nil {
		fmt.Println("Docker error:", dockerErr)
	}

	kubernetesStatus, kubernetesErr := checkKubernetes()
	fmt.Println("Kubernetes:", kubernetesStatus)
	if kubernetesErr != nil {
		fmt.Println("Kubernetes error:", kubernetesErr)
	}

	tailscaleStatus, tailscaleErr := checkTailscale()
	fmt.Println("Tailscale:", tailscaleStatus)
	if tailscaleErr != nil {
		fmt.Println("Tailscale error:", tailscaleErr)
	}
}
