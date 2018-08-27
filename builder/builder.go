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

// File "builder.go" contains Builder struct itself, along with main
// builder login.
package builder

import (
	// stdlib
	"fmt"
	"os"
)

type Builder struct {
	// Neccessary programs list.
	NeccessaryPrograms map[string]string
	// Old GOPATH.
	OldGOPATH string
	// Old PATH.
	OldPATH string
	// Go's workspace path.
	Workspace string
}

// Actually build Caddy binary.
func (b *Builder) buildCaddy() error {
	command := fmt.Sprintf("%s build -a -o %s github.com/mholt/caddy/caddy", b.NeccessaryPrograms["go"], flags.BUILD_OUTPUT)
	err := cw.Execute(command)
	if err != nil {
		return err
	}

	return nil
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
	log.Printf("System's path: %s", os.Getenv("PATH"))
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
	// Prepare environment.
	b.prepareEnvironment()
	// Print configuration.
	b.printConfiguration()

	// From this point we assume that we have prepared GOPATH
	// and found paths to all neccessary binaries.
	// We begin with go get'ing Caddy.
	log.Print("Starting installing Caddy and plugins...")
	log.Print("Obtaining Caddy sources...")
	err1 := cw.Execute(fmt.Sprintf("%s get -d -u github.com/mholt/caddy/caddy", b.NeccessaryPrograms["go"]))
	if err1 != nil {
		log.Print("Error occured while getting Caddy sources:")
		log.Fatal(err1.Error())
	}

	// Install all requested plugins.
	// For now it is controlled by checking flagger's value, but this
	// will change.
	for i := range pl.PluginsList {
		handler := pl.PluginsList[i]
		handler.Install(b.Workspace)
	}

	// We're set! Build Caddy!
	log.Print("Building Caddy...")
	err2 := b.buildCaddy()
	if err2 != nil {
		log.Fatal(err2.Error())
	}
}
