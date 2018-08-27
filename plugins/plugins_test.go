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

package plugins

import (
	// stdlib
	"bytes"
	lt "log"
	//"os"
	"testing"

	// local
	"lab.pztrn.name/pztrn/caddybuilder/cmdworker"
	"lab.pztrn.name/pztrn/caddybuilder/flagger"
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
