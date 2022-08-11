# Exasol Test Setup Abstraction Server 0.2.3, released 2022-??-??

Code name: Add GetConnectionInfo() method to Go client

## Summary

This release makes the `GetConnectionInfo()` method of the Go client public. It allows retrieving connection parameters for the database. We also added documentation for the public methods.

## Features

* #20: Made GetConnectionInfo() in Go client public

## Dependency Updates

### Server for the Exasol Test-Setup Abstraction

#### Test Dependency Updates

* Updated `org.junit.jupiter:junit-jupiter-engine:5.8.2` to `5.9.0`
* Updated `org.junit.jupiter:junit-jupiter-params:5.8.2` to `5.9.0`

#### Plugin Dependency Updates

* Updated `com.exasol:error-code-crawler-maven-plugin:1.1.1` to `1.1.2`
* Updated `org.apache.maven.plugins:maven-enforcer-plugin:3.0.0` to `3.1.0`

### Go-client

#### Compile Dependency Updates

* Updated `golang:1.17` to `1.18`
* Updated `github.com/exasol/exasol-driver-go:v0.4.3` to `v0.4.5`

#### Test Dependency Updates

* Updated `github.com/antchfx/xmlquery:v1.3.11` to `v1.3.12`
