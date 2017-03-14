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
