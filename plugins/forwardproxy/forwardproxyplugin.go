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

package forwardproxy

import (
    // stdlib
    "fmt"
    "io/ioutil"
    "os"
    "path/filepath"
    "strings"
)

type ForwardproxyPlugin struct {
    documentationUrl  string
    importLine        string
    pluginDescription string
    pluginName        string
    sourcesUrl        string
}

// Get plugin description. Short one :).
func (fp *ForwardproxyPlugin) GetPluginDescription() string {
    return fp.pluginDescription
}

// Get plugin documentation URL.
func (fp *ForwardproxyPlugin) GetPluginDocumentationURL() string {
    return fp.documentationUrl
}

// Get plugin's import line.
// This line will be used to replace plugins initialization placehodler
// in Caddy's run.go.
func (fp *ForwardproxyPlugin) GetPluginImportLine() string {
    return fp.importLine
}

// Get plugin name.
func (fp *ForwardproxyPlugin) GetPluginName() string {
    return fp.pluginName
}

// Get plugin's sources URL, for using with Builder.
func (fp *ForwardproxyPlugin) GetPluginSourcesURL() string {
    return fp.sourcesUrl
}

// Plugin initialization.
func (fp *ForwardproxyPlugin) Initialize() {
    fp.pluginName = "forwardproxy"
    fp.pluginDescription = "Enables Caddy webserver to act as a Secure Web Proxy. Effectively, forwards HTTP requests in following form: GET someplace.else/thing.html HTTP/1.1 and establishes secure TCP tunnels for: CONNECT someotherweb.server HTTP/1.1."
    fp.sourcesUrl = "https://github.com/caddyserver/forwardproxy"
    fp.documentationUrl = "https://caddyserver.com/docs/http.forwardproxy"
    fp.importLine = "github.com/caddyserver/forwardproxy"
}

// Installation
func (fp *ForwardproxyPlugin) Install(workspace_path string) {
    // Do nothing if user don't want to install this plugin.
    if !ctx.Flags.BUILD_WITH_FORWARDPROXY {
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

    ctx.Log.Printf("Installing plugin: %s", fp.GetPluginName())

    err1 := ctx.CmdWorker.Execute(fmt.Sprintf("go get -d -u %s", fp.GetPluginImportLine()))
    if err1 != nil {
        ctx.Log.Fatalf("Failed to get plugin's sources: %s", err1.Error())
    }

    // Replace default "This is where other plugins get plugged in (imported)"
    // line with plugin import.
    replace_to := fmt.Sprintf("_ \"%s\"\n\t// This is where other plugins get plugged in (imported)", fp.GetPluginImportLine())
    fh = strings.Replace(fh, "// This is where other plugins get plugged in (imported)", replace_to, 1)
    // Write file.
    ioutil.WriteFile(rungo, []byte(fh), os.ModePerm)
}
