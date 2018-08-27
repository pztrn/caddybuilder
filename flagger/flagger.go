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

Copyright (c) 2017-2018, Stanislav N. aka pztrn <pztrn at pztrn dot name>

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

	BUILD_WITH_AUTHZ         bool
	BUILD_WITH_AWSES         bool
	BUILD_WITH_AWSLAMBDA     bool
	BUILD_WITH_CACHE         bool
	BUILD_WITH_CGI           bool
	BUILD_WITH_CORS          bool
	BUILD_WITH_DATADOG       bool
	BUILD_WITH_EXPIRES       bool
	BUILD_WITH_FILEMANAGER   bool
	BUILD_WITH_FILTER        bool
	BUILD_WITH_FORWARDPROXY  bool
	BUILD_WITH_GIT           bool
	BUILD_WITH_GOPKG         bool
	BUILD_WITH_GRPC          bool
	BUILD_WITH_HUGO          bool
	BUILD_WITH_IPFILTER      bool
	BUILD_WITH_JSONP         bool
	BUILD_WITH_JWT           bool
	BUILD_WITH_LOCALE        bool
	BUILD_WITH_MAILOUT       bool
	BUILD_WITH_MINIFY        bool
	BUILD_WITH_NOBOTS        bool
	BUILD_WITH_PROMETHEUS    bool
	BUILD_WITH_PROXYPROTOCOL bool
	BUILD_WITH_RATELIMIT     bool
	BUILD_WITH_REALIP        bool
	BUILD_WITH_REAUTH        bool
	BUILD_WITH_RESTIC        bool
	BUILD_WITH_UPLOAD        bool
	BUILD_WITH_WEBDAV        bool

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

	flag.BoolVar(&f.BUILD_WITH_AUTHZ, "authz", false, "Build Caddy with AUTHZ plugin")
	flag.BoolVar(&f.BUILD_WITH_AUTHZ, "awses", false, "Build Caddy with AWSES plugin")
	flag.BoolVar(&f.BUILD_WITH_AWSLAMBDA, "awslambda", false, "Build Caddy with AWSLAMBDA plugin")
	flag.BoolVar(&f.BUILD_WITH_CACHE, "cache", false, "Build Caddy with CACHE plugin")
	flag.BoolVar(&f.BUILD_WITH_CGI, "cgi", false, "Build Caddy with CGI plugin")
	flag.BoolVar(&f.BUILD_WITH_CORS, "cors", false, "Build Caddy with CORS plugin")
	flag.BoolVar(&f.BUILD_WITH_DATADOG, "datadog", false, "Build Caddy with DATADOG plugin")
	flag.BoolVar(&f.BUILD_WITH_EXPIRES, "expires", false, "Build Caddy with EXPIRES plugin")
	flag.BoolVar(&f.BUILD_WITH_FILEMANAGER, "filemanager", false, "Build Caddy with FILEMANAGER plugin. Includes HUGO and JEKYLL.")
	flag.BoolVar(&f.BUILD_WITH_FILTER, "filter", false, "Build Caddy with FILTER plugin")
	flag.BoolVar(&f.BUILD_WITH_FORWARDPROXY, "forwardproxy", false, "Build Caddy with FORWARDPROXY plugin")
	flag.BoolVar(&f.BUILD_WITH_GIT, "git", false, "Build Caddy with GIT plugin")
	flag.BoolVar(&f.BUILD_WITH_GOPKG, "gopkg", false, "Build Caddy with GOPKG plugin")
	flag.BoolVar(&f.BUILD_WITH_GRPC, "grpc", false, "Build Caddy with GRPC plugin")
	flag.BoolVar(&f.BUILD_WITH_IPFILTER, "ipfilter", false, "Build Caddy with IPFILTER plugin")
	flag.BoolVar(&f.BUILD_WITH_JSONP, "jsonp", false, "Build Caddy with JSONP plugin")
	flag.BoolVar(&f.BUILD_WITH_JWT, "jwt", false, "Build Caddy with JWT plugin")
	flag.BoolVar(&f.BUILD_WITH_LOCALE, "locale", false, "Build Caddy with LOCALE plugin")
	flag.BoolVar(&f.BUILD_WITH_MAILOUT, "mailout", false, "Build Caddy with MAILOUT plugin")
	flag.BoolVar(&f.BUILD_WITH_MINIFY, "minify", false, "Build Caddy with MINIFY plugin")
	flag.BoolVar(&f.BUILD_WITH_NOBOTS, "nobots", false, "Build Caddy with NOBOTS plugin")
	flag.BoolVar(&f.BUILD_WITH_PROMETHEUS, "prometheus", false, "Build Caddy with PROMETHEUS plugin")
	flag.BoolVar(&f.BUILD_WITH_PROXYPROTOCOL, "proxyprotocol", false, "Build Caddy with PROXYPROTOCOL plugin")
	flag.BoolVar(&f.BUILD_WITH_RATELIMIT, "ratelimit", false, "Build Caddy with RATELIMIT plugin")
	flag.BoolVar(&f.BUILD_WITH_REALIP, "realip", false, "Build Caddy with REALIP plugin")
	flag.BoolVar(&f.BUILD_WITH_REAUTH, "reauth", false, "Build Caddy with REAUTH plugin")
	flag.BoolVar(&f.BUILD_WITH_RESTIC, "restic", false, "Build Caddy with RESTIC plugin")
	flag.BoolVar(&f.BUILD_WITH_UPLOAD, "upload", false, "Build Caddy with UPLOAD plugin")
	flag.BoolVar(&f.BUILD_WITH_WEBDAV, "webdav", false, "Build Caddy with WEBDAV plugin")

	flag.BoolVar(&f.DO_NOT_REMOVE_CURRENT_GOPATH, "donotremovegopath", false, "Do not remove previously created GOPATH (improves speed)")

	// Build output path.
	flag.StringVar(&f.BUILD_OUTPUT, "output", "/usr/local/bin/", "Resulted binary. It will be recreated if already exist.")

	if len(os.Args) == 1 || os.Args[1] == "-h" {
		flag.PrintDefaults()
		fmt.Println(HELP_POSTFIX)
		os.Exit(1)
	} else {
		flag.Parse()
	}
}
