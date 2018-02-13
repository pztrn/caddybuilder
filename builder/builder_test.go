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

package builder

import (
	// stdlib
	"bytes"
	"io/ioutil"
	lt "log"
	"os"
	"os/exec"
	"strings"
	"testing"

	// local
	"github.com/pztrn/caddybuilder/cmdworker"
	"github.com/pztrn/caddybuilder/flagger"
	"github.com/pztrn/caddybuilder/plugins"
)

var (
	b *Builder
	c *cmdworker.CmdWorker
	f *flagger.Flagger
	p *plugins.Plugins
)

// Preparation for tests.
func prepareToTests() {
	// Initialize flagger.
	f = flagger.New()
	f.DO_NOT_REMOVE_CURRENT_GOPATH = true
	f.BUILD_OUTPUT = "/tmp/caddybuilder-gopath/bin/caddy.test"
	f.BUILD_WITH_REALIP = true
	f.BUILD_WITH_UPLOAD = true

	// Initialize dummy logger.
	buf := bytes.NewBuffer([]byte(""))
	l := lt.New(buf, "", lt.Lmicroseconds|lt.LstdFlags)

	// CmdWorker.
	c = cmdworker.New(l)
	c.Initialize()

	// Initialize plugins.
	p = plugins.New(c, f, l)
	p.Initialize()

	b = New(c, f, l, p)
	b.Initialize()
}

// Main test, which will group all other tests.
func TestInitialization(t *testing.T) {
	prepareToTests()
}

// Test checking for programs existing WITHOUT defined PATH.
func TestCheckForProgramsWithoutPath(t *testing.T) {
	oldPath := os.Getenv("PATH")
	defer func() { os.Setenv("PATH", oldPath) }()
	os.Setenv("PATH", "")

	err := b.checkForPrograms()
	if err == nil {
		t.Fatal("testCheckForProgramsWithoutPath: Found needed programs without defined PATH!")
		t.FailNow()
	}
}

// Test checking for programs existing WITH defined PATH.
func TestCheckForPrograms(t *testing.T) {
	err := b.checkForPrograms()
	if err != nil {
		t.Fatal("testCheckForPrograms:", err.Error())
		t.FailNow()
	}

	for binary, path := range b.NeccessaryPrograms {
		if path == "" {
			t.Fatal("testCheckForPrograms:", binary, "have no path!")
			t.FailNow()
		}
	}
}

// Test Caddy building.
// We will call Builder.Proceed and check resulted binary for installed
// plugins.
func TestBuilderBuild(t *testing.T) {
	if !testing.Short() {
		t.Log("Building Caddy...")
		b.Proceed()
		// Get installed plugins list.
		cmd := exec.Command("/tmp/caddybuilder-gopath/bin/caddy.test", "-plugins")
		stdout, _ := cmd.StdoutPipe()
		stderr, _ := cmd.StderrPipe()
		// Go, go, go!
		err := cmd.Start()
		if err != nil {
			t.Fatalf("Failed to check installed plugins: %s", err.Error())
		}

		stdout_output, err1 := ioutil.ReadAll(stdout)
		stderr_output, err2 := ioutil.ReadAll(stderr)

		if err1 != nil {
			t.Fatalf("Failed to obtain stdout output!")
			t.FailNow()
		}

		if err2 != nil {
			t.Fatalf("Failed to obtain stderr output!")
			t.FailNow()
		}

		// Wait until command finishes.
		err3 := cmd.Wait()
		if err3 != nil {
			// This means that some error occured in run time.
			t.Log("\tStdout:")
			t.Log(string(stdout_output))
			t.Log("\tStderr:")
			t.Log(string(stderr_output))
			t.Fatalf("Error occured while getting list of installed plugins: %s", err1.Error())
			t.FailNow()
		}

		stdout_as_string := string(stdout_output)
		if !strings.Contains(stdout_as_string, "http.upload") && !strings.Contains(stdout_as_string, "http.realip") {
			t.Log("Required plugins wasn't installed! We need upload and realip, got:")
			t.Fatal(stdout_as_string)
			t.FailNow()
		}
	} else {
		t.Skip("Test skipped due to -short parameter")
	}
}
