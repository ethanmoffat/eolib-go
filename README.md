# eolib-go

Core library for writing Endless Online applications using the Go programming language.

## Usage

### Referencing the package

The library may be referenced in a project via go get:

```
go get github.com/ethanmoffat/eolib-go/v3
```

### Sample code

A sample server skeleton using eolib-go is [available here](https://gist.github.com/ethanmoffat/95eed4ef0eeb524c8a505acb1bcbf956).

### Switch data

Some packets and structs have a code field (e.g. `LoginReplyServerPacket.ReplyCode`) and a data field (e.g. `ReplyCodeData`) that must hold the data type matching the code. Generated factory functions set both together. They are named `New<Type>With<Value>`, where `<Type>` is the type name without the `ClientPacket`/`ServerPacket` suffix:

```go
// a case with data takes the case data
ok := server.NewLoginReplyWithOk(&server.LoginReplyReplyCodeDataOk{Characters: characters})

// a code without data to set takes no parameters
wrongUser := server.NewLoginReplyWithWrongUser()

// nested switches are flattened, and every code along the path is set
banned := server.NewInitInitWithBannedTemporary(&server.InitInitBanTypeDataTemporary{MinutesRemaining: 30})

// a default case takes the code, and returns an error if the code has its own case
reply, err := server.NewAccountReplyWithDefault(sessionId, &server.AccountReplyReplyCodeDataDefault{SequenceStart: 12})
```

Numeric cases with data are named after their data type, e.g. `server.NewInitInitWithBanTypeData0`. If the data passed to a factory is nil, the data field is left nil and serializing the result returns an error.

To read the data, use a type assertion or type switch on the data field. Checking the code as well isn't necessary: deserialization always sets the data type matching the code, and serialization rejects data that doesn't match it.

```go
if ok, isOk := p.ReplyCodeData.(*server.LoginReplyReplyCodeDataOk); isOk {
	fmt.Println("characters:", len(ok.Characters))
}
```

## Development Environment

### Installing go

Development was done using go 1.20.5 on Ubuntu Linux. Development on Windows is untested.

[gvm](https://github.com/moovweb/gvm) is recommended as a mechanism to manage installations of different versions of go.

To set up gvm on Linux using bash:
```bash
bash < <(curl -s -S -L https://raw.githubusercontent.com/moovweb/gvm/master/binscripts/gvm-installer)
source $HOME/.gvm/scripts/gvm
```

The `source` command may be added to your `~/.bash_profile` so gvm is available in all login shells.

```bash
echo "source $HOME/.gvm/scripts/gvm" >> ~/.bash_profile
```

Once gvm is installed and sourced, use the following commands to install and use go 1.20.5:

```bash
gvm install go1.20.5 -B
gvm use go1.20.5
```

The `gvm use` command is required on each subsequent terminal or session in which `go` commands must be run.

### VSCode setup

If go is installed via gvm, you must update your `$GOROOT` variable to point at it, or VSCode won't be able to find the go binary.

Get the path to the GOROOT via `go env`:

```bash
~ [ethan@BEASTMODE] $  gvm use go1.20.5
Now using version go1.20.5
~ [ethan@BEASTMODE] $  go env GOROOT
/home/ethan/.gvm/gos/go1.20.5
```

In VSCode, search for the `GOROOT` setting and update the path to point at this version of go.

Note that a gvm extension exists for VSCode but it is out of date and no longer functions properly.

### Building the code

A `Makefile` is provided to ease the process of building, testing, and code generation. Use `make help` to see all available targets. Running `make` by itself should be enough for most uses. `make test` is also available to run all tests.

`make` runs `golangci-lint`; install `v2.11.4` before building locally. The recommended binary install is:

```bash
curl -sSfL https://golangci-lint.run/install.sh | sh -s -- -b "$(go env GOPATH)/bin" v2.11.4
```

The GitHub Actions workflow installs the same pinned version in CI before running the build.

Building the library on Windows is left as an exercise to the reader.

## Versioning and releases

eolib-go uses [Semantic Versioning](https://semver.org/) (`MAJOR.MINOR.PATCH`, with an optional `-beta.N` or `-rc.N`
suffix). The current major version lives at the repository root, with the `/vN` module path suffix (e.g.
`github.com/ethanmoffat/eolib-go/v3`). Version 1 (`github.com/ethanmoffat/eolib-go`) is only available from its
`v1.x.y` tags, and the 2.x tags are yanked. Changes are tracked in [CHANGELOG.md](CHANGELOG.md), following
[Keep a Changelog](https://keepachangelog.com/en/1.1.0/).

To release a new version:

1. Make sure the `[Unreleased]` section of `CHANGELOG.md` lists the changes.
2. Run `./scripts/prepare-release.sh x.y.z-suffix --tag`. It updates every file that references the version (the
   changelog section and links, and the module path in the README's `go get` command), runs
   `scripts/validate-release.sh`, commits the changes as "Release x.y.z-suffix" and creates the tag. Use `--commit` to
   commit without tagging, `--date` to set the changelog date, or no option to only update the files for review.
3. Push master, and wait for CI to pass.
4. Push the tag `vx.y.z-suffix`. The release workflow runs `validate-release.sh` again, which also checks that the
   commit is on `origin/master`, before building anything. It then builds and tests the code, requests the new version
   from the Go module proxy so it is available to `go get`, and publishes a GitHub release. It is marked as a
   prerelease when the version has a suffix.
