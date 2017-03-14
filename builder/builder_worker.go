// Caddybuilder - a friendly tool to build Caddy executable.
// Copyright (c) 2017, Stanislav N. aka pztrn <pztrn at pztrn dot name>
//

// File "builder_worker.go" contains worker functions which related
// to commands executing.
package builder

import (
    // stdlib
    "errors"
    "fmt"
    "io/ioutil"
    "os"
    "os/exec"
    "path/filepath"
    "strings"
)

// Actually build Caddy binary.
func (b *Builder) buildCaddy() error {
    command := fmt.Sprintf("%s build -a -o %s %s/src/github.com/mholt/caddy/caddy", b.NeccessaryPrograms["go"], flags.BUILD_OUTPUT, b.Workspace)
    log.Printf("\tExecuting: %s", command)

    //os.Chdir(fmt.Sprintf("%s/src/github.com/mholt/caddy/caddy/", b.Workspace))
    // Prepare command.
    cmd := exec.Command(b.NeccessaryPrograms["go"], "build", "-a", "-o", flags.BUILD_OUTPUT, "github.com/mholt/caddy/caddy/")
    //stdout, _ := cmd.StdoutPipe()
    stderr, _ := cmd.StderrPipe()
    // Go, go, go!
    err := cmd.Start()
    if err != nil {
        return errors.New(fmt.Sprintf("Failed to build Caddy: %s", err.Error()))
    }
    // Wait until command finishes.
    err1 := cmd.Wait()
    if err1 != nil {
        // This means that some error occured in run time.
        stderr_output, _ := ioutil.ReadAll(stderr)
        log.Print("\tCommand output:")
        log.Print(string(stderr_output))
        return errors.New(fmt.Sprintf("Error occured while building Caddy: %s", err1.Error()))
    }

    return nil
}

// Get sources from passed path.
func (b *Builder) goGet(path string) error {
    command := fmt.Sprintf("%s get -d -u %s", b.NeccessaryPrograms["go"], path)
    log.Printf("\tExecuting: %s", command)

    // Prepare command.
    cmd := exec.Command(b.NeccessaryPrograms["go"], "get", "-d", "-u", path)
    //stdout, _ := cmd.StdoutPipe()
    stderr, _ := cmd.StderrPipe()
    // Go, go, go!
    err := cmd.Start()
    if err != nil {
        return errors.New(fmt.Sprintf("Failed to start command '%s': %s", command, err.Error()))
    }
    // Wait until command finishes.
    err1 := cmd.Wait()
    if err1 != nil {
        // This means that some error occured in run time.
        stderr_output, _ := ioutil.ReadAll(stderr)
        log.Print("\tCommand output:")
        log.Print(string(stderr_output))
        return errors.New(fmt.Sprintf("Error occured while executing '%s': %s", command, err1.Error()))
    }

    return nil
}

// Replace default string with plugin import in run.go.
// ToDo: refactor it.
func (b *Builder) installPlugin(name string) {
    // Get plugin remote path.
    // ToDo: refactor, it's shitty.
    var plugin_path string = ""
    if flags.BUILD_WITH_AWSLAMBDA && name == "awslambda" {
        plugin_path = "github.com/coopernurse/caddy-awslambda"
    }
    if flags.BUILD_WITH_CORS && name == "cors" {
        plugin_path = "github.com/captncraig/cors"
    }
    if flags.BUILD_WITH_EXPIRES && name == "expires" {
        plugin_path = "github.com/epicagency/caddy-expires"
    }
    if flags.BUILD_WITH_FILEMANAGER && name == "filemanager" {
        plugin_path = "github.com/hacdias/caddy-filemanager"
    }
    if flags.BUILD_WITH_GIT && name == "git" {
        plugin_path = "github.com/abiosoft/caddy-git"
    }
    if flags.BUILD_WITH_HUGO && name == "hugo" {
        plugin_path = "github.com/hacdias/caddy-hugo"
    }
    if flags.BUILD_WITH_IPFILTER && name == "ipfilter" {
        plugin_path = "github.com/pyed/ipfilter"
    }
    if flags.BUILD_WITH_JSONP && name == "jsonp" {
        plugin_path = "github.com/pschlump/caddy-jsonp"
    }
    if flags.BUILD_WITH_JWT && name == "jwt" {
        plugin_path = "github.com/BTBurke/caddy-jwt"
    }
    if flags.BUILD_WITH_LOCALE && name == "locale" {
        plugin_path = "github.com/simia-tech/caddy-locale"
    }
    if flags.BUILD_WITH_MAILOUT && name == "mailout" {
        plugin_path = "github.com/SchumacherFM/mailout"
    }
    if flags.BUILD_WITH_MINIFY && name == "minify" {
        plugin_path = "github.com/hacdias/caddy-minify"
    }
    if flags.BUILD_WITH_MULTIPASS && name == "multipass" {
        plugin_path = "github.com/namsral/multipass"
    }
    if flags.BUILD_WITH_PROMETHEUS && name == "prometheus" {
        plugin_path = "github.com/miekg/caddy-prometheus"
    }
    if flags.BUILD_WITH_RATELIMIT && name == "ratelimit" {
        plugin_path = "github.com/xuqingfeng/caddy-rate-limit"
    }
    if flags.BUILD_WITH_REALIP && name == "realip" {
        plugin_path = "github.com/captncraig/caddy-realip"
    }
    if flags.BUILD_WITH_SEARCH && name == "search" {
        plugin_path = "github.com/pedronasser/caddy-search"
    }
    if flags.BUILD_WITH_UPLOAD && name == "upload" {
        plugin_path = "github.com/wmark/caddy.upload"
    }

    // Do nothing if nothing to install.
    if len(plugin_path) == 0 {
        return
    }

    // Path to run.go.
    rungo := filepath.Join(b.Workspace, "src", "github.com", "mholt", "caddy", "caddy", "caddymain", "run.go")
    // Read file.
    fh_bytes, err := ioutil.ReadFile(rungo)
    if err != nil {
        log.Fatalf("Cannot open run.go: %s", err.Error())
    }
    fh := string(fh_bytes)

    log.Printf("Installing plugin: %s", name)

    err1 := b.goGet(plugin_path)
    if err1 != nil {
        log.Fatalf("Failed to get plugin's sources: %s", err1.Error())
    }

    // Replace default "This is where other plugins get plugged in (imported)"
    // line with plugin import.
    replace_to := fmt.Sprintf("_ \"%s\"\n\t// This is where other plugins get plugged in (imported)", plugin_path)
    fh = strings.Replace(fh, "// This is where other plugins get plugged in (imported)", replace_to, 1)
    // Write file.
    ioutil.WriteFile(rungo, []byte(fh), os.ModePerm)
}
