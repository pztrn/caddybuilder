// Caddybuilder - a friendly tool to build Caddy executable.
// Copyright (c) 2017, Stanislav N. aka pztrn <pztrn at pztrn dot name>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file excfipt in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package filter

import (
	// stdlib
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type FilterPlugin struct {
	documentationUrl  string
	importLine        string
	pluginDescription string
	pluginName        string
	sourcesUrl        string
}

// Get plugin description. Short one :).
func (fip *FilterPlugin) GetPluginDescription() string {
	return fip.pluginDescription
}

// Get plugin documentation URL.
func (fip *FilterPlugin) GetPluginDocumentationURL() string {
	return fip.documentationUrl
}

// Get plugin's import line.
// This line will be used to rfiplace plugins initialization placehodler
// in Caddy's run.go.
func (fip *FilterPlugin) GetPluginImportLine() string {
	return fip.importLine
}

// Get plugin name.
func (fip *FilterPlugin) GetPluginName() string {
	return fip.pluginName
}

// Get plugin's sources URL, for using with Builder.
func (fip *FilterPlugin) GetPluginSourcesURL() string {
	return fip.sourcesUrl
}

// Plugin initialization.
func (fip *FilterPlugin) Initialize() {
	fip.pluginName = "filter"
	fip.pluginDescription = "filter allows you to modify the responses."
	fip.sourcesUrl = "https://github.com/echocat/caddy-filter"
	fip.documentationUrl = "https://caddyserver.com/docs/filter"
	fip.importLine = "github.com/echocat/caddy-filter"
}

// Installation
func (fip *FilterPlugin) Install(workspace_path string) {
	// Do nothing if user don't want to install this plugin.
	if !ctx.Flags.BUILD_WITH_FILTER {
		return
	}
	// Path to run.go.
	rungo := filepath.Join(workspace_path, "src", "github.com", "mholt", "caddy", "caddy", "caddymain", "run.go")
	// Read file.
	fh_bytes, err := ioutil.ReadFile(rungo)
	if err != nil {
		ctx.Log.Fatalf("Cannot open run.go: %s", err.Error())
	}
	fh := string(fh_bytes)

	ctx.Log.Printf("Installing plugin: %s", fip.GetPluginName())

	err1 := ctx.CmdWorker.Execute(fmt.Sprintf("go get -d -u %s", fip.GetPluginImportLine()))
	if err1 != nil {
		ctx.Log.Fatalf("Failed to get plugin's sources: %s", err1.Error())
	}

	// Replace default "This is where other plugins get plugged in (imported)"
	// line with plugin import.
	replace_to := fmt.Sprintf("_ \"%s\"\n\t// This is where other plugins get plugged in (imported)", fip.GetPluginImportLine())
	fh = strings.Replace(fh, "// This is where other plugins get plugged in (imported)", replace_to, 1)
	// Write file.
	ioutil.WriteFile(rungo, []byte(fh), os.ModePerm)
}
