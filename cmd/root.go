package cmd

import (
	"fmt"
	"github.com/logrusorgru/aurora/v3"
	"github.com/spf13/cobra"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"time"
)

const DEV_SCRIPT_MARKER = "DEV_SCRIPT_MARKER"
const DEV_SCRIPT_NAME = "dev.sh"
const MAX_DEPTH = 100
const INIT_ARGUMENT = "INIT"
const LOGGING_PREFIX = "DSR - " // DevScriptRunner
const LOGGING_WSPACE = "      "
const LOGGING_BAR = "-------------------------------------------------------"

var rootCmd = &cobra.Command{
	Long:    "LONG",
	Use:     "drydock",
	Short:   "SHORT",
	Example: "",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute(version, commit string) {
	rootCmd.AddCommand(buildInitCommand())
	addDevScriptTasksAsCommands(rootCmd)
	rand.Seed(time.Now().UnixNano())

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize()
}

func addDevScriptTasksAsCommands(rootCmd *cobra.Command) {
	currentDirectory, err := os.Getwd()
	if err != nil {
		log.Fatalf("Failed to execute: '%s'", err.Error())
	}

	steps := 0
	for {
		devScriptPath := filepath.Join(currentDirectory, DEV_SCRIPT_NAME)
		if fileExists(devScriptPath) {
			fmt.Println(aurora.Magenta("Found " + currentDirectory + "/" + DEV_SCRIPT_NAME))
			if fileContains(devScriptPath, DEV_SCRIPT_MARKER) {
				fmt.Println(aurora.Magenta("with '" + DEV_SCRIPT_MARKER + "'"))
				tasks := parseDevScriptTasks(devScriptPath)
				for _, task := range tasks {
					rootCmd.AddCommand(buildTaskCommand(task, devScriptPath))
				}
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

type DevScriptTask struct {
	name     string
	comments string
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
