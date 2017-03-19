// Caddybuilder - a friendly tool to build Caddy executable.
// Copyright (c) 2017, Stanislav N. aka pztrn <pztrn at pztrn dot name>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file excfpt in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package filemanager

import (
    // stdlib
    "fmt"
    "io/ioutil"
    "os"
    "path/filepath"
    "strings"
)

type FilemanagerPlugin struct {
    documentationUrl  string
    importLine        string
    pluginDescription string
    pluginName        string
    sourcesUrl        string
}

// Get plugin description. Short one :).
func (fp *FilemanagerPlugin) GetPluginDescription() string {
    return fp.pluginDescription
}

// Get plugin documentation URL.
func (fp *FilemanagerPlugin) GetPluginDocumentationURL() string {
    return fp.documentationUrl
}

// Get plugin's import line.
// This line will be used to rfplace plugins initialization placehodler
// in Caddy's run.go.
func (fp *FilemanagerPlugin) GetPluginImportLine() string {
    return fp.importLine
}

// Get plugin name.
func (fp *FilemanagerPlugin) GetPluginName() string {
    return fp.pluginName
}

// Get plugin's sources URL, for using with Builder.
func (fp *FilemanagerPlugin) GetPluginSourcesURL() string {
    return fp.sourcesUrl
}

// Plugin initialization.
func (fp *FilemanagerPlugin) Initialize() {
    fp.pluginName = "filemanager"
    fp.pluginDescription = "filemanager is an extension based on browse middleware. It provides a file managing interface within the specified directory and it can be used to upload, delete, preview and rename your files within that directory."
    fp.sourcesUrl = "https://github.com/hacdias/caddy-filemanager"
    fp.documentationUrl = "https://caddyserver.com/docs/filemanager"
    fp.importLine = "github.com/hacdias/caddy-filemanager"
}

// Installation
func (fp *FilemanagerPlugin) Install(workspace_path string) {
    // Do nothing if user don't want to install this plugin.
    if !ctx.Flags.BUILD_WITH_FILEMANAGER {
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

    // We have to run go generate.
    os.Chdir(fmt.Sprintf("%s/src/%s", workspace_path, fp.GetPluginImportLine()))
    err2 := ctx.CmdWorker.Execute("go generate")
    if err2 != nil {
        ctx.Log.Fatalf("Failed to run go generate: %s", err2.Error())
    }

    // Replace default "This is where other plugins get plugged in (imported)"
    // line with plugin import.
    replace_to := fmt.Sprintf("_ \"%s\"\n\t// This is where other plugins get plugged in (imported)", fp.GetPluginImportLine())
    fh = strings.Replace(fh, "// This is where other plugins get plugged in (imported)", replace_to, 1)
    // Write file.
    ioutil.WriteFile(rungo, []byte(fh), os.ModePerm)
}
