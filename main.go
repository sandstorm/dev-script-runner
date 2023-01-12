// main.go
package main

import (
	"embed"
	"fmt"
	"github.com/logrusorgru/aurora/v3"
	"log"
	"main/cmd"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"syscall"
)

const DEV_SCRIPT_MARKER = "DEV_SCRIPT_MARKER"
const DEV_SCRIPT_NAME = "dev.sh"
const MAX_DEPTH = 100
const INIT_ARGUMENT = "INIT"
const LOGGING_PREFIX = "DSR - " // DevScriptRunner
const LOGGING_WSPACE = "      "
const LOGGING_BAR = "-------------------------------------------------------"

// Assets represents the embedded files.
//
//go:embed templates/dev.sh templates/dev_setup.sh
var Assets embed.FS

// set by goreleaser; see https://goreleaser.com/environment/
var (
	version = "dev"
	commit  = "none"
)

func main() {
	cmd.Execute(version, commit)
}

func main2() {
	// `os.Args[0]` will always be the path of the script
	// `os.Args[1]` will either be INIT or a name of the dev.sh script
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
		// Show usage
		fmt.Println(aurora.Bold("USAGE:"))
		fmt.Println(" ", aurora.Bold("dev <TASK_NAME>"), "- to run a name of your dev.sh")
		fmt.Println(" ", aurora.Bold("dev INIT"), "- to create a `dev.sh` in the current folder")
		// Explicitly show that we exit here
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
		// we can access embedded assets by using the path use din the annotation
		copyAssetToPath("templates/dev.sh", devShTargetPath)
		if !fileExists(devSetupShTargetPath) {
			// We do not want to add dev_setup.sh if INIT was already run.
			// The file might have been deleted on purpose.
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
			fmt.Println(aurora.Magenta(LOGGING_PREFIX + "Found " + currentDirectory + "/" + DEV_SCRIPT_NAME))
			if fileContains(devScriptPath, DEV_SCRIPT_MARKER) {
				fmt.Println(aurora.Magenta(LOGGING_WSPACE + "Found marker '" + DEV_SCRIPT_MARKER + "'"))
				execDevScriptWithArguments(devScriptPath, os.Args[1:])
				break
			} else {
				fmt.Println(aurora.Yellow(LOGGING_WSPACE + "Marker '" + DEV_SCRIPT_MARKER + "' is missing in " + DEV_SCRIPT_NAME))
				fmt.Println(aurora.Yellow(LOGGING_WSPACE + "Moving up, looking for new " + DEV_SCRIPT_NAME))
				// Not breaking here as we want to move up
			}
		}
		if currentDirectory == "/" || steps >= MAX_DEPTH {
			fmt.Println(aurora.Yellow(aurora.Bold(LOGGING_PREFIX + "No " + DEV_SCRIPT_NAME + " with " + DEV_SCRIPT_MARKER + " found. Nothing to do here :(")))
			fmt.Println(aurora.Magenta(aurora.Bold(LOGGING_BAR)))
			break
		}
		// Moving up one directory
		currentDirectory = filepath.Dir(currentDirectory)
		steps += 1
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
	markerIsPresent, err := regexp.Match(needle, b)
	if err != nil {
		log.Fatalf("Failed to execute: '%s'", err.Error())
	}
	return markerIsPresent
}

func copyAssetToPath(embedPath string, targetPath string) {
	data, err := Assets.ReadFile(embedPath)
	if err != nil {
		log.Fatalf("Failed to execute: '%s'", err.Error())
	}
	// Shell scripts need to be executable -> 0755
	err = os.WriteFile(targetPath, []byte(data), 0755)
	if err != nil {
		log.Fatalf("Failed to execute: '%s'", err.Error())
	}
	fmt.Println(aurora.Magenta(LOGGING_PREFIX + targetPath + " was created."))
}

type DevScriptTask struct {
	name     string
	comments string
}

func taskExists(tasks []DevScriptTask, calledTask string) bool {
	for _, task := range tasks {
		if task.name == calledTask {
			return true
		}
	}
	return false
}

func availableTasks(tasks []DevScriptTask) string {
	result := ""
	for _, task := range tasks {
		result += LOGGING_WSPACE + "  " + task.name + "\n"
	}
	return result
}

func parseDevScriptTasks(devScriptPath string) []DevScriptTask {
	// https://regex101.com/r/5LVRcP/1 -> Iteration 1 without comments before
	// https://regex101.com/r/XyB410/1 -> Final Iteration with comments before ;)
	devScriptBytes, err := os.ReadFile(devScriptPath)
	if err != nil {
		log.Fatalf("Failed to execute: '%s'", err.Error())
	}

	devScriptString := string(devScriptBytes)
	compiledRegex := regexp.MustCompile(`(?m)(?P<comments>(?:^#.*(?:\n|\r\n|\r))*)^(?:function )?(?P<name>[a-zA-Z0-9_-]+)\s?(?:\(\))?\s?{`)
	captureGroupNames := compiledRegex.SubexpNames()
	commentsIndex := sort.StringSlice(captureGroupNames).Search("comments")
	taskIndex := sort.StringSlice(captureGroupNames).Search("name")
	matches := compiledRegex.FindAllStringSubmatch(devScriptString, -1)

	var results = []DevScriptTask{}
	for _, match := range matches {
		task := match[taskIndex]
		comments := match[commentsIndex]
		if !strings.HasPrefix(task, "_") {
			results = append(results, DevScriptTask{name: task, comments: comments})
		}
	}
	return results
}

func execDevScriptWithArguments(devScriptPath string, arguments []string) {
	tasks := parseDevScriptTasks(devScriptPath)
	calledTask := os.Args[1]
	if taskExists(tasks, calledTask) {
		fmt.Println(aurora.Magenta(LOGGING_WSPACE + "Found task '" + calledTask + "'"))
		fmt.Println(aurora.Magenta(aurora.Bold(LOGGING_PREFIX + "Executing dev script :)")))
		fmt.Println(aurora.Magenta(aurora.Bold(LOGGING_BAR)))
		err := os.Chdir(filepath.Dir(devScriptPath))
		if err != nil {
			log.Fatalf("Failed to execute: '%s'", err.Error())
		}
		// In case the script is not executable
		err = os.Chmod(devScriptPath, 0755)
		if err != nil {
			log.Fatalf("Failed to execute: '%s'", err.Error())
		}

		// We tried using exec.Command(devScriptPath, arguments...) which failed for interactive
		// terminal calls, e.g. `docker compose exec neos /bin/bash` This is running our command
		// as a child process. We are now replacing the process of this helper with the call of
		// the `dev.sh` using `syscall.Exec()`
		err = syscall.Exec(devScriptPath, append([]string{devScriptPath}, arguments...), os.Environ())
		if err != nil {
			log.Fatalf("Failed to run shell script: '%s'", err.Error())
		}
		// IMPORTANT: As we are replacing the process of the helper nothing else will be
		// called, except the error handler!
	} else {
		fmt.Println(aurora.Yellow(aurora.Bold(LOGGING_WSPACE + "No task '" + calledTask + "' found. Nothing to do here :(")))
		fmt.Println(aurora.Yellow(aurora.Bold(LOGGING_WSPACE + "Try one of the following:")))
		fmt.Print(aurora.Yellow(availableTasks(tasks)))
		fmt.Println(aurora.Magenta(aurora.Bold(LOGGING_BAR)))
	}
}
