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

package cmdworker

import (
	// stdlib
	"errors"
	"fmt"
	"io/ioutil"
	"os/exec"
	"strings"
)

type CmdWorker struct{}

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
