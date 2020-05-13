package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

// Build a CSV for import to Pivotal Tracker by uncommenting the track you'd like to output

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("Usage: %s local|oss", os.Args[0])
	}
	trackName := os.Args[1]

	var modules []string
	switch trackName {
	case "local":
		modules = []string{
			// Onboarding Part 1: Local Track / developer experience / step-by-step
			"introduction_local.prolific",
			"deploy_local.prolific",
			"app_development_local.prolific",
			"routing_deep_dive.prolific",
			"services_redis_cups.prolific",
			"authorization.prolific",
			"networking_monitoring.prolific",
			"concourse.prolific",
			"opsman.prolific",
			"projects.prolific",
		}
	case "oss":
		modules = []string{
			//// Onboarding Part 2: OSS Track / operator & CF developer experience / guided exploration
			"introduction.prolific",
			"deploy_oss.prolific",
			"app_development.prolific",
			"feedback.prolific",
			"routing_deep_dive.prolific",
			"services_redis_cups.prolific",
			"networking_monitoring.prolific",
			"bosh_troubleshooting.prolific",
			"concourse.prolific",
			"opsman.prolific",
			"projects.prolific",
			"clean_up.prolific",
		}
	default:
		log.Fatalf("Unrecognized track type '%s', only 'local' and 'oss' are support", trackName)
	}

	workingDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	outputFile, err := os.Create(filepath.Join(workingDir, "onboarding-tracker.csv"))
	if err != nil {
		panic(err)
	}

	defer outputFile.Close()

	for i, module := range modules {
		moduleFilePath := filepath.Join(workingDir, module)
		cmd := exec.Command("prolific", moduleFilePath)
		cmd.Stderr = os.Stderr
		csvContent, err := cmd.Output()
		if err != nil {
			log.Fatalf("Failed to run prolific: %s", err)
		}

		if i != 0 {
			csvContent = bytes.TrimLeft(csvContent, "Title, Type, Description, Labels,Task\n")
		}

		_, err = outputFile.Write(csvContent)
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Printf("Generating CSV with selected modules: %s", modules)
}
