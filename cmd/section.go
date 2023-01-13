package cmd

import (
	"github.com/gookit/color"
	"github.com/spf13/cobra"
	"main/utils"
)

// This is a hack to add a section to cobra rendering
func buildSectionTitle(title string) string {
	title = " " + title + " "
	beforeAndAfterSeparatorLength := (len(utils.SECTION_SEPARATOR) - len(title)) / 2
	remainingSeparator := utils.SECTION_SEPARATOR[0 : beforeAndAfterSeparatorLength-1]
	return "\n" + remainingSeparator + title + remainingSeparator + "\n"
}

func buildSectionCommand(title string) *cobra.Command {
	return &cobra.Command{
		Short:                 color.Gray.Text(buildSectionTitle(title)),
		DisableFlagParsing:    true,
		DisableAutoGenTag:     true,
		DisableFlagsInUseLine: true,
		DisableSuggestions:    true,
		Run: func(cmd *cobra.Command, args []string) {
			// if Run is not present the command will be listed
			// in the "additional commands" section of cobra
			// DO NOTHING
		},
	}
}
