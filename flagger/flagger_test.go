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

package flagger

import (
	// stdlib
	"os"
	"testing"
)

var (
	f *Flagger
)

// Preparation for tests.
func TestFlaggerPreparation(t *testing.T) {
	f = &Flagger{}
}

func TestFlaggerParamsParsing(t *testing.T) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	os.Args = []string{"caddybuilder", "-cors", "-realip", "-webdav"}
	f.Initialize()

	if f.BUILD_WITH_CORS && f.BUILD_WITH_REALIP && f.BUILD_WITH_WEBDAV {
		// All ok.
	} else {
		if !f.BUILD_WITH_CORS {
			t.Fatal("BUILD_WITH_CORS = false")
		}
		if !f.BUILD_WITH_REALIP {
			t.Fatal("BUILD_WITH_REALIP = false")
		}
		if !f.BUILD_WITH_WEBDAV {
			t.Fatal("BUILD_WITH_WEBDAV = false")
		}
		t.Fatal("testParamsParsing: Parameters parsing failed!")
		t.FailNow()
	}
}
