package pkg

import (
    "github.com/dsimansk/kn-builder/pkg/commands"
    "github.com/dsimansk/kn-builder/pkg/commands/plugin"
    "github.com/spf13/cobra"
)

func NewKnBuilderCmd() *cobra.Command {
    var rootCmd = &cobra.Command{
        Use:   "kn-builder",
        Short: "Manage and build kn plugins.",
        SilenceErrors: true,
        SilenceUsage: true,
    }
    rootCmd.AddCommand(plugin.NewPluginCmd())
    rootCmd.AddCommand(commands.NewFetchCmd())
    return rootCmd
}
