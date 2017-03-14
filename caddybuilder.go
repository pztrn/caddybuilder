// Caddybuilder - a friendly tool to build Caddy executable.
// Copyright (c) 2017, Stanislav N. aka pztrn <pztrn at pztrn dot name>
//

package main

import (
    // stdlib
    l "log"
    "os"

    // local
    "github.com/pztrn/caddybuilder/builder"
    "github.com/pztrn/caddybuilder/flagger"
)

var (
    b *builder.Builder
    flags *flagger.Flagger
    log *l.Logger
)

func main() {
    // Initializing logger.
    log = l.New(os.Stdout, "", l.Lmicroseconds | l.LstdFlags)

    flags = flagger.New()
    flags.Initialize()

    b = builder.New(flags, log)
    b.Initialize()
    b.Proceed()
}
