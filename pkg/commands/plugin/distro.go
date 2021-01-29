package plugin

import (
	"fmt"
	"github.com/dsimansk/knb/pkg/utils"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
)

const KnConfigFile = ".kn.yaml"

type DistroConfig struct {
	Plugins []Plugin `yaml:"plugins"`
}

func NewDistroGenerateCmd() *cobra.Command {
	var generateCmd = &cobra.Command{
		Use:   "distro",
		Short: "Generate required files to build `kn` with inline plugins.",
		RunE: func(cmd *cobra.Command, args []string) error {
			if _, err := os.Stat(KnConfigFile); os.IsNotExist(err) {
				return fmt.Errorf("kn distro configuration file '%s' doesn't exist, make sure the command is executed in knative/client root directory", KnConfigFile)
			}
			rawConf, err := ioutil.ReadFile(KnConfigFile)
			if err != nil {
				return err
			}
			conf := &DistroConfig{}
			if err := yaml.Unmarshal(rawConf, conf); err != nil {
				return err
			}

			registerFile := filepath.Join("pkg", "kn", "root", "plugin_register.go")

			for _, p := range conf.Plugins {
				importPath := p.PluginImportPath
				if importPath == "" {
					importPath = p.Module + "/plugin"
				}
				if err := utils.AppendImport(registerFile, importPath); err != nil {
					return err
				}
				_, err := exec.Command("go", "mod", "edit", "-require", p.Module+"@"+p.Version).Output()
				if err != nil {
					return fmt.Errorf("go mod edit -require failed: %s", err.Error())
				}

				if len(p.Replace) > 0 {
					for _, r := range p.Replace {
						_, err := exec.Command("go", "mod", "edit", "-replace", r.Module+"="+r.Module+"@"+r.Version).Output()
						if err != nil {
							return fmt.Errorf("go mod edit -replace failed: %s", err.Error())
						}
					}
				}
			}
			if err := exec.Command("gofmt", "-s", "-w", registerFile).Run(); err != nil {
				return fmt.Errorf("gofmt failed: %s", err.Error())
			}
			return nil
		},
	}
	return generateCmd
}
