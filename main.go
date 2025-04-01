package main

import (
	"log"
	"os"
	"os/exec"
)

func runTerraform(args ...string) error {
	cmd := exec.Command("terraform", args...)
	cmd.Dir = "terraform"
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	
	// Use AWS CLI configuration
	cmd.Env = append(os.Environ(),
		"AWS_SDK_LOAD_CONFIG=1",      // Makes Go SDK use CLI config
		"AWS_PROFILE=default",        // Use your CLI profile name
	)
	
	return cmd.Run()
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: go run main.go [init|plan|apply|destroy]")
	}

	switch os.Args[1] {
	case "init":
		log.Println("Initializing Terraform...")
		if err := runTerraform("init"); err != nil {
			log.Fatal("Init failed:", err)
		}
	case "plan":
		log.Println("Generating plan...")
		if err := runTerraform("plan"); err != nil {
			log.Fatal("Plan failed:", err)
		}
	case "apply":
		log.Println("Applying configuration...")
		if err := runTerraform("apply", "-auto-approve"); err != nil {
			log.Fatal("Apply failed:", err)
		}
	case "destroy":
		log.Println("Destroying infrastructure...")
		if err := runTerraform("destroy", "-auto-approve"); err != nil {
			log.Fatal("Destroy failed:", err)
		}
	default:
		log.Fatal("Unknown command")
	}
}
