# Exasol Test Setup Abstraction Server 0.3.7, released 2024-04-09

Code name: Fix CVE-2024-29025 in io.netty:netty-codec-http:jar:4.1.107.Final:runtime

## Summary

This release fixes vulnerability CVE-2024-29025 in `io.netty:netty-codec-http:jar:4.1.107.Final:runtime`.

## Security

* #46: Fixed CVE-2024-29025 in `io.netty:netty-codec-http:jar:4.1.107.Final:runtime`

## Dependency Updates

### Server for the Exasol Test-Setup Abstraction

#### Compile Dependency Updates

* Updated `com.exasol:exasol-test-setup-abstraction-java:2.1.1` to `2.1.2`

#### Plugin Dependency Updates

* Updated `com.exasol:error-code-crawler-maven-plugin:2.0.0` to `2.0.2`
* Updated `org.apache.maven.plugins:maven-assembly-plugin:3.6.0` to `3.7.1`
* Updated `org.apache.maven.plugins:maven-compiler-plugin:3.12.1` to `3.13.0`
* Updated `org.codehaus.mojo:exec-maven-plugin:3.1.0` to `3.2.0`
* Updated `org.jacoco:jacoco-maven-plugin:0.8.11` to `0.8.12`
* Updated `org.sonarsource.scanner.maven:sonar-maven-plugin:3.10.0.2594` to `3.11.0.3922`

### Go-client

#### Compile Dependency Updates

* Updated `github.com/exasol/exasol-driver-go:v1.0.4` to `v1.0.6`

#### Test Dependency Updates

* Updated `github.com/antchfx/xmlquery:v1.3.18` to `v1.4.0`
