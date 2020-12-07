package plugin

import (
    "bytes"
    "fmt"
    "io/ioutil"
    "os"
    "path/filepath"
    "text/template"

    "github.com/spf13/cobra"
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

            if _, err := os.Stat(registerFile); err == nil {
                return appendImport(registerFile, importPath)
            }

            f, err := os.Create(registerFile)
            if err != nil {
                return err
            }

            t, err := template.New("register").Parse(registerTemplate)
            if err != nil {
                return err
            }

            // TODO: determine import path from git metadata etc.
            err = t.Execute(f, importPath)
            if err != nil {
                return err
            }

            return nil
        },
    }
    registerCmd.Flags().StringVar(&importPath, "import", "", "Import path of plugin.")
    registerCmd.MarkFlagRequired("import")
    registerCmd.Flags().StringVar(&srcDir, "source", "client", "Path to kn source dir. You can fetch source with `kn-builder fetch`.")

    return registerCmd
}

func appendImport(file, importPath string) error {
    content, err := ioutil.ReadFile(file)
    if err != nil {
        return err
    }
    hook := "// Add #plugins# import here. Don't remove this line, it triggers an automatic replacement."
    content = bytes.Replace(content, []byte(hook), []byte(fmt.Sprintf("%s\n    _ \"%s\"", hook, importPath)), 1)
    return ioutil.WriteFile(file, content, 644)
}

var registerTemplate =`
// Copyright © 2020 The OpenShift Knative Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package root

import (
    // Add #plugins# import here. Don't remove this line, it triggers an automatic replacement.
    _ "{{.}}"
)

// RegisterInlinePlugins is an empty function which however forces the
// compiler to run all init() methods of the registered imports
func RegisterInlinePlugins() {}`
