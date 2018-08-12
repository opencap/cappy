package domains

import (
	"github.com/spf13/cobra"
)

var RootCommand = &cobra.Command{
	Use:   "domains",
	Short: "Domain related commands",
}

func init() {
	RootCommand.AddCommand(infoCommand)
	RootCommand.AddCommand(setupCommand)
	RootCommand.AddCommand(notifyCommand)
}
