package cappy

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/mitchellh/go-homedir"
	"github.com/opencap/cappy/cmd/cappy/domains"
	"github.com/opencap/cappy/cmd/cappy/keys"
	"github.com/opencap/cappy/internal/pkg/context"
	"github.com/opencap/cappy/internal/pkg/key"
	"github.com/spf13/cobra"
	"os"
	"path"
)

var (
	dataDir string
)

const keysDir = "./key"

var RootCmd = &cobra.Command{
	Use:           "cappy",
	Short:         "Making crypto convenient",
	Long:          `cappy is a tool to communicate with ...`,
	SilenceErrors: true,
	SilenceUsage:  true,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		if err := os.MkdirAll(dataDir, 0755); err != nil {
			return err
		}

		m, err := key.NewManager(path.Join(dataDir, keysDir))
		if err != nil {
			return fmt.Errorf("failed to create key manager: %v", err)
		}
		context.Instance().SetKeyManager(m)

		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		return queryCommand.RunE(queryCommand, args)
	},
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(color.RedString("%v", err))
		os.Exit(1)
	}
}

func init() {
	defaultDataDir, err := homedir.Dir()
	if err != nil {
		defaultDataDir = "."
	}
	defaultDataDir = path.Join(defaultDataDir, ".opencap")

	RootCmd.PersistentFlags().StringVar(&dataDir, "data-dir", defaultDataDir, "data directory")
}

func init() {
	RootCmd.AddCommand(queryCommand)
	RootCmd.AddCommand(domains.RootCommand)
	RootCmd.AddCommand(keys.RootCommand)
}
