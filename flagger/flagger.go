// Caddybuilder - a friendly tool to build Caddy executable.
// Copyright (c) 2017, Stanislav N. aka pztrn <pztrn at pztrn dot name>
//

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

	// Build with AWSLAMBDA: https://github.com/coopernurse/caddy-awslambda
	BUILD_WITH_AWSLAMBDA bool
	// Build with CORS: https://github.com/captncraig/cors
	BUILD_WITH_CORS bool
	// Build with EXPIRES: https://github.com/epicagency/caddy-expires
	BUILD_WITH_EXPIRES bool
	// Build with FILEMANAGER: https://github.com/hacdias/caddy-filemanager
	BUILD_WITH_FILEMANAGER bool
	// Build with GIT: https://github.com/abiosoft/caddy-git
	BUILD_WITH_GIT bool
	// Build with HUGO: https://github.com/hacdias/caddy-hugo
	BUILD_WITH_HUGO bool
	// Build with IPFILTER: https://github.com/pyed/ipfilter
	BUILD_WITH_IPFILTER bool
	// Build with JSONP: https://github.com/pschlump/caddy-jsonp
	BUILD_WITH_JSONP bool
	// Build with JWT: https://github.com/BTBurke/caddy-jwt
	BUILD_WITH_JWT bool
	// Build with LOCALE: https://github.com/simia-tech/caddy-locale
	BUILD_WITH_LOCALE bool
	// Build with MAILOUT: https://github.com/SchumacherFM/mailout
	BUILD_WITH_MAILOUT bool
	// Build with MINIFY: https://github.com/hacdias/caddy-minify
	BUILD_WITH_MINIFY bool
	// Build with MULTIPASS: https://github.com/namsral/multipass
	BUILD_WITH_MULTIPASS bool
	// Build with PROMETHEUS: https://github.com/miekg/caddy-prometheus
	BUILD_WITH_PROMETHEUS bool
	// Build with RATELIMIT: https://github.com/xuqingfeng/caddy-rate-limit
	BUILD_WITH_RATELIMIT bool
	// Build with REALIP: https://github.com/captncraig/caddy-realip
	BUILD_WITH_REALIP bool
	// Build with SEARCH: https://github.com/pedronasser/caddy-search
	BUILD_WITH_SEARCH bool
	// Build with UPLOAD: https://github.com/wmark/caddy.upload
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
	flag.StringVar(&f.BUILD_OUTPUT, "output", "/usr/local/bin/caddy", "Path where resulting binary will be placed. Specify full path with binary name (e.g. /usr/local/bin/caddy)!")

	if len(os.Args) == 1 || os.Args[1] == "-h" {
		flag.PrintDefaults()
		fmt.Println(HELP_POSTFIX)
		os.Exit(1)
	} else {
		flag.Parse()
	}
}
