package canary

import (
	"github.com/spf13/cobra"
	canary_config "github.com/spinnaker/spin/cmd/canary/canary-config"
)

type canaryOptions struct{}

const (
	canaryShort   = ""
	canaryLong    = ""
	canaryExample = ""
)

func NewCanaryCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "canary",
		Aliases: []string{},
		Short:   canaryShort,
		Long:    canaryLong,
		Example: canaryExample,
	}

	// create subcommands
	cmd.AddCommand(canary_config.NewCanaryConfigCmd())
	return cmd
}
