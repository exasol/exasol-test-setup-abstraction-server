# Exasol Test Setup Abstraction Server 0.3.0, released 2022-10-28

Code name: Specify Exasol version

## Summary

In this release we changed the way how to configure the test setup. Instead of calling `Create("config.json")` you now use `New().CloudSetupConfigFilePath("config.json").Start()`. This emphasizes that the config file is optional and it allows specifying the Exasol Docker DB version to use with `New().DockerDbVersion("7.1.15").Start()`.

## Features

* #22: Added option to specify the Exasol Docker DB version

## Dependency Updates

### Server for the Exasol Test-Setup Abstraction

#### Compile Dependency Updates

* Updated `com.exasol:exasol-test-setup-abstraction-java:0.3.2` to `1.0.0`
* Updated `io.javalin:javalin:4.6.4` to `5.1.2`

#### Test Dependency Updates

* Removed `commons-codec:commons-codec:1.15`
* Updated `io.rest-assured:rest-assured:5.1.1` to `5.2.0`
* Updated `org.junit.jupiter:junit-jupiter-engine:5.9.0` to `5.9.1`
* Updated `org.junit.jupiter:junit-jupiter-params:5.9.0` to `5.9.1`
* Updated `org.mockito:mockito-junit-jupiter:4.7.0` to `4.8.1`

### Go-client

#### Compile Dependency Updates

* Updated `github.com/exasol/exasol-driver-go:v0.4.5` to `v0.4.6`

#### Test Dependency Updates

* Updated `github.com/stretchr/testify:v1.8.0` to `v1.8.1`
