// Caddybuilder - a friendly tool to build Caddy executable.
// Copyright (c) 2017-2018, Stanislav N. aka pztrn <pztrn at pztrn dot name>
//
// Permission is hereby granted, free of charge, to any person obtaining
// a copy of this software and associated documentation files (the
// "Software"), to deal in the Software without restriction, including
// without limitation the rights to use, copy, modify, merge, publish,
// distribute, sublicense, and/or sell copies of the Software, and to
// permit persons to whom the Software is furnished to do so, subject
// to the following conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
// IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY
// CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT,
// TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE
// OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

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
	fp.sourcesUrl = "https://github.com/hacdias/filemanager"
	fp.documentationUrl = "https://caddyserver.com/docs/http.filemanager"
	fp.importLine = "github.com/hacdias/filemanager"
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

	// For some reason one of dependencies didn't installed with previous
	// command. Okay, will install it manually.
	// Some explanations:
	//   * "github.com/gorilla/websocket" - Caddy vendored it, so it's not
	//     in GOPATH.
	//   * "github.com/hacdias/varutils" - needed for Hugo plugin.
	deps_to_install := []string{
		"github.com/asdine/storm",
		"github.com/dgrijalva/jwt-go",
		"github.com/gorilla/websocket",
		"github.com/hacdias/varutils",
		"github.com/mholt/archiver",
		"github.com/mitchellh/mapstructure",
		"github.com/filebrowser/filebrowser",
	}
	for i := range deps_to_install {
		err2 := ctx.CmdWorker.Execute(fmt.Sprintf("go get -d -u %s", deps_to_install[i]))
		if err2 != nil {
			ctx.Log.Fatalf("Failed to get dependency for filemanager: %s", deps_to_install[i])
		}
	}

	// We have to run go generate.
	os.Chdir(fmt.Sprintf("%s/src/%s", workspace_path, fp.GetPluginImportLine()))
	err3 := ctx.CmdWorker.Execute("go generate")
	if err3 != nil {
		ctx.Log.Fatalf("Failed to run go generate: %s", err3.Error())
	}

	// Three plugins to install - filemanager, hugo, jekyll.
	plugins_to_install := []string{"filemanager", "hugo", "jekyll"}
	for i := range plugins_to_install {
		// Replace default "This is where other plugins get plugged in (imported)"
		// line with plugin import.
		replace_to := fmt.Sprintf("_ \"%s\"\n\t// This is where other plugins get plugged in (imported)", fp.GetPluginImportLine()+"/caddy/"+plugins_to_install[i])
		fh = strings.Replace(fh, "// This is where other plugins get plugged in (imported)", replace_to, 1)
	}

	// Write file.
	ioutil.WriteFile(rungo, []byte(fh), os.ModePerm)
}
