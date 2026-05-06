# Exasol Test Setup Abstraction Server 1.0.1, released 2026-??-??

Code name: Fix CVE-2026-32287 in Go dependency `github.com/antchfx/xpath`

## Summary

This release updates dependencies to fix CVE-2026-32287 in Go dependency `github.com/antchfx/xpath`.

## Security

* #68: Fixed CVE-2026-32287 in Go dependency `github.com/antchfx/xpath`

## Dependency Updates

### Server for the Exasol Test-Setup Abstraction

#### Compile Dependency Updates

* Updated `com.exasol:exasol-test-setup-abstraction-java:2.1.10` to `2.1.11`
* Updated `io.javalin:javalin:6.7.0` to `7.2.0`

#### Test Dependency Updates

* Added `org.junit.jupiter:junit-jupiter-api:6.0.3`
* Removed `org.junit.jupiter:junit-jupiter-engine:6.0.1`
* Updated `org.junit.jupiter:junit-jupiter-params:6.0.1` to `6.0.3`
* Updated `org.mockito:mockito-junit-jupiter:5.21.0` to `5.23.0`

#### Plugin Dependency Updates

* Updated `com.exasol:error-code-crawler-maven-plugin:2.0.5` to `2.0.7`
* Updated `io.github.git-commit-id:git-commit-id-maven-plugin:9.0.2` to `10.0.0`
* Updated `org.apache.maven.plugins:maven-assembly-plugin:3.7.1` to `3.8.0`
* Updated `org.apache.maven.plugins:maven-compiler-plugin:3.14.1` to `3.15.0`
* Updated `org.apache.maven.plugins:maven-failsafe-plugin:3.5.4` to `3.5.5`
* Updated `org.apache.maven.plugins:maven-jar-plugin:3.4.2` to `3.5.0`
* Updated `org.apache.maven.plugins:maven-resources-plugin:3.3.1` to `3.5.0`
* Updated `org.apache.maven.plugins:maven-surefire-plugin:3.5.4` to `3.5.5`
* Updated `org.codehaus.mojo:versions-maven-plugin:2.19.1` to `2.21.0`
* Updated `org.sonarsource.scanner.maven:sonar-maven-plugin:5.2.0.4988` to `5.5.0.6356`

### Go-client

#### Compile Dependency Updates

* Updated `golang:1.24.0` to `1.25.0`
* Updated `github.com/exasol/exasol-driver-go:v1.0.14` to `v1.0.16`

#### Test Dependency Updates

* Updated `github.com/antchfx/xmlquery:v1.5.0` to `v1.5.1`

#### Other Dependency Updates

* Updated `toolchain:go1.25.5` to `go1.25.9`
