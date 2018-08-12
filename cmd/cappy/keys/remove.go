package keys

import (
	"github.com/opencap/cappy/internal/cmd/keys"
	"github.com/spf13/cobra"
)

var removeCommand = &cobra.Command{
	Use:     "remove",
	Aliases: []string{"rm", "delete", "del"},
	Short:   "Delete a key",
	RunE: func(cmd *cobra.Command, args []string) error {
		var id *string

		if len(args) > 0 {
			id = &args[0]
		}

		return keys.Remove(id)
	},
}
