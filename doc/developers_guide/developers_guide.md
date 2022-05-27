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
go test ./...
```

### Linter

To install golangci-lint on your machine, follow [these instruction](https://golangci-lint.run/usage/install/#local-installation). Then run

```shell
cd go-client
golangci-lint run
```
