package domains

import (
	"fmt"
	"github.com/opencap/cappy/internal/cmd/domains"
	"github.com/spf13/cobra"
)

var notifyCommand = &cobra.Command{
	Use:     "notify",
	Aliases: []string{"register", "associate", "update"},
	Short:   "Notify changes to a domain",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) > 1 {
			return fmt.Errorf("invalid number of arguments")
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		var domain *string

		if len(args) > 0 {
			domain = &args[0]
		}

		return domains.Notify(domain)
	},
}
