// main.go
package main

import (
	"embed"
	"github.com/logrusorgru/aurora/v3"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
)

const DEV_SCRIPT_MARKER = "DEV_SCRIPT_MARKER"
const DEV_SCRIPT_NAME = "dev.sh"
const MAX_DEPTH = 100
const INIT_ARGUMENT = "INIT"

// Assets represents the embedded files.
//
//go:embed templates/dev.sh templates/dev_setup.sh
var Assets embed.FS

func main() {
	// TODO: remove later
	// os.Chdir("/Users/florian/src/solarwatt-home-app-flutter/app/lib/features/counter_example")
	os.Chdir("/Users/florian/src/go-dev-script-runner/test")
	if os.Args[1] == INIT_ARGUMENT {
		runInit()
	} else {
		runTask()
	}
}

func runInit() {
	log.Println(aurora.Green("RUNNING INIT"))
	currentDirectory, err := os.Getwd()
	if err != nil {
		log.Fatalf("Failed to execute: '%s'", err.Error())
	}

	devShTargetPath := filepath.Join(currentDirectory, "dev.sh")
	devSetupShTargetPath := filepath.Join(currentDirectory, "dev_setup.sh")

	if !fileExists(devShTargetPath) {
		copyAssetToPath("templates/dev.sh", devShTargetPath)
		if !fileExists(devSetupShTargetPath) {
			copyAssetToPath("templates/dev_setup.sh", devSetupShTargetPath)
		} else {
			log.Println("dev_setup.sh is already present.")
		}
	} else {
		log.Println("dev.sh is already present. Skipping INIT")
	}

	os.Exit(0)
}

func runTask() {
	currentDirectory, err := os.Getwd()
	if err != nil {
		log.Fatalf("Failed to execute: '%s'", err.Error())
	}

	steps := 0
	for {
		devScriptPath := filepath.Join(currentDirectory, DEV_SCRIPT_NAME)
		if fileExists(devScriptPath) {
			log.Println("Found " + DEV_SCRIPT_NAME + " in path " + currentDirectory)
			if fileContains(devScriptPath, DEV_SCRIPT_MARKER) {
				log.Println(DEV_SCRIPT_NAME + " contains marker " + DEV_SCRIPT_MARKER)
				execDevScriptWithArguments(devScriptPath, os.Args[1:])
				break
			} else {
				log.Println("Marker is missing in " + DEV_SCRIPT_NAME)
			}
		} else {
			currentDirectory = filepath.Dir(currentDirectory)
			log.Println("MOVING UP:", currentDirectory)
			steps += 1
		}
		if currentDirectory == "/" || steps >= MAX_DEPTH {
			break
		}
	}
}

func fileExists(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func fileContains(filePath string, needle string) bool {
	b, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Failed to execute: '%s'", err.Error())
	}
	isContained, err := regexp.Match(needle, b)
	if err != nil {
		log.Fatalf("Failed to execute: '%s'", err.Error())
	}
	return isContained
}

func copyAssetToPath(embedPath string, targetPath string) {
	data, err := Assets.ReadFile(embedPath)
	if err != nil {
		log.Fatalf("Failed to execute: '%s'", err.Error())
	}
	err = os.WriteFile(targetPath, []byte(data), 0755)
	if err != nil {
		log.Fatalf("Failed to execute: '%s'", err.Error())
	}
	log.Println(targetPath + " was created.")
}

func execDevScriptWithArguments(devScriptPath string, arguments []string) {
	err := os.Chdir(filepath.Dir(devScriptPath))
	if err != nil {
		log.Fatalf("Failed to execute: '%s'", err.Error())
	}

	cmd := exec.Command(devScriptPath, arguments...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	// Run wil wait for the process to finish
	err = cmd.Run()
	if err != nil {
		log.Fatalf("Failed to execute: '%s'", err.Error())
	}
}
