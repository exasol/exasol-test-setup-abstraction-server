# Developers Guide

## Running Build Locally

```shell
./build.sh
```

## Server

To install the Server to `$HOME/.test-setup-abstraction-server`, run

```shell
cd server
mvn install
```

## Go Client

### Run tests

```shell
cd go-client
go test ./...
```

### Linter

To install golangci-lint on your machine, follow [these instruction](https://golangci-lint.run/usage/install/#local-installation).

```shell
cd go-client
golangci-lint run
```
