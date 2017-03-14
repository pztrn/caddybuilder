// Caddybuilder - a friendly tool to build Caddy executable.
// Copyright (c) 2017, Stanislav N. aka pztrn <pztrn at pztrn dot name>
//

package builder

import (
    // stdlib
    "bytes"
    lt "log"
    "os"
    "testing"

    // local
    "github.com/pztrn/caddybuilder/flagger"
)

var (
    b *Builder
    f *flagger.Flagger
)

// Preparation for tests.
func prepareToTests() {
    // Initialize flagger.
    f = flagger.New()
    f.DO_NOT_REMOVE_CURRENT_GOPATH = true
    f.BUILD_OUTPUT = "/tmp/caddybuilder-root/bin/"
    f.BUILD_WITH_CORS = true
    f.BUILD_WITH_REALIP = true

    // Initialize dummy logger.
    buf := bytes.NewBuffer([]byte(""))
    l := lt.New(buf, "", lt.Lmicroseconds | lt.LstdFlags)

    b = New(f, l)
    b.Initialize()
}

// Main test, which will group all other tests.
func TestBuilderPreparation(t *testing.T) {
    prepareToTests()
}

// Test checking for programs existing WITHOUT defined PATH.
func TestCheckForProgramsWithoutPath(t *testing.T) {
    prepareToTests()
    oldPath := os.Getenv("PATH")
    defer func() { os.Setenv("PATH", oldPath) }()
    os.Setenv("PATH", "")

    err := b.checkForPrograms()
    if err == nil {
        t.Fatal("testCheckForProgramsWithoutPath: Found needed programs without defined PATH!")
        t.FailNow()
    }
}

// Test checking for programs existing WITH defined PATH.
func TestCheckForPrograms(t *testing.T) {
    prepareToTests()
    err := b.checkForPrograms()
    if err != nil {
        t.Fatal("testCheckForPrograms:", err.Error())
        t.FailNow()
    }

    for binary, path := range b.NeccessaryPrograms {
        if path == "" {
            t.Fatal("testCheckForPrograms:", binary, "have no path!")
            t.FailNow()
        }
    }
}

// Test go get execution.
func TestGoGet(t *testing.T) {
    prepareToTests()
    b.checkEnvironmentVariables()
    _ = b.checkForPrograms()
    b.prepareEnvironment()
    // We will test with Caddy sources.
    err := b.goGet("github.com/mholt/caddy")
    if err != nil {
        t.Log("Error occured while getting Caddy sources:")
        t.Fatal(err.Error())
        t.FailNow()
    }
}

// Test Caddy building.
func TestBuild(t *testing.T) {
    t.Skip("Temporary skipped")
}
