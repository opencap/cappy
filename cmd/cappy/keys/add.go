package keys

import (
	"github.com/opencap/cappy/internal/cmd/keys"
	"github.com/spf13/cobra"
)

var addCommand = &cobra.Command{
	Use:     "add",
	Aliases: []string{"new", "create"},
	Short:   "Create a new key",
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		_, err = keys.Add()
		return
	},
}
