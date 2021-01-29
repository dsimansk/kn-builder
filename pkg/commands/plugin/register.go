package plugin

import (
	"fmt"
	"github.com/dsimansk/knb/pkg/utils"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
)

func NewPluginRegisterCmd() *cobra.Command {
	var importPath string
	var srcDir string
	var registerCmd = &cobra.Command{
		Use:   "register",
		Short: "Register plugin to kn.",
		RunE: func(cmd *cobra.Command, args []string) error {
			if _, err := os.Stat(srcDir); os.IsNotExist(err) {
				return fmt.Errorf("plugin dir '%s' doesn't exist", srcDir)
			}
			registerFile := filepath.Join(srcDir, "pkg", "kn", "root", "plugin_register.go")

			return utils.AppendImport(registerFile, importPath)
		},
	}
	registerCmd.Flags().StringVar(&importPath, "import", "", "Import path of plugin.")
	registerCmd.MarkFlagRequired("import")
	registerCmd.Flags().StringVar(&srcDir, "source", "client", "Path to kn source dir. You can fetch source with `kn-builder fetch`.")

	return registerCmd
}
