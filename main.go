// main.go
package main

import (
	"embed"
	"fmt"
	"github.com/logrusorgru/aurora/v3"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
)

const DEV_SCRIPT_MARKER = "DEV_SCRIPT_MARKER"
const DEV_SCRIPT_NAME = "dev.sh"
const MAX_DEPTH = 100
const INIT_ARGUMENT = "INIT"
const LOGGING_PREFIX = "DSR - "
const LOGGING_WSPACE = "      "
const LOGGING_BAR = "-------------------------------------------------------"

// Assets represents the embedded files.
//
//go:embed templates/dev.sh templates/dev_setup.sh
var Assets embed.FS

func main() {
	// `os.Args[0]` will always be the path of the script
	// `os.Args[1]` will either be INIT or a task of the dev.sh script
	if len(os.Args) > 1 {
		if os.Args[1] == INIT_ARGUMENT {
			// If we called `dev INIT`
			fmt.Println(aurora.Magenta(aurora.Bold(LOGGING_BAR)))
			fmt.Println(aurora.Magenta(aurora.Bold(LOGGING_PREFIX + "Running 'dev INIT' ...")))
			runInit()
		} else {
			// If we called `dev <sometask>`
			fmt.Println(aurora.Magenta(aurora.Bold(LOGGING_BAR)))
			fmt.Println(aurora.Magenta(aurora.Bold(LOGGING_PREFIX + "Running 'dev " + strings.Join(os.Args[1:], " ") + "' ...")))
			runTask()
		}
	} else {
		fmt.Println(aurora.Bold("USAGE:"))
		fmt.Println(" ", aurora.Bold("dev <sometask>"), "- to run a task of your dev.sh")
		fmt.Println(" ", aurora.Bold("dev INIT"), "- to create a `dev.sh` in the current folder")
		os.Exit(0)
	}
}

func runInit() {
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
			fmt.Println(aurora.Yellow(LOGGING_PREFIX + "dev_setup.sh is already present!"))
		}
	} else {
		fmt.Println(aurora.Yellow(LOGGING_PREFIX + "dev.sh is already present."))
		fmt.Println(aurora.Yellow(aurora.Bold(LOGGING_PREFIX + "Skipping INIT!")))
	}
	fmt.Println(aurora.Magenta(aurora.Bold(LOGGING_BAR)))
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
			fmt.Println(aurora.Magenta(LOGGING_PREFIX + "Found " + DEV_SCRIPT_NAME + " in "))
			fmt.Println(aurora.Magenta(LOGGING_WSPACE + currentDirectory))
			if fileContains(devScriptPath, DEV_SCRIPT_MARKER) {
				fmt.Println(aurora.Magenta(LOGGING_PREFIX + "Found marker in "))
				fmt.Println(aurora.Magenta(LOGGING_WSPACE + currentDirectory + "/" + DEV_SCRIPT_NAME))
				execDevScriptWithArguments(devScriptPath, os.Args[1:])
				break
			} else {
				fmt.Println(aurora.Yellow(LOGGING_PREFIX + "Marker '" + DEV_SCRIPT_MARKER + "' is missing in "))
				fmt.Println(aurora.Yellow(LOGGING_WSPACE + currentDirectory + "/" + DEV_SCRIPT_NAME))
				fmt.Println(aurora.Yellow(aurora.Bold(LOGGING_PREFIX + "Nothing to do here :(")))
				break
			}
		} else {
			// Moving up one directory
			currentDirectory = filepath.Dir(currentDirectory)
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
	fmt.Println(aurora.Magenta(LOGGING_PREFIX + targetPath + " was created."))
}

func execDevScriptWithArguments(devScriptPath string, arguments []string) {
	fmt.Println(aurora.Magenta(aurora.Bold(LOGGING_PREFIX + "Executing dev script :)")))
	fmt.Println(aurora.Magenta(aurora.Bold(LOGGING_BAR)))
	err := os.Chdir(filepath.Dir(devScriptPath))
	if err != nil {
		log.Fatalf("Failed to execute: '%s'", err.Error())
	}

	err = os.Chmod(devScriptPath, 0755)
	if err != nil {
		log.Fatalf("Failed to execute: '%s'", err.Error())
	}

	cmd := exec.Command(devScriptPath, arguments...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	// Run wil wait for the process to finish
	err = cmd.Run()
	if err != nil {
		log.Fatalf("Failed to run shell script: '%s'", err.Error())
	}
}
