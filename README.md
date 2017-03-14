# Caddy Builder

This application builds Caddy webserver with neccessary plugins.
Great for deploying automation and already used in production.

This project was created after failing to obtain code which is used
on https://caddyserver.com/download to make Caddy builds with
custom feature set.

After I wrote the very first version of Caddybuilder I found
https://github.com/caddyserver/devportal, but it's not what I want.

# Installing

It is enough to do:

```
go get github.com/pztrn/caddybuilder
```

**ONLY *NIX SYSTEMS ARE SUPPORTED FOR NOW!**

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

# Testing

Issue this command to execute all tests:

```
go test -test.v ./...
```

# ToDo

- [ ] Preserve Caddybuilder's GOPATH between launches.
- [ ] More perfect plugins subsystem.
- [ ] Go away from standart ``log`` module.
- [ ] Fix inability to get processes output on error.
- [ ] Configuration files with possibility to pin to tag or revision of Caddy
and/or plugins.
- [ ] Support for Windows (PRs welcome, I have none of them).
- [ ] Switch to native Git library for working with git?
