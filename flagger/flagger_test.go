// Caddybuilder - a friendly tool to build Caddy executable.
// Copyright (c) 2017, Stanislav N. aka pztrn <pztrn at pztrn dot name>
//

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
func prepareToTests() {
    f = &Flagger{}
}

// Main test, which will group all other tests.
func TestFlaggerPreparation(t *testing.T) {
    prepareToTests()
}

func TestParamsParsing(t *testing.T) {
    oldArgs := os.Args
    defer func() { os.Args = oldArgs }()

    os.Args = []string{"caddybuilder", "-cors", "-realip", "-search"}
    f.Initialize()

    if f.BUILD_WITH_CORS && f.BUILD_WITH_REALIP && f.BUILD_WITH_SEARCH {
        // All ok.
    } else {
        if !f.BUILD_WITH_CORS {
            t.Fatal("BUILD_WITH_CORS = false")
        }
        if !f.BUILD_WITH_REALIP {
            t.Fatal("BUILD_WITH_REALIP = false")
        }
        if !f.BUILD_WITH_SEARCH {
            t.Fatal("BUILD_WITH_SEARCH = false")
        }
        t.Fatal("testParamsParsing: Parameters parsing failed!")
        t.FailNow()
    }
}
