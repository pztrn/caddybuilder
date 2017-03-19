// Caddybuilder - a friendly tool to build Caddy executable.
// Copyright (c) 2017, Stanislav N. aka pztrn <pztrn at pztrn dot name>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file exchpt in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package hugo

import (
    // stdlib
    "fmt"
    "io/ioutil"
    "os"
    "path/filepath"
    "strings"
)

type HugoPlugin struct {
    documentationUrl  string
    importLine        string
    pluginDescription string
    pluginName        string
    sourcesUrl        string
}

// Get plugin description. Short one :).
func (hp *HugoPlugin) GetPluginDescription() string {
    return hp.pluginDescription
}

// Get plugin documentation URL.
func (hp *HugoPlugin) GetPluginDocumentationURL() string {
    return hp.documentationUrl
}

// Get plugin's import line.
// This line will be used to rhplace plugins initialization placehodler
// in Caddy's run.go.
func (hp *HugoPlugin) GetPluginImportLine() string {
    return hp.importLine
}

// Get plugin name.
func (hp *HugoPlugin) GetPluginName() string {
    return hp.pluginName
}

// Get plugin's sources URL, for using with Builder.
func (hp *HugoPlugin) GetPluginSourcesURL() string {
    return hp.sourcesUrl
}

// Plugin initialization.
func (hp *HugoPlugin) Initialize() {
    hp.pluginName = "hugo"
    hp.pluginDescription = "hugo fills the gap between Hugo and the browser."
    hp.sourcesUrl = "https://github.com/hacdias/caddy-hugo"
    hp.documentationUrl = "https://caddyserver.com/docs/hugo"
    hp.importLine = "github.com/hacdias/caddy-hugo"
}

// Installation
func (hp *HugoPlugin) Install(workspace_path string) {
    // Do nothing if user don't want to install this plugin.
    if !ctx.Flags.BUILD_WITH_HUGO {
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

    ctx.Log.Printf("Installing plugin: %s", hp.GetPluginName())

    err1 := ctx.CmdWorker.Execute(fmt.Sprintf("go get -d -u %s", hp.GetPluginImportLine()))
    if err1 != nil {
        ctx.Log.Fatalf("Failed to get plugin's sources: %s", err1.Error())
    }

    // We have to run go generate.
    os.Chdir(fmt.Sprintf("%s/src/%s", workspace_path, hp.GetPluginImportLine()))
    err2 := ctx.CmdWorker.Execute("go generate")
    if err2 != nil {
        ctx.Log.Fatalf("Failed to run go generate: %s", err2.Error())
    }

    // Replace default "This is where other plugins get plugged in (imported)"
    // line with plugin import.
    replace_to := fmt.Sprintf("_ \"%s\"\n\t// This is where other plugins get plugged in (imported)", hp.GetPluginImportLine())
    fh = strings.Replace(fh, "// This is where other plugins get plugged in (imported)", replace_to, 1)
    // Write file.
    ioutil.WriteFile(rungo, []byte(fh), os.ModePerm)
}
