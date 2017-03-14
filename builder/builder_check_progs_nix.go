// Caddybuilder - a friendly tool to build Caddy executable.
// Copyright (c) 2017, Stanislav N. aka pztrn <pztrn at pztrn dot name>
//

// build: !windows

package builder

import (
    // stdlib
    "errors"
    "fmt"
    "path/filepath"
    "os"
    "strings"
)


// Checks for neccessary programs.
// It will take PATH, split it and iterate over it to check if
// needed binary exists.
func (b *Builder) checkForPrograms() error {
    path, path_defined := os.LookupEnv("PATH")
    if !path_defined || path == "" {
        return errors.New("PATH environment variable isn't defined! Cannot continue!")
    }

    path_splitted := strings.Split(path, ":")

    for i, _ := range b.NeccessaryPrograms {
        prgname := i

        var found bool = false

        for i := range path_splitted {
            p := path_splitted[i]

            path_to_try := filepath.Join(p, prgname)
            _, err := os.Stat(path_to_try)
            if err != nil {
                continue
            }

            b.NeccessaryPrograms[prgname] = path_to_try
            found = true
            break
        }

        if !found {
            return errors.New(fmt.Sprintf("Program not found in PATH: %s", prgname))
        }

    }

    return nil
}
