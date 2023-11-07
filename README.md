# eolib-go

Core library for writing Endless Online applications using the Go programming language.

## Usage

### Referencing the package

The library may be referenced in a project via go get:

```
go get github.com/ethanmoffat/eolib-go@v1.1.2
```

### Sample code

A sample server skeleton using eolib-go is [available here](https://gist.github.com/ethanmoffat/95eed4ef0eeb524c8a505acb1bcbf956).

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

Building the library on Windows is left as an exercise to the reader.
