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

package minify

import (
	// stdlib
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type MinifyPlugin struct {
	documentationUrl  string
	importLine        string
	pluginDescription string
	pluginName        string
	sourcesUrl        string
}

// Get plugin description. Short one :).
func (mip *MinifyPlugin) GetPluginDescription() string {
	return mip.pluginDescription
}

// Get plugin documentation URL.
func (mip *MinifyPlugin) GetPluginDocumentationURL() string {
	return mip.documentationUrl
}

// Get plugin's import line.
// This line will be used to replace plugins initialization placehodler
// in Caddy's run.go.
func (mip *MinifyPlugin) GetPluginImportLine() string {
	return mip.importLine
}

// Get plugin name.
func (mip *MinifyPlugin) GetPluginName() string {
	return mip.pluginName
}

// Get plugin's sources URL, for using with Builder.
func (mip *MinifyPlugin) GetPluginSourcesURL() string {
	return mip.sourcesUrl
}

// Plugin initialization.
func (mip *MinifyPlugin) Initialize() {
	mip.pluginName = "minify"
	mip.pluginDescription = "minify enables minifying on your website. By default it will minify every css, javascript, html, svg, xml and json files. You can exclude some paths or single files so they're not minified."
	mip.sourcesUrl = "https://github.com/hacdias/caddy-minify"
	mip.documentationUrl = "https://caddyserver.com/docs/minify"
	mip.importLine = "github.com/hacdias/caddy-minify"
}

// Installation
func (mip *MinifyPlugin) Install(workspace_path string) {
	// Do nothing if user don't want to install this plugin.
	if !ctx.Flags.BUILD_WITH_MINIFY {
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

	ctx.Log.Printf("Installing plugin: %s", mip.GetPluginName())

	err1 := ctx.CmdWorker.Execute(fmt.Sprintf("go get -d -u %s", mip.GetPluginImportLine()))
	if err1 != nil {
		ctx.Log.Fatalf("Failed to get plugin's sources: %s", err1.Error())
	}

	// Replace default "This is where other plugins get plugged in (imported)"
	// line with plugin import.
	replace_to := fmt.Sprintf("_ \"%s\"\n\t// This is where other plugins get plugged in (imported)", mip.GetPluginImportLine())
	fh = strings.Replace(fh, "// This is where other plugins get plugged in (imported)", replace_to, 1)
	// Write file.
	ioutil.WriteFile(rungo, []byte(fh), os.ModePerm)
}
