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

// This package responsible for parsing CLI parameters (also known as
// flags).
package flagger

import (
	// stdlib
	"flag"
	"fmt"
	"os"
)

var (
	// Help string which will be shown on "-h".
	HELP_STRING = `Caddybuilder - a tool for building Caddy webserver binary with needed
plugins.

Copyright (c) 2017, Stanislav N. aka pztrn <pztrn at pztrn dot name>

Syntax:

  caddybuilder [-flag1] [-flag2] [-flag3 value] ...

Available parameters:
`

	// Help postfix.
	// First line is empty to delimit text from flag package paramteters.
	HELP_POSTFIX = `
Caddybuilder can be configured with environment variables:

  CB_WORKSPACE=/path
        Go's workspace to use. Caddy and plugins sources will be installed
        there.

WARNING: Caddybuilder using git command (not bindings to libgit2 or something
else), therefore bugs will appear!

Exitcodes:
  1 - No parameters was specified and this help message was shown.
`
)

type Flagger struct {
	// Build flags - with what plugins Caddybuilder will build Caddy.

	BUILD_WITH_AWSLAMBDA bool
	BUILD_WITH_CORS bool
	BUILD_WITH_EXPIRES bool
	BUILD_WITH_FILEMANAGER bool
	BUILD_WITH_FILTER bool
	BUILD_WITH_GIT bool
	BUILD_WITH_HUGO bool
	BUILD_WITH_IPFILTER bool
	BUILD_WITH_JSONP bool
	BUILD_WITH_JWT bool
	BUILD_WITH_LOCALE bool
	BUILD_WITH_MAILOUT bool
	BUILD_WITH_MINIFY bool
	BUILD_WITH_MULTIPASS bool
	BUILD_WITH_PROMETHEUS bool
	BUILD_WITH_RATELIMIT bool
	BUILD_WITH_REALIP bool
	BUILD_WITH_SEARCH bool
	BUILD_WITH_UPLOAD bool

	// Output - where binary will be placed.
	BUILD_OUTPUT string

	// Do not remove already composed GOPATH?
	DO_NOT_REMOVE_CURRENT_GOPATH bool
}

// This function initializes Go's flag package.
func (f *Flagger) Initialize() {
	if len(os.Args) == 1 || os.Args[1] == "-h" {
		fmt.Println(HELP_STRING)
	}

	flag.BoolVar(&f.BUILD_WITH_AWSLAMBDA, "awslambda", false, "Build Caddy with AWSLAMBDA plugin")
	flag.BoolVar(&f.BUILD_WITH_CORS, "cors", false, "Build Caddy with CORS plugin")
	flag.BoolVar(&f.BUILD_WITH_EXPIRES, "expires", false, "Build Caddy with EXPIRES plugin")
	flag.BoolVar(&f.BUILD_WITH_FILEMANAGER, "filemanager", false, "Build Caddy with FILEMANAGER plugin")
	flag.BoolVar(&f.BUILD_WITH_FILTER, "filter", false, "Build Caddy with FILTER plugin")
	flag.BoolVar(&f.BUILD_WITH_GIT, "git", false, "Build Caddy with GIT plugin")
	flag.BoolVar(&f.BUILD_WITH_HUGO, "hugo", false, "Build Caddy with HUGO plugin")
	flag.BoolVar(&f.BUILD_WITH_IPFILTER, "ipfilter", false, "Build Caddy with IPFILTER plugin")
	flag.BoolVar(&f.BUILD_WITH_JSONP, "jsonp", false, "Build Caddy with JSONP plugin")
	flag.BoolVar(&f.BUILD_WITH_JWT, "jwt", false, "Build Caddy with JWT plugin")
	flag.BoolVar(&f.BUILD_WITH_LOCALE, "locale", false, "Build Caddy with LOCALE plugin")
	flag.BoolVar(&f.BUILD_WITH_MAILOUT, "mailout", false, "Build Caddy with MAILOUT plugin")
	flag.BoolVar(&f.BUILD_WITH_MINIFY, "minify", false, "Build Caddy with MINIFY plugin")
	flag.BoolVar(&f.BUILD_WITH_MULTIPASS, "multipass", false, "Build Caddy with MULTIPASS plugin")
	flag.BoolVar(&f.BUILD_WITH_PROMETHEUS, "prometheus", false, "Build Caddy with PROMETHEUS plugin")
	flag.BoolVar(&f.BUILD_WITH_RATELIMIT, "ratelimit", false, "Build Caddy with RATELIMIT plugin")
	flag.BoolVar(&f.BUILD_WITH_REALIP, "realip", false, "Build Caddy with REALIP plugin")
	flag.BoolVar(&f.BUILD_WITH_SEARCH, "search", false, "Build Caddy with SEARCH plugin")
	flag.BoolVar(&f.BUILD_WITH_UPLOAD, "upload", false, "Build Caddy with UPLOAD plugin")

	flag.BoolVar(&f.DO_NOT_REMOVE_CURRENT_GOPATH, "donotremovegopath", false, "Do not remove previously created GOPATH (improves speed)")

	// Build output path.
	flag.StringVar(&f.BUILD_OUTPUT, "output", "/usr/local/bin/", "Directory where resulting binary will be placed. It will be recreated if already exist.")

	if len(os.Args) == 1 || os.Args[1] == "-h" {
		flag.PrintDefaults()
		fmt.Println(HELP_POSTFIX)
		os.Exit(1)
	} else {
		flag.Parse()
	}
}
