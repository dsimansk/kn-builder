package plugin

import "github.com/spf13/cobra"

type Plugin struct {
    Name string
    Description string
    ImportPath string
    CmdParts []string
}

func NewPluginCmd() *cobra.Command {
    var pluginCmd = &cobra.Command{
        Use:   "plugin",
        Short: "Manage kn plugins.",
    }
    pluginCmd.AddCommand(NewPluginInitCmd())
    pluginCmd.AddCommand(NewPluginRegisterCmd())
    return pluginCmd
}
