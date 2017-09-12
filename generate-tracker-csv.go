package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

// Build a CSV for import to Pivotal Tracker by uncommenting the track you'd like to output

func main() {
	modules := []string{
		// Onboarding Part 1: Local Track / developer experience / step-by-step
		"introduction_local.prolific",
		"deploy_local.prolific",
		"app_development_local.prolific",
		"services_redis_cups.prolific",
		"authorization.prolific",
		"networking_monitoring.prolific",
		"concourse_local.prolific",

		//// Onboarding Part 2: OSS Track / operator & CF developer experience / guided exploration
		// "introduction.prolific",
		// "deploy_oss.prolific",
		// "app_development.prolific",
		// "services_redis_cups.prolific",
		// "networking_monitoring.prolific",
		// "bosh_troubleshooting.prolific",
		// "bosh_release_redis.prolific",
		// "concourse.prolific",
		// "clean_up.prolific",
	}

	workingDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	outputFile, err := os.Create(filepath.Join(workingDir, "onboarding-tracker.csv"))
	if err != nil {
		panic(err)
	}

	// outputFile, err = os.OpenFile(outputFile, os.O_APPEND, 0666)
	defer outputFile.Close()

	for i, module := range modules {
		moduleFilePath := filepath.Join(workingDir, module)
		cmd := exec.Command("prolific", moduleFilePath)
		csvContent, err := cmd.Output()
		if err != nil {
			fmt.Println(err.Error())
		}

		if i != 0 {
			csvContent = bytes.TrimLeft(csvContent, "Title, Type, Description, Labels\n")
		}

		_, err = outputFile.Write(csvContent)
		if err != nil {
			panic(err)
		}
	}

	fmt.Printf("Generating CSV with selected modules: %s", modules)
}
