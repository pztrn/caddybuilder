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

package awslambda

import (
	// stdlib
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type AwsLambdaPlugin struct {
	documentationUrl  string
	importLine        string
	pluginDescription string
	pluginName        string
	sourcesUrl        string
}

// Get plugin description. Short one :).
func (alp *AwsLambdaPlugin) GetPluginDescription() string {
	return alp.pluginDescription
}

// Get plugin documentation URL.
func (alp *AwsLambdaPlugin) GetPluginDocumentationURL() string {
	return alp.documentationUrl
}

// Get plugin's import line.
// This line will be used to replace plugins initialization placehodler
// in Caddy's run.go.
func (alp *AwsLambdaPlugin) GetPluginImportLine() string {
	return alp.importLine
}

// Get plugin name.
func (alp *AwsLambdaPlugin) GetPluginName() string {
	return alp.pluginName
}

// Get plugin's sources URL, for using with Builder.
func (alp *AwsLambdaPlugin) GetPluginSourcesURL() string {
	return alp.sourcesUrl
}

// Plugin initialization.
func (alp *AwsLambdaPlugin) Initialize() {
	alp.pluginName = "awslambda"
	alp.pluginDescription = "awslambda proxies requests to AWS Lambda functions using the AWS Lambda Invoke operation. It provides an alternative to AWS API Gateway and provides a simple way to declaratively proxy requests to a set of Lambda functions without per-function configuration."
	alp.sourcesUrl = "https://github.com/coopernurse/caddy-awslambda"
	alp.documentationUrl = "https://caddyserver.com/docs/http.awslambda"
	alp.importLine = "github.com/coopernurse/caddy-awslambda"
}

// Installation
func (alp *AwsLambdaPlugin) Install(workspace_path string) {
	// Do nothing if user don't want to install this plugin.
	if !ctx.Flags.BUILD_WITH_AWSLAMBDA {
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

	ctx.Log.Printf("Installing plugin: %s", alp.GetPluginName())

	err1 := ctx.CmdWorker.Execute(fmt.Sprintf("go get -d -u %s", alp.GetPluginImportLine()))
	if err1 != nil {
		ctx.Log.Fatalf("Failed to get plugin's sources: %s", err1.Error())
	}

	// Replace default "This is where other plugins get plugged in (imported)"
	// line with plugin import.
	replace_to := fmt.Sprintf("_ \"%s\"\n\t// This is where other plugins get plugged in (imported)", alp.GetPluginImportLine())
	fh = strings.Replace(fh, "// This is where other plugins get plugged in (imported)", replace_to, 1)
	// Write file.
	ioutil.WriteFile(rungo, []byte(fh), os.ModePerm)
}
