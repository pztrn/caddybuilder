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

package main

import (
	// stdlib
	l "log"
	"os"

	// local
	"github.com/pztrn/caddybuilder/builder"
	"github.com/pztrn/caddybuilder/cmdworker"
	"github.com/pztrn/caddybuilder/common"
	"github.com/pztrn/caddybuilder/flagger"
	"github.com/pztrn/caddybuilder/plugins"
)

var (
	b     *builder.Builder
	cw    *cmdworker.CmdWorker
	flags *flagger.Flagger
	log   *l.Logger
	pl    *plugins.Plugins
)

func main() {
	// Initializing logger.
	log = l.New(os.Stdout, "", l.Lmicroseconds|l.LstdFlags)

	flags = flagger.New()
	flags.Initialize()

	log.Print("This is Caddybuilder, version " + common.CADDYBUILDER_VERSION)

	cw = cmdworker.New(log)
	cw.Initialize()

	pl = plugins.New(cw, flags, log)
	pl.Initialize()

	b = builder.New(cw, flags, log, pl)
	b.Initialize()
	b.Proceed()
}
