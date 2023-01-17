package cmd

import (
	"github.com/spf13/cobra"
)

// This is a hack add a hint to cobra rendering
func buildTestTaskCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "test",
		Short: "Test Task",
		Run: func(cmd *cobra.Command, args []string) {
		},
	}
}
