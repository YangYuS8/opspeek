package main

import (
	"fmt"
	"os/exec"
)

type CheckResult struct {
	Name   string
	Status string
	Err    error
}

func runCheck(name string, command string, args ...string) CheckResult {
	cmd := exec.Command(command, args...)
	err := cmd.Run()
	if err != nil {
		return CheckResult{Name: name, Status: "NOT AVAILABLE", Err: err}
	}

	return CheckResult{Name: name, Status: "RUNNING", Err: nil}
}

func checkDocker() CheckResult {
	return runCheck("Docker", "docker", "info")
}

func checkKubernetes() CheckResult {
	return runCheck("Kubernets", "kubectl", "version", "--client")
}

func checkTailscale() CheckResult {
	return runCheck("Tailscale", "tailscale", "status")
}

func printResult(result CheckResult) {
	fmt.Println(result.Name + ":", result.Status)
	if result.Err != nil {
		fmt.Println(result.Name+" error:", result.Err)
	}
}

func main() {
	fmt.Println("=== opspeek ===")

	printResult(checkDocker())
	printResult(checkKubernetes())
	printResult(checkTailscale())
}
