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

package jwt

import (
    // stdlib
    "fmt"
    "io/ioutil"
    "os"
    "path/filepath"
    "strings"
)

type JwtPlugin struct {
    documentationUrl  string
    importLine        string
    pluginDescription string
    pluginName        string
    sourcesUrl        string
}

// Get plugin description. Short one :).
func (jp *JwtPlugin) GetPluginDescription() string {
    return jp.pluginDescription
}

// Get plugin documentation URL.
func (jp *JwtPlugin) GetPluginDocumentationURL() string {
    return jp.documentationUrl
}

// Get plugin's import line.
// This line will be used to replace plugins initialization placehodler
// in Caddy's run.go.
func (jp *JwtPlugin) GetPluginImportLine() string {
    return jp.importLine
}

// Get plugin name.
func (jp *JwtPlugin) GetPluginName() string {
    return jp.pluginName
}

// Get plugin's sources URL, for using with Builder.
func (jp *JwtPlugin) GetPluginSourcesURL() string {
    return jp.sourcesUrl
}

// Plugin initialization.
func (jp *JwtPlugin) Initialize() {
    jp.pluginName = "jwt"
    jp.pluginDescription = "This middleware implements an authorization layer based on JSON Web Tokens (JWT)."
    jp.sourcesUrl = "https://github.com/BTBurke/caddy-jwt"
    jp.documentationUrl = "https://caddyserver.com/docs/jwt"
    jp.importLine = "github.com/BTBurke/caddy-jwt"
}

// Installation
func (jp *JwtPlugin) Install(workspace_path string) {
    // Do nothing if user don't want to install this plugin.
    if !ctx.Flags.BUILD_WITH_JWT {
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

    ctx.Log.Printf("Installing plugin: %s", jp.GetPluginName())

    err1 := ctx.CmdWorker.Execute(fmt.Sprintf("go get -d -u %s", jp.GetPluginImportLine()))
    if err1 != nil {
        ctx.Log.Fatalf("Failed to get plugin's sources: %s", err1.Error())
    }

    // Replace default "This is where other plugins get plugged in (imported)"
    // line with plugin import.
    replace_to := fmt.Sprintf("_ \"%s\"\n\t// This is where other plugins get plugged in (imported)", jp.GetPluginImportLine())
    fh = strings.Replace(fh, "// This is where other plugins get plugged in (imported)", replace_to, 1)
    // Write file.
    ioutil.WriteFile(rungo, []byte(fh), os.ModePerm)
}
