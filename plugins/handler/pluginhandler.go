// Caddybuilder - a friendly tool to build Caddy executable.
// Copyright (c) 2017-2018 Stanislav N. aka pztrn <pztrn at pztrn dot name>
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

// This package contains interface for plugins, which is used by Builder.
package pluginhandler

type PluginHandler interface {
	// Get plugin description. Short one :).
	GetPluginDescription() string
	// Get plugin documentation URL.
	GetPluginDocumentationURL() string
	// Get plugin's import line.
	// This line will be used to replace plugins initialization placehodler
	// in Caddy's run.go.
	GetPluginImportLine() string
	// Get plugin name.
	GetPluginName() string
	// Get plugin's sources URL, for using with Builder.
	GetPluginSourcesURL() string
	// Plugin initialization.
	Initialize()
	// Installation procedure.
	Install(workspace_path string)
}
