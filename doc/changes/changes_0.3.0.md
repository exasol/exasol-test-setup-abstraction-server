# Exasol Test Setup Abstraction Server 0.3.0, released 2022-10-27

Code name: Specify Exasol version

## Summary

In this release we changed the way how to configure the test setup. Instead of calling `Create("config.json")` you now use `New().CloudSetupConfigFilePath("config.json").Start()`. This emphasizes that the config file is optional and it allows specifying the Exasol Docker DB version to use with `New().DockerDbVersion("7.1.15").Start()`.

## Features

* #22: Added option to specify the Exasol Docker DB version

## Dependency Updates

### Server for the Exasol Test-Setup Abstraction

#### Test Dependency Updates

* Updated `org.mockito:mockito-junit-jupiter:4.6.1` to `4.7.0`
