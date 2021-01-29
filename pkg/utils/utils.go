package utils

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"text/template"
)

func AppendImport(file, importPath string) error {
	if _, err := os.Stat(file); os.IsNotExist(err) {
		f, err := os.Create(file)
		if err != nil {
			return err
		}
		t, err := template.New("register").Parse(registerTemplate)
		if err != nil {
			return err
		}
		return t.Execute(f, importPath)
	}

	content, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}
	hook := "// Add #plugins# import here. Don't remove this line, it triggers an automatic replacement."
	content = bytes.Replace(content, []byte(hook), []byte(fmt.Sprintf("%s\n    _ \"%s\"", hook, importPath)), 1)
	return ioutil.WriteFile(file, content, 644)
}

var registerTemplate = `
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
