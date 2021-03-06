# Server for the Exasol Test-Setup Abstraction 0.2.0, released 2022-06-02

Code name: Add BucketFS support

## Summary

This release adds support for accessing BucketFS, allows creating a DB connection with autocommit deactivated and fixes the failing startup when docker is not running on localhost. We also upgraded project-keeper to generate dependency list and dependency changes also for the Go module.

## Features

* #7: Support for creating a connection with no autocommit
* #9: Added API for accessing BucketFS
* #3: Upgraded to latest project-keeper to generate dependency information for the Go module

## Bug Fixes

* #5: Fixed failing startup when docker is not running on localhost

## Dependency Updates

### Server for the Exasol Test-Setup Abstraction

#### Compile Dependency Updates

* Updated `com.exasol:exasol-test-setup-abstraction-java:0.3.0` to `0.3.2`
* Updated `io.javalin:javalin:4.5.0` to `4.6.0`

#### Test Dependency Updates

* Added `commons-codec:commons-codec:1.15`
* Updated `io.rest-assured:rest-assured:5.0.1` to `5.1.0`

#### Plugin Dependency Updates

* Added `org.apache.maven.plugins:maven-failsafe-plugin:3.0.0-M5`

### Go-client

#### Compile Dependency Updates

* Updated `github.com/exasol/exasol-driver-go:v0.4.0` to `v0.4.2`

#### Test Dependency Updates

* Added `github.com/antchfx/xmlquery:v1.3.11`
