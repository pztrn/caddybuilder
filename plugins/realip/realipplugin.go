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

package realip

import (
	// stdlib
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type RealipPlugin struct {
	documentationUrl  string
	importLine        string
	pluginDescription string
	pluginName        string
	sourcesUrl        string
}

// Get plugin description. Short one :).
func (rip *RealipPlugin) GetPluginDescription() string {
	return rip.pluginDescription
}

// Get plugin documentation URL.
func (rip *RealipPlugin) GetPluginDocumentationURL() string {
	return rip.documentationUrl
}

// Get plugin's import line.
// This line will be used to replace plugins initialization placehodler
// in Caddy's run.go.
func (rip *RealipPlugin) GetPluginImportLine() string {
	return rip.importLine
}

// Get plugin name.
func (rip *RealipPlugin) GetPluginName() string {
	return rip.pluginName
}

// Get plugin's sources URL, for using with Builder.
func (rip *RealipPlugin) GetPluginSourcesURL() string {
	return rip.sourcesUrl
}

// Plugin initialization.
func (rip *RealipPlugin) Initialize() {
	rip.pluginName = "realip"
	rip.pluginDescription = "The real IP module is useful in situations where you are running caddy behind a proxy server. In these situations, the actual client IP will be stored in an HTTP header, usually X-Forwarded-For."
	rip.sourcesUrl = "https://github.com/captncraig/caddy-realip"
	rip.documentationUrl = "https://caddyserver.com/docs/realip"
	rip.importLine = "github.com/captncraig/caddy-realip"
}

// Installation
func (rip *RealipPlugin) Install(workspace_path string) {
	// Do nothing if user don't want to install this plugin.
	if !ctx.Flags.BUILD_WITH_REALIP {
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

	ctx.Log.Printf("Installing plugin: %s", rip.GetPluginName())

	err1 := ctx.CmdWorker.Execute(fmt.Sprintf("go get -d -u %s", rip.GetPluginImportLine()))
	if err1 != nil {
		ctx.Log.Fatalf("Failed to get plugin's sources: %s", err1.Error())
	}

	// Replace default "This is where other plugins get plugged in (imported)"
	// line with plugin import.
	replace_to := fmt.Sprintf("_ \"%s\"\n\t// This is where other plugins get plugged in (imported)", rip.GetPluginImportLine())
	fh = strings.Replace(fh, "// This is where other plugins get plugged in (imported)", replace_to, 1)
	// Write file.
	ioutil.WriteFile(rungo, []byte(fh), os.ModePerm)
}
