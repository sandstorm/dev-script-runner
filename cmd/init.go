package cmd

import (
	"github.com/gookit/color"
	"github.com/spf13/cobra"
	"log"
	"main/utils"
	"os"
	"path/filepath"
)

func buildInitCommand() *cobra.Command {
	var cmd = &cobra.Command{
		// IMPORTANT: never color Use! You will not be able to run the command otherwise.
		Use:     "DSR_INIT",
		Short:   "creates dev.sh and dev_setup.sh in current folder",
		Long:    "This creates an example create a dev.sh and dev_setup.sh in your current directory.",
		Args:    cobra.NoArgs,
		GroupID: utils.GROUP_ID_UTILS,

		Run: func(cmd *cobra.Command, args []string) {
			currentDirectory, err := os.Getwd()
			if err != nil {
				log.Fatalf("Failed to execute: '%s'", err.Error())
			}
			devShTargetPath := filepath.Join(currentDirectory, "dev.sh")
			devSetupShTargetPath := filepath.Join(currentDirectory, "dev_setup.sh")

			if !utils.FileExists(devShTargetPath) {
				// we can access embedded assets by using the path use din the annotation
				utils.CopyAssetToPath("templates/dev.sh", devShTargetPath)
				if !utils.FileExists(devSetupShTargetPath) {
					// We do not want to add dev_setup.sh if INIT was already run.
					// The file might have been deleted on purpose.
					utils.CopyAssetToPath("templates/dev_setup.sh", devSetupShTargetPath)
				} else {
					color.Yellow.Println("dev_setup.sh is already present!")
				}
			} else {
				color.Yellow.Println("dev.sh is already present.")
				color.Style{color.Yellow, color.Bold}.Println("Skipping INIT!")
			}
			os.Exit(0)
		},
	}
	return cmd
}
