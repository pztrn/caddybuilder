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

// File "exported.go" contains New() call.
package builder

import (
	// stdlib
	l "log"

	// local
	"github.com/pztrn/caddybuilder/flagger"
)

var (
	flags *flagger.Flagger
	log   *l.Logger
)

func New(f *flagger.Flagger, l *l.Logger) *Builder {
	flags = f
	log = l
	return &Builder{}
}
