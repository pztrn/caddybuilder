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

package cmdworker

import (
    // stdlib
    "errors"
    "fmt"
    "io/ioutil"
    "os/exec"
    "strings"
)

type CmdWorker struct {}

// Executes arbitrary command.
func (cw *CmdWorker) Execute(command string) error {
    log.Printf("\tExecuting command: %s", command)
    // First parameter is a command, others are parameters.
    cmd_splitted := strings.Split(command, " ")

    // Prepare command.
    cmd := exec.Command(cmd_splitted[0], cmd_splitted[1:]...)
    stdout, _ := cmd.StdoutPipe()
    stderr, _ := cmd.StderrPipe()
    // Go, go, go!
    err := cmd.Start()
    if err != nil {
        return errors.New(fmt.Sprintf("Failed to start command '%s': %s", command, err.Error()))
    }
    stdout_output, _ := ioutil.ReadAll(stdout)
    stderr_output, _ := ioutil.ReadAll(stderr)
    // Wait until command finishes.
    err1 := cmd.Wait()
    if err1 != nil {
        // This means that some error occured in run time.
        log.Print("\tStdout:")
        log.Print(string(stdout_output))
        log.Print("\tStderr:")
        log.Print(string(stderr_output))
        return errors.New(fmt.Sprintf("Error occured while executing '%s': %s", command, err1.Error()))
    }

    return nil
}

// Initializes CmdWorker.
func (cw *CmdWorker) Initialize() {}

