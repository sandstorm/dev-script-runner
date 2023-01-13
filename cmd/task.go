package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"main/utils"
	"os"
	"path/filepath"
	"strings"
	"syscall"
)

func buildTaskCommand(task utils.DevScriptTask, devScriptPath string) *cobra.Command {
	var cmd = &cobra.Command{
		// IMPORTANT: never color Use! You will not be able to run the command otherwise.
		Use:     task.Usage,
		Short:   task.Short,
		Long:    task.Long,
		Args:    cobra.ArbitraryArgs,
		GroupID: utils.GROUP_ID_TASKS,
		// If this is true all flags will be passed to the command as arguments.
		DisableFlagParsing: true,
		// Will disable the addition of [flags] to the usage
		// As we currently do not parse comments with a specific syntax to provide
		// useful information. The only flags we support are -h and --help they
		// are always part of the usage template.
		DisableFlagsInUseLine: true,
		Run: func(cmd *cobra.Command, args []string) {
			// We DisableFlagParsing but we still want to use the -h and --help flag.
			// This is why we evaluate the args manually and show the help of the
			// command if needed.
			argsAsString := strings.Join(args, " ")
			if argsAsString == "--help" || argsAsString == "-h" {
				cmd.Help()
			} else {
				execDevScriptWithArguments(devScriptPath, append([]string{task.Usage}, args...))
			}
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
