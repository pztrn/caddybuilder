// Caddybuilder - a friendly tool to build Caddy executable.
// Copyright (c) 2017, Stanislav N. aka pztrn <pztrn at pztrn dot name>
//

package builder

import (
    // stdlib
    l "log"

    // local
    "github.com/pztrn/caddybuilder/flagger"
)

var (
    flags *flagger.Flagger
    log *l.Logger
)

func New(f *flagger.Flagger, l *l.Logger) *Builder {
    flags = f
    log = l
    return &Builder{}
}
