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

package restic

import (
	// stdlib
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type ResticPlugin struct {
	documentationUrl  string
	importLine        string
	pluginDescription string
	pluginName        string
	sourcesUrl        string
}

// Get plugin description. Short one :).
func (rp *ResticPlugin) GetPluginDescription() string {
	return rp.pluginDescription
}

// Get plugin documentation URL.
func (rp *ResticPlugin) GetPluginDocumentationURL() string {
	return rp.documentationUrl
}

// Get plugin's import line.
// This line will be used to replace plugins initialization placehodler
// in Caddy's run.go.
func (rp *ResticPlugin) GetPluginImportLine() string {
	return rp.importLine
}

// Get plugin name.
func (rp *ResticPlugin) GetPluginName() string {
	return rp.pluginName
}

// Get plugin's sources URL, for using with Builder.
func (rp *ResticPlugin) GetPluginSourcesURL() string {
	return rp.sourcesUrl
}

// Plugin initialization.
func (rp *ResticPlugin) Initialize() {
	rp.pluginName = "restic"
	rp.pluginDescription = "Serves a restic repository at a path. Restic is a secure and fast backup program, see https://restic.github.io for details."
	rp.sourcesUrl = "https://github.com/restic/caddy"
	rp.documentationUrl = "https://caddyserver.com/docs/http.restic"
	rp.importLine = "github.com/restic/caddy"
}

// Installation
func (rp *ResticPlugin) Install(workspace_path string) {
	// Do nothing if user don't want to install this plugin.
	if !ctx.Flags.BUILD_WITH_RESTIC {
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

	ctx.Log.Printf("Installing plugin: %s", rp.GetPluginName())

	err1 := ctx.CmdWorker.Execute(fmt.Sprintf("go get -d -u %s", rp.GetPluginImportLine()))
	if err1 != nil {
		ctx.Log.Fatalf("Failed to get plugin's sources: %s", err1.Error())
	}

	// Replace default "This is where other plugins get plugged in (imported)"
	// line with plugin import.
	replace_to := fmt.Sprintf("_ \"%s\"\n\t// This is where other plugins get plugged in (imported)", rp.GetPluginImportLine())
	fh = strings.Replace(fh, "// This is where other plugins get plugged in (imported)", replace_to, 1)
	// Write file.
	ioutil.WriteFile(rungo, []byte(fh), os.ModePerm)
}
