# Caddy Builder

This application builds Caddy webserver with neccessary plugins.
Great for deploying automation and already used in production.

This project was created after failing to obtain code which is used
on https://caddyserver.com/download to make Caddy builds with
custom feature set.

After I wrote the very first version of Caddybuilder I found
https://github.com/caddyserver/devportal, but it's not what I want.

**Note:** there are some plugins that missing from official Caddy
website. They might or might not work!

Repository on Github is a mirror. Real development is going at
[pztrn's Lab](https://lab.pztrn.name/pztrn/caddybuilder).
Please, report bugs there.

# Installing

It is enough to do:

```
go get lab.pztrn.name/pztrn/caddybuilder
```

**ONLY NIX SYSTEMS ARE SUPPORTED FOR NOW!**

# Running

Execute this command to get list of available options:

```
$GOPATH/bin/caddybuilder -h
```

By default no additional plugins will be installed, so you should pass
some flags to caddybuilder binary to get plugin included.

**Note:** by default every new Caddy build will be placed as
``/usr/local/bin/caddy``. Specify ``-output`` parameter with value
to overwrite it!

# Developing

For now I'm commiting to master branch, until first release. After that
master branch will be "always stable".

To make a PR you should fork this repository, create own branch (e.g.
"$nick_changes"), do some coding, and make a PR against "development"
branch.

Every commit should be covered by tests. PRs with new functions and without
tests are unacceptable and will be rejected.

## Plugins

Take a look into any plugin, they're pretty straightforward. ``context`` and
``handler`` packages are not plugins, but:

* ``context`` is a structure with some useful things that passed to plugins.
* ``handler`` is an interface for plugins.

Plugins require no tests to be written, instead it will be tested as part
of ``plugins`` package. It will test all plugins that have been initialized
in ``plugins/plugins.go``. If any error arises - ``go test`` will print it
out.

If plugin require separate test (e.g. it do some requests or doint something
more that replacing "default replace line" in ``run.go``) - test should be
written in ``plugins/plugins_test.go`` and named like ``TestNamePluginAction``,
where:

* ``Name`` - should be replaced with plugin name (e.g. "Realip").
* ``Action`` - should be replaced to function name or action name.

# Testing

Issue this command to execute all tests:

```
go test -test.v ./...
```

Full tests execution takes approx. 1-2 minutes. Specify ``-short`` to skip
tests for ``go get`` execution and Caddy building. This can be useful if
you're writing a plugin:

```
go test -test.v -short ./...
```

# ToDos

## 0.1.0

- [ ] Preserve Caddybuilder's GOPATH between launches.
- [x] More perfect plugins subsystem.
- [x] Fix inability to get processes output on error.

## 0.2.0

- [ ] Configuration files with possibility to pin to tag or revision of Caddy
and/or plugins.
- [ ] Go away from standart ``log`` module.

## 0.3.0

- [ ] Support for Windows (PRs welcome, I have none of them).

## 0.4.0

- [ ] Switch to native Git library for working with git?
