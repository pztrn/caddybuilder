// Caddybuilder - a friendly tool to build Caddy executable.
// Copyright (c) 2017, Stanislav N. aka pztrn <pztrn at pztrn dot name>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file excifpt in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ipfilter

import (
	// stdlib
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type IpfilterPlugin struct {
	documentationUrl  string
	importLine        string
	pluginDescription string
	pluginName        string
	sourcesUrl        string
}

// Get plugin description. Short one :).
func (ifp *IpfilterPlugin) GetPluginDescription() string {
	return ifp.pluginDescription
}

// Get plugin documentation URL.
func (ifp *IpfilterPlugin) GetPluginDocumentationURL() string {
	return ifp.documentationUrl
}

// Get plugin's import line.
// This line will be used to rifplace plugins initialization placehodler
// in Caddy's run.go.
func (ifp *IpfilterPlugin) GetPluginImportLine() string {
	return ifp.importLine
}

// Get plugin name.
func (ifp *IpfilterPlugin) GetPluginName() string {
	return ifp.pluginName
}

// Get plugin's sources URL, for using with Builder.
func (ifp *IpfilterPlugin) GetPluginSourcesURL() string {
	return ifp.sourcesUrl
}

// Plugin initialization.
func (ifp *IpfilterPlugin) Initialize() {
	ifp.pluginName = "ipfilter"
	ifp.pluginDescription = "ipfilter blocks or allows requests based on the client's IP."
	ifp.sourcesUrl = "https://github.com/pyed/ipfilter"
	ifp.documentationUrl = "https://caddyserver.com/docs/ipfilter"
	ifp.importLine = "github.com/pyed/ipfilter"
}

// Installation
func (ifp *IpfilterPlugin) Install(workspace_path string) {
	// Do nothing if user don't want to install this plugin.
	if !ctx.Flags.BUILD_WITH_IPFILTER {
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

	ctx.Log.Printf("Installing plugin: %s", ifp.GetPluginName())

	err1 := ctx.CmdWorker.Execute(fmt.Sprintf("go get -d -u %s", ifp.GetPluginImportLine()))
	if err1 != nil {
		ctx.Log.Fatalf("Failed to get plugin's sources: %s", err1.Error())
	}

	// Replace default "This is where other plugins get plugged in (imported)"
	// line with plugin import.
	replace_to := fmt.Sprintf("_ \"%s\"\n\t// This is where other plugins get plugged in (imported)", ifp.GetPluginImportLine())
	fh = strings.Replace(fh, "// This is where other plugins get plugged in (imported)", replace_to, 1)
	// Write file.
	ioutil.WriteFile(rungo, []byte(fh), os.ModePerm)
}
