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

// File "environment.go" contains functions which working with environment.
package builder

import (
	// stdlib
	"errors"
	"os"
	"strings"
)

// Checks environment variables we support.
// Values from these variables will overwrite values in Builder structure.
func (b *Builder) checkEnvironmentVariables() {
	log.Print("Checking environment variables...")

	// Go's workspace.
	workspace_path, workspace_path_defined := os.LookupEnv("CB_WORKSPACE")
	if workspace_path_defined {
		b.Workspace = workspace_path
	}
}

// Prepares environment:
//   * Creates GOPATH directory.
//   * Sets new GOPATH.
func (b *Builder) prepareEnvironment() error {
	log.Print("Preparing environment...")

	// Checking for GOPATH directory existing.
	// It should be deleted if so, so every time we're starting with
	// fresh GOPATH.
	// ToDo: flag to re-use sources, maybe default behaviour?
	_, err := os.Stat(b.Workspace)
	if err == nil {
		if !flags.DO_NOT_REMOVE_CURRENT_GOPATH {
			log.Print("Old GOPATH found, removing...")
			os.RemoveAll(b.Workspace)
		}
	}

	os.MkdirAll(b.Workspace, os.ModePerm)

	// Set GOPATH. Old GOPATH resides in b.OldGOPATH.
	current_gopath, gopath_defined := os.LookupEnv("GOPATH")
	if !gopath_defined {
		return errors.New("No GOPATH defined! Something weird going on!")
	}
	b.OldGOPATH = current_gopath
	os.Setenv("GOPATH", b.Workspace)

	// Set PATH. We should include NEW_GOPATH/bin in it. Old path
	// resides in b.OldPath.
	// ToDo: make Windows-compatible?
	b.OldPATH = os.Getenv("PATH")
	os.Setenv("PATH", b.OldPATH + ":" + b.Workspace + "/bin")

	// Output directory. Recreate if exist.
	_, err1 := os.Stat(flags.BUILD_OUTPUT)
	if err1 == nil {
		// We should not even try to remove /usr/local things.
		if strings.Contains(flags.BUILD_OUTPUT, "/usr/") {
			log.Print("Output set somewhere to /usr/, will not delete destination.")
		} else {
			os.RemoveAll(flags.BUILD_OUTPUT)
			os.MkdirAll(flags.BUILD_OUTPUT, os.ModePerm)
		}
	}

	return nil
}
