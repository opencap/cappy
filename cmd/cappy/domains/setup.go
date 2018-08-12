package domains

import (
	"github.com/opencap/cappy/internal/cmd/domains"
	"github.com/spf13/cobra"
)

var setupCommand = &cobra.Command{
	Use:        "setup",
	SuggestFor: []string{"add", "new", "create"},
	Short:      "Use your own domain with a 3rd party server",
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		var (
			domain, host   *string
			port           *uint16
			newKey, dnssig *bool
		)

		if len(args) > 0 {
			domain = &args[0]
		}

		return domains.Setup(domain, host, port, newKey, dnssig)
	},
}
