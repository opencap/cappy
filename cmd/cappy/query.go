package cappy

import (
	"fmt"
	"github.com/opencap/cappy/pkg/api"
	"github.com/opencap/cappy/pkg/formats"
	"github.com/spf13/cobra"
)

var queryCommand = &cobra.Command{
	Use:     "query",
	Aliases: []string{""},
	Short:   "Query aliases",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println(args)

		client, err := api.NewClient("localhost", 41145)
		if err != nil {
			return fmt.Errorf("client creation failed: %v", err)
		}

		subTypeId, addrData, _, err := client.Lookup("aka.cash", "alice", 0)
		if err != nil {
			return fmt.Errorf("lookup failed: %v", err)
		}

		for _, addr := range formats.Format(addrData, 0, subTypeId) {
			fmt.Printf(" - %s\n", addr)
		}

		return nil
	},
}
