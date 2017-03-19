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

package plugins

import (
	// stdlib
	"bytes"
	lt "log"
	//"os"
	"testing"

	// local
	"github.com/pztrn/caddybuilder/cmdworker"
	"github.com/pztrn/caddybuilder/flagger"
)

var (
	c *cmdworker.CmdWorker
	f *flagger.Flagger
	p *Plugins
)

// Preparation for tests.
func prepareToTests() {
	// Initialize flagger.
	f = flagger.New()
	f.DO_NOT_REMOVE_CURRENT_GOPATH = true
	f.BUILD_OUTPUT = "/tmp/caddybuilder-gopath/bin/caddy.test"
	f.BUILD_WITH_CORS = true
	f.BUILD_WITH_REALIP = true

	// Initialize dummy logger.
	buf := bytes.NewBuffer([]byte(""))
	l := lt.New(buf, "", lt.Lmicroseconds|lt.LstdFlags)

	// CmdWorker.
	c = cmdworker.New(l)
	c.Initialize()

	p = New(c, f, l)
	p.Initialize()
}

func TestPluginsPreparation(t *testing.T) {
	prepareToTests()
}

// Next tests are checked against all available plugins, so plugin
// writers must not create own tests for them.

func TestGetPluginDescription(t *testing.T) {
	for i := range p.PluginsList {
		desc := p.PluginsList[i].GetPluginDescription()
		if len(desc) == 0 {
			t.Fatalf("Plugin %s: empty description", i)
			t.FailNow()
		}
	}
	t.Logf("Tested %d plugins", len(p.PluginsList))
}

func TestGetPluginDocumentationURL(t *testing.T) {
	for i := range p.PluginsList {
		desc := p.PluginsList[i].GetPluginDocumentationURL()
		if len(desc) == 0 {
			t.Fatalf("Plugin %s: empty documentation URL", i)
			t.FailNow()
		}
	}
	t.Logf("Tested %d plugins", len(p.PluginsList))
}

func TestGetPluginImportLine(t *testing.T) {
	for i := range p.PluginsList {
		desc := p.PluginsList[i].GetPluginImportLine()
		if len(desc) == 0 {
			t.Fatalf("Plugin %s: empty import line", i)
			t.FailNow()
		}
	}
	t.Logf("Tested %d plugins", len(p.PluginsList))
}

func TestGetPluginName(t *testing.T) {
	for i := range p.PluginsList {
		desc := p.PluginsList[i].GetPluginName()
		if len(desc) == 0 {
			t.Fatalf("Plugin %s: empty plugin name", i)
			t.FailNow()
		}
	}
	t.Logf("Tested %d plugins", len(p.PluginsList))
}

func TestGetPluginSourcesURL(t *testing.T) {
	for i := range p.PluginsList {
		desc := p.PluginsList[i].GetPluginSourcesURL()
		if len(desc) == 0 {
			t.Fatalf("Plugin %s: empty sources URL", i)
			t.FailNow()
		}
	}
	t.Logf("Tested %d plugins", len(p.PluginsList))
}
