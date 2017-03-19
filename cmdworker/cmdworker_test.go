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

package cmdworker

import (
	// stdlib
	"bytes"
	lt "log"
	"testing"

	// local
	"github.com/pztrn/caddybuilder/flagger"
	"github.com/pztrn/caddybuilder/plugins"
)

var (
	c *CmdWorker
	f *flagger.Flagger
	p *plugins.Plugins
)

// Preparation for tests.
func prepareToTests() {
	// Initialize dummy logger.
	buf := bytes.NewBuffer([]byte(""))
	l := lt.New(buf, "", lt.Lmicroseconds|lt.LstdFlags)

	c = New(l)
	c.Initialize()
}

func TestCmdWorkerPreparation(t *testing.T) {
	prepareToTests()
}

// Test command execution.
func TestExecute(t *testing.T) {
	err := c.Execute("echo 'success'")
	if err != nil {
		t.Fatal(err.Error())
		t.FailNow()
	}
}
