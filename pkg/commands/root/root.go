package root

import (
	"github.com/dsimansk/knb/pkg/commands/plugin"
	"github.com/spf13/cobra"
)

func NewKnBuilderCmd() *cobra.Command {
	var rootCmd = &cobra.Command{
		Use:           "knb",
		Short:         "Manage and build kn inline plugins.",
		SilenceErrors: true,
		SilenceUsage:  true,
	}
	rootCmd.AddCommand(plugin.NewPluginCmd())
	return rootCmd
}
