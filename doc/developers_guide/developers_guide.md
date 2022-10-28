# Developers Guide

## Running Build Locally

```shell
./build.sh
```

## Server

### Install Server

To install the server to `$HOME/.test-setup-abstraction-server/`, run

```shell
cd server
mvn install
```

## Go Client

### Run tests

You must [install the server](#install-server) before running the client tests.

```shell
cd go-client
go test -p 1 ./...
```

### Linter

To install golangci-lint on your machine, follow [these instruction](https://golangci-lint.run/usage/install/#local-installation). Then run

```shell
cd go-client
golangci-lint run
```

## Building Releases

Use [release-droid](https://github.com/exasol/release-droid) as usual for building releases.

To allow Go projects to use the go-client module, additionally create and push a lightweight tag for the new version:

```shell
version=0.2.1
tag=go-client/v$version
git tag $tag
git push origin $tag
```

See [the Go documentation](https://go.dev/doc/modules/managing-source#multiple-module-source) for details.
