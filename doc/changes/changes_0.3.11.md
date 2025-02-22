# Exasol Test Setup Abstraction Server 0.3.11, released 2025-02-11

Code name: Fix CVE-2024-45338 in `golang.org/x/net`

## Summary

This release fixes vulnerability CVE-2024-45338 in dependency `golang.org/x/net`.

## Security

* #54: Fixed CVE-2024-45338 in dependency `golang.org/x/net`

## Dependency Updates

### Server for the Exasol Test-Setup Abstraction

#### Compile Dependency Updates

* Updated `com.exasol:exasol-test-setup-abstraction-java:2.1.6` to `2.1.7`
* Updated `io.javalin:javalin:6.3.0` to `6.4.0`
* Removed `org.eclipse.jetty:jetty-server:11.0.24`

#### Test Dependency Updates

* Updated `org.junit.jupiter:junit-jupiter-engine:5.11.3` to `5.11.4`
* Updated `org.junit.jupiter:junit-jupiter-params:5.11.3` to `5.11.4`
* Updated `org.mockito:mockito-junit-jupiter:5.14.2` to `5.15.2`

#### Plugin Dependency Updates

* Updated `org.apache.maven.plugins:maven-failsafe-plugin:3.5.1` to `3.5.2`
* Updated `org.apache.maven.plugins:maven-site-plugin:3.9.1` to `3.21.0`
* Updated `org.apache.maven.plugins:maven-surefire-plugin:3.5.1` to `3.5.2`
* Updated `org.codehaus.mojo:versions-maven-plugin:2.17.1` to `2.18.0`
* Updated `org.sonarsource.scanner.maven:sonar-maven-plugin:4.0.0.4121` to `5.0.0.4389`

### Go-client

#### Compile Dependency Updates

* Updated `golang:1.22` to `1.23`
* Updated `github.com/stretchr/testify:v1.9.0` to `v1.10.0`
* Updated `github.com/exasol/exasol-driver-go:v1.0.10` to `v1.0.12`

#### Test Dependency Updates

* Updated `github.com/antchfx/xmlquery:v1.4.2` to `v1.4.4`
