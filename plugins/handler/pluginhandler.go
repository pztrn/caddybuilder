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
