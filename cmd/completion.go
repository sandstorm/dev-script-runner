package cmd

import (
	"github.com/spf13/cobra"
	"os"
)

func buildCompletionCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "DSR_COMPLETION [bash|zsh|fish|powershell]",
		Short: "Generate completion script for common shells",
		Long: `Generate Autocomplete Scripts for common shells for the DevScriptRunner tool

To load completions:

Bash:

$ source <(dev DSR_COMPLETION bash)

# To load completions for each session, execute once:
Linux:
  $ dev DSR_COMPLETION bash > /etc/bash_completion.d/dev
MacOS:
  $ dev DSR_COMPLETION bash > /usr/local/etc/bash_completion.d/dev

Zsh:

# If shell completion is not already enabled in your environment you will need
# to enable it.  You can execute the following once:

$ echo "autoload -U compinit; compinit" >> ~/.zshrc

# To load completions for each session, execute once:
$ dev DSR_COMPLETION zsh > "${fpath[1]}/_dev"

# You will need to start a new shell for this setup to take effect.

Fish:

$ dev DSR_COMPLETION fish | source

# To load completions for each session, execute once:
$ dev DSR_COMPLETION fish > ~/.config/fish/completions/dev.fish
`,
		DisableFlagsInUseLine: true,
		ValidArgs:             []string{"bash", "zsh", "fish", "powershell"},
		Args:                  cobra.ExactValidArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			switch args[0] {
			case "bash":
				cmd.Root().GenBashCompletion(os.Stdout)
			case "zsh":
				cmd.Root().GenZshCompletion(os.Stdout)
			case "fish":
				cmd.Root().GenFishCompletion(os.Stdout, true)
			case "powershell":
				cmd.Root().GenPowerShellCompletion(os.Stdout)
			}
		},
	}
}
