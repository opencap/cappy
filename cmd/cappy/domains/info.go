package domains

import (
	"fmt"
	"github.com/opencap/cappy/internal/cmd/domains"
	"github.com/spf13/cobra"
)

var infoCommand = &cobra.Command{
	Use:   "info",
	Short: "Display information about a domain",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return fmt.Errorf("invalid number of arguments")
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		var (
			domain *string
		)

		if len(args) > 0 {
			domain = &args[0]
		}

		return domains.Info(domain)
	},
}
