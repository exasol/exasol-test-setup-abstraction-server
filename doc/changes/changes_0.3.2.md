# Exasol Test Setup Abstraction Server 0.3.2, released 2023-01-31

Code name: Improve Docker Container Reuse

## Summary

This release avoids starting new Docker containers for each working directory in which the server was started.

## Bugfixes

* #31: Fixed starting new Docker containers for each working directory
## Dependency Updates

### Server for the Exasol Test-Setup Abstraction

#### Compile Dependency Updates

* Added `com.exasol:exasol-testcontainers:6.5.1`
* Updated `io.javalin:javalin:5.3.1` to `5.3.2`

#### Test Dependency Updates

* Updated `org.mockito:mockito-junit-jupiter:5.0.0` to `5.1.1`

### Go-client

#### Test Dependency Updates

* Updated `github.com/antchfx/xmlquery:v1.3.14` to `v1.3.15`
