package cmd

import (
	"fmt"
	"github.com/gookit/color"
	"github.com/spf13/cobra"
	"log"
	"main/utils"
	"os"
	"path/filepath"
	"strconv"
)

var RootCmd = &cobra.Command{
	Use: "dev",
	Long: `
DevScriptRunner is a helper to run task from a dev.sh file
from within a nested folder structure also providing autocompletion
and other nifty feature ;)
`,
	Short:   "SHORT",
	Example: "",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute(version, commit string) {
	// IMPORTANT: the order of tasks should resemble the order in the dev.sh
	cobra.EnableCommandSorting = false
	RootCmd.AddCommand(buildSectionCommand("your tasks"))
	addDevScriptTasksAsCommands(RootCmd)
	RootCmd.AddCommand(buildSectionCommand("utils"))
	RootCmd.AddCommand(buildInitCommand())
	// RootCmd.AddCommand(buildCompletionCommand())
	RootCmd.AddCommand(buildSectionCommand("other"))
	if err := RootCmd.Execute(); err != nil {
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
		devScriptPath := filepath.Join(currentDirectory, utils.DEV_SCRIPT_NAME)
		if utils.FileExists(devScriptPath) {
			if utils.FileContains(devScriptPath, utils.DEV_SCRIPT_MARKER) {
				tasks := utils.ParseDevScriptTasks(devScriptPath)
				rootCmd.Long = rootCmd.Long + color.Magenta.Text("\nDEV SCRIPT WITH ")
				rootCmd.Long = rootCmd.Long + color.Bold.Text(strconv.Itoa(len(tasks)))
				rootCmd.Long = rootCmd.Long + color.Magenta.Text(" TASKS FOUND AT:\n  ")
				rootCmd.Long = rootCmd.Long + color.Magenta.Text(devScriptPath)
				if len(tasks) > 0 {
					for _, task := range tasks {
						rootCmd.AddCommand(buildTaskCommand(task, devScriptPath))
					}
				} else {
					rootCmd.AddCommand(buildNoTaskCommand("NO TASKS IN DEV SCRIPT!"))
				}
				break
			}
		}
		if currentDirectory == "/" || steps >= utils.MAX_DEPTH {
			rootCmd.AddCommand(buildNoTaskCommand("NO DEV SCRIPT WITH VALID MARKER FOUND!"))
			break
		}
		// Moving up one directory
		currentDirectory = filepath.Dir(currentDirectory)
		steps += 1
	}
}
