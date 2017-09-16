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

package webdav

import (
    // stdlib
    "fmt"
    "io/ioutil"
    "os"
    "path/filepath"
    "strings"
)

type WebdavPlugin struct {
    documentationUrl  string
    importLine        string
    pluginDescription string
    pluginName        string
    sourcesUrl        string
}

// Get plugin description. Short one :).
func (wp *WebdavPlugin) GetPluginDescription() string {
    return wp.pluginDescription
}

// Get plugin documentation URL.
func (wp *WebdavPlugin) GetPluginDocumentationURL() string {
    return wp.documentationUrl
}

// Get plugin's import line.
// This line will be used to replace plugins initialization placehodler
// in Caddy's run.go.
func (wp *WebdavPlugin) GetPluginImportLine() string {
    return wp.importLine
}

// Get plugin name.
func (wp *WebdavPlugin) GetPluginName() string {
    return wp.pluginName
}

// Get plugin's sources URL, for using with Builder.
func (wp *WebdavPlugin) GetPluginSourcesURL() string {
    return wp.sourcesUrl
}

// Plugin initialization.
func (wp *WebdavPlugin) Initialize() {
    wp.pluginName = "webdav"
    wp.pluginDescription = "WebDAV capabilities with support for path restriction rules and users."
    wp.sourcesUrl = "https://github.com/hacdias/caddy-webdav"
    wp.documentationUrl = "https://caddyserver.com/docs/http.webdav"
    wp.importLine = "github.com/hacdias/caddy-webdav"
}

// Installation
func (wp *WebdavPlugin) Install(workspace_path string) {
    // Do nothing if user don't want to install this plugin.
    if !ctx.Flags.BUILD_WITH_WEBDAV {
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

    ctx.Log.Printf("Installing plugin: %s", wp.GetPluginName())

    err1 := ctx.CmdWorker.Execute(fmt.Sprintf("go get -d -u %s", wp.GetPluginImportLine()))
    if err1 != nil {
        ctx.Log.Fatalf("Failed to get plugin's sources: %s", err1.Error())
    }

    // Replace default "This is where other plugins get plugged in (imported)"
    // line with plugin import.
    replace_to := fmt.Sprintf("_ \"%s\"\n\t// This is where other plugins get plugged in (imported)", wp.GetPluginImportLine())
    fh = strings.Replace(fh, "// This is where other plugins get plugged in (imported)", replace_to, 1)
    // Write file.
    ioutil.WriteFile(rungo, []byte(fh), os.ModePerm)
}
