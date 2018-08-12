package keys

import (
	"github.com/spf13/cobra"
)

var RootCommand = &cobra.Command{
	Use:   "keys",
	Short: "Key related commands",
}

func init() {
	RootCommand.AddCommand(addCommand)
	RootCommand.AddCommand(removeCommand)
	RootCommand.AddCommand(listCommand)
}
