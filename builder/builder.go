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

// File "builder.go" contains Builder struct itself, along with main
// builder login.
package builder

type Builder struct {
	// Neccessary programs list.
	NeccessaryPrograms map[string]string
	// Old GOPATH.
	OldGOPATH string
	// Go's workspace path.
	Workspace string
}

// Initializes Builder.
func (b *Builder) Initialize() {
	// Neccessary programs list.
	b.NeccessaryPrograms = map[string]string{
		"go":  "",
		"git": "",
	}

	// Set some defaults. They can be overriden with
	// b.checkEnvironmentVariables().

	// Workspace path.
	b.Workspace = "/tmp/caddybuilder-gopath"
}

// Prints current Builder configuration.
func (b *Builder) printConfiguration() {
	log.Printf("Go's workspace path: %s", b.Workspace)
	log.Printf("Caddy binary will be installed at: %s", flags.BUILD_OUTPUT)

	log.Print("Binaries found:")

	for prgname, path := range b.NeccessaryPrograms {
		log.Printf("\t%s: %s", prgname, path)
	}
}

// Proceeds with Caddy building.
// It will do git clone of Caddy and will look into Flagger for
// plugins we should install.
func (b *Builder) Proceed() {
	// Check environment variables.
	b.checkEnvironmentVariables()
	// Check for neccessary programs.
	log.Print("Checking for neccessary programs...")
	err := b.checkForPrograms()
	if err != nil {
		log.Fatal("Failed to find binary: ", err.Error())
	}
	// Print them out.
	b.printConfiguration()
	// Prepare environment.
	b.prepareEnvironment()

	// From this point we assume that we have prepared GOPATH
	// and found paths to all neccessary binaries.
	// We begin with go get'ing Caddy.
	log.Print("Starting installing Caddy and plugins...")
	err1 := b.goGet("github.com/mholt/caddy/caddy")
	if err1 != nil {
		log.Print("Error occured while getting Caddy sources:")
		log.Fatal(err1.Error())
	}

	// Install all requested plugins.
	// For now it is controlled by checking flagger's value, but this
	// will change.
	b.installPlugin("awslambda")
	b.installPlugin("cors")
	b.installPlugin("expires")
	b.installPlugin("filemanager")
	b.installPlugin("git")
	b.installPlugin("hugo")
	b.installPlugin("ipfilter")
	b.installPlugin("jsonp")
	b.installPlugin("jwt")
	b.installPlugin("locale")
	b.installPlugin("mailout")
	b.installPlugin("minify")
	b.installPlugin("multipass")
	b.installPlugin("prometheus")
	b.installPlugin("ratelimit")
	b.installPlugin("realip")
	b.installPlugin("search")
	b.installPlugin("upload")

	// We're set! Build Caddy!
	log.Print("Building Caddy...")
	err2 := b.buildCaddy()
	if err2 != nil {
		log.Fatal(err2.Error())
	}
}
