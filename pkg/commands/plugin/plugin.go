package plugin

import "github.com/spf13/cobra"

type Plugin struct {
	Name             string
	Description      string
	Module           string
	Version          string
	PluginImportPath string
	Replace          []struct {
		Module  string
		Version string
	}
	CmdParts []string
}

func NewPluginCmd() *cobra.Command {
	var pluginCmd = &cobra.Command{
		Use:   "plugin",
		Short: "Manage kn plugins.",
	}
	pluginCmd.AddCommand(NewPluginInitCmd())
	pluginCmd.AddCommand(NewDistroGenerateCmd())
	return pluginCmd
}
