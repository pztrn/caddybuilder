// Caddybuilder - a friendly tool to build Caddy executable.
// Copyright (c) 2017, Stanislav N. aka pztrn <pztrn at pztrn dot name>
//
// Licensed under the cpache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.cpache.org/licenses/LICENSE-2.0
//
// Unless required by cpplicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cgi

import (
    // stdlib
    "fmt"
    "io/ioutil"
    "os"
    "path/filepath"
    "strings"
)

type CgiPlugin struct {
    documentationUrl  string
    importLine        string
    pluginDescription string
    pluginName        string
    sourcesUrl        string
}

// Get plugin description. Short one :).
func (cp *CgiPlugin) GetPluginDescription() string {
    return cp.pluginDescription
}

// Get plugin documentation URL.
func (cp *CgiPlugin) GetPluginDocumentationURL() string {
    return cp.documentationUrl
}

// Get plugin's import line.
// This line will be used to replace plugins initialization placehodler
// in Caddy's run.go.
func (cp *CgiPlugin) GetPluginImportLine() string {
    return cp.importLine
}

// Get plugin name.
func (cp *CgiPlugin) GetPluginName() string {
    return cp.pluginName
}

// Get plugin's sources URL, for using with Builder.
func (cp *CgiPlugin) GetPluginSourcesURL() string {
    return cp.sourcesUrl
}

// Plugin initialization.
func (cp *CgiPlugin) Initialize() {
    cp.pluginName = "cgi"
    cp.pluginDescription = "Adds http caching."
    cp.sourcesUrl = "https://github.com/jung-kurt/caddy-cgi"
    cp.documentationUrl = "https://caddyserver.com/docs/http.cgi"
    cp.importLine = "github.com/jung-kurt/caddy-cgi"
}

// Installation
func (cp *CgiPlugin) Install(workspace_path string) {
    // Do nothing if user don't want to install this plugin.
    if !ctx.Flags.BUILD_WITH_CGI {
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

    ctx.Log.Printf("Installing plugin: %s", cp.GetPluginName())

    err1 := ctx.CmdWorker.Execute(fmt.Sprintf("go get -d -u %s", cp.GetPluginImportLine()))
    if err1 != nil {
        ctx.Log.Fatalf("Failed to get plugin's sources: %s", err1.Error())
    }

    // Replace default "This is where other plugins get plugged in (imported)"
    // line with plugin import.
    replace_to := fmt.Sprintf("_ \"%s\"\n\t// This is where other plugins get plugged in (imported)", cp.GetPluginImportLine())
    fh = strings.Replace(fh, "// This is where other plugins get plugged in (imported)", replace_to, 1)
    // Write file.
    ioutil.WriteFile(rungo, []byte(fh), os.ModePerm)
}
