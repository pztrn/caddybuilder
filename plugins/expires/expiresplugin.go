// Caddybuilder - a friendly tool to build Caddy executable.
// Copyright (c) 2017, Stanislav N. aka pztrn <pztrn at pztrn dot name>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package expires

import (
	// stdlib
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type ExpiresPlugin struct {
	documentationUrl  string
	importLine        string
	pluginDescription string
	pluginName        string
	sourcesUrl        string
}

// Get plugin description. Short one :).
func (ep *ExpiresPlugin) GetPluginDescription() string {
	return ep.pluginDescription
}

// Get plugin documentation URL.
func (ep *ExpiresPlugin) GetPluginDocumentationURL() string {
	return ep.documentationUrl
}

// Get plugin's import line.
// This line will be used to replace plugins initialization placehodler
// in Caddy's run.go.
func (ep *ExpiresPlugin) GetPluginImportLine() string {
	return ep.importLine
}

// Get plugin name.
func (ep *ExpiresPlugin) GetPluginName() string {
	return ep.pluginName
}

// Get plugin's sources URL, for using with Builder.
func (ep *ExpiresPlugin) GetPluginSourcesURL() string {
	return ep.sourcesUrl
}

// Plugin initialization.
func (ep *ExpiresPlugin) Initialize() {
	ep.pluginName = "expires"
	ep.pluginDescription = "expires allows you to set expiration header relative to the request time."
	ep.sourcesUrl = "https://github.com/epicagency/caddy-expires"
	ep.documentationUrl = "https://caddyserver.com/docs/expires"
	ep.importLine = "github.com/epicagency/caddy-expires"
}

// Installation
func (ep *ExpiresPlugin) Install(workspace_path string) {
	// Do nothing if user don't want to install this plugin.
	if !ctx.Flags.BUILD_WITH_EXPIRES {
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

	ctx.Log.Printf("Installing plugin: %s", ep.GetPluginName())

	err1 := ctx.CmdWorker.Execute(fmt.Sprintf("go get -d -u %s", ep.GetPluginImportLine()))
	if err1 != nil {
		ctx.Log.Fatalf("Failed to get plugin's sources: %s", err1.Error())
	}

	// Replace default "This is where other plugins get plugged in (imported)"
	// line with plugin import.
	replace_to := fmt.Sprintf("_ \"%s\"\n\t// This is where other plugins get plugged in (imported)", ep.GetPluginImportLine())
	fh = strings.Replace(fh, "// This is where other plugins get plugged in (imported)", replace_to, 1)
	// Write file.
	ioutil.WriteFile(rungo, []byte(fh), os.ModePerm)
}
