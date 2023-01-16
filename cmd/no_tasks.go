package cmd

import (
	"github.com/gookit/color"
	"github.com/spf13/cobra"
	"main/utils"
)

// This is a hack add a hint to cobra rendering
func buildNoTaskCommand(reason string) *cobra.Command {
	return &cobra.Command{
		Short:                 color.Magenta.Text("\n  " + reason + "\n"),
		DisableFlagParsing:    true,
		DisableAutoGenTag:     true,
		DisableFlagsInUseLine: true,
		DisableSuggestions:    true,
		GroupID:               utils.GROUP_ID_TASKS,
		Run: func(cmd *cobra.Command, args []string) {
			// if Run is not present the command will be listed
			// in the "additional commands" section of cobra
			// DO NOTHING
		},
	}
}
