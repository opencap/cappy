package keys

import (
	"github.com/opencap/cappy/internal/cmd/keys"
	"github.com/spf13/cobra"
)

var listCommand = &cobra.Command{
	Use:   "list",
	Short: "List key",
	RunE: func(cmd *cobra.Command, args []string) error {
		return keys.List()
	},
}
