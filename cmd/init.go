package cmd

import (
	"github.com/gookit/color"
	"github.com/spf13/cobra"
)

func buildInitCommand() *cobra.Command {
	var execRootCmd = &cobra.Command{
		Use:   "INIT",
		Short: "TODO",
		Long:  color.Sprintf(`Usage:	TODO`),
		Args:  cobra.RangeArgs(1, 2),

		Run: func(cmd *cobra.Command, args []string) {
			color.Sprintf(`RUN COMMAND`)
		},
	}
	return execRootCmd
}
