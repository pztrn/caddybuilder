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
	os.Setenv("PATH", b.OldPATH+":"+b.Workspace+"/bin")

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
