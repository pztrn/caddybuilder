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

// build: !windows

package builder

import (
	// stdlib
	"errors"
	"fmt"
	"os"
	"path/filepath"
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
