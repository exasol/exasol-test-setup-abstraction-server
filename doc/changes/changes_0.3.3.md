# Exasol Test Setup Abstraction Server 0.3.3, released 2023-04-??

Code name: Upgrade dependencies on top of 0.3.2

## Summary

This release fixes CVE-2022-41723 in dependency `golang.org/x/net`.

## Features

* #34: Updated dependencies

## Dependency Updates

### Server for the Exasol Test-Setup Abstraction

#### Compile Dependency Updates

* Updated `io.javalin:javalin:5.3.2` to `5.4.2`

#### Runtime Dependency Updates

* Updated `org.slf4j:slf4j-jdk14:2.0.6` to `2.0.7`

#### Test Dependency Updates

* Updated `org.mockito:mockito-junit-jupiter:5.1.1` to `5.2.0`

#### Plugin Dependency Updates

* Updated `com.exasol:error-code-crawler-maven-plugin:1.2.1` to `1.2.2`
* Updated `org.apache.maven.plugins:maven-assembly-plugin:3.4.2` to `3.5.0`
* Updated `org.apache.maven.plugins:maven-enforcer-plugin:3.1.0` to `3.2.1`
* Updated `org.apache.maven.plugins:maven-failsafe-plugin:3.0.0-M7` to `3.0.0-M8`
* Updated `org.apache.maven.plugins:maven-surefire-plugin:3.0.0-M7` to `3.0.0-M8`
* Updated `org.codehaus.mojo:versions-maven-plugin:2.13.0` to `2.14.2`

### Go-client

#### Compile Dependency Updates

* Updated `golang:1.18` to `1.19`
* Updated `github.com/exasol/exasol-driver-go:v0.4.6` to `v0.4.7`

#### Test Dependency Updates

* Updated `github.com/stretchr/testify:v1.8.1` to `v1.8.2`
