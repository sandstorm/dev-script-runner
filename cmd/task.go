package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"main/utils"
	"os"
	"path/filepath"
	"syscall"
)

func buildTaskCommand(task utils.DevScriptTask, devScriptPath string) *cobra.Command {
	var cmd = &cobra.Command{
		Use:   task.Name,
		Short: task.Comments,
		Long:  task.Comments,
		Args:  cobra.ArbitraryArgs,
		Run: func(cmd *cobra.Command, args []string) {
			execDevScriptWithArguments(devScriptPath, append([]string{task.Name}, args...))
		},
	}
	return cmd
}

func execDevScriptWithArguments(devScriptPath string, arguments []string) {
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
}
