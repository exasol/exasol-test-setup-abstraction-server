# Exasol Test Setup Abstraction Server 0.3.10, released 2024-11-19

Code name: Fix startup with Exasol 8.32.0

## Summary

This release fixes the startup with Exasol 8.32.0.

The release also upgrades the default Exasol version to 8.32.0.

**Breaking Change:** Starting with this release, Java 17 is required for running the project.

## Features

* #52: Fixed startup with Exasol 8.32.0

## Dependency Updates

### Server for the Exasol Test-Setup Abstraction

#### Plugin Dependency Updates

* Added `com.exasol:quality-summarizer-maven-plugin:0.2.0`
* Updated `io.github.zlika:reproducible-build-maven-plugin:0.16` to `0.17`
* Updated `org.apache.maven.plugins:maven-clean-plugin:2.5` to `3.4.0`
* Updated `org.apache.maven.plugins:maven-failsafe-plugin:3.2.5` to `3.5.1`
* Updated `org.apache.maven.plugins:maven-install-plugin:2.4` to `3.1.3`
* Updated `org.apache.maven.plugins:maven-jar-plugin:3.4.1` to `3.4.2`
* Updated `org.apache.maven.plugins:maven-resources-plugin:2.6` to `3.3.1`
* Updated `org.apache.maven.plugins:maven-site-plugin:3.3` to `3.9.1`
* Updated `org.apache.maven.plugins:maven-surefire-plugin:3.2.5` to `3.5.1`
* Updated `org.codehaus.mojo:versions-maven-plugin:2.16.2` to `2.17.1`

### Go-client

#### Compile Dependency Updates

* Updated `golang:1.21` to `1.22`
* Updated `github.com/exasol/exasol-driver-go:v1.0.7` to `v1.0.10`

#### Test Dependency Updates

* Updated `github.com/antchfx/xmlquery:v1.4.0` to `v1.4.2`
