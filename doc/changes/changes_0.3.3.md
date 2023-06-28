# Exasol Test Setup Abstraction Server 0.3.3, released 2023-06-29

Code name: Upgrade dependencies on top of 0.3.2

## Summary

This release fixes CVE-2022-41723 in dependency `golang.org/x/net` and CVE-2023-34462 in `io.netty:netty-handler`. It also upgrades the default Exasol version to 7.1.19.

## Bugfixes

* #34: Updated dependencies
* #36: Updated dependencies

## Dependency Updates

### Server for the Exasol Test-Setup Abstraction

#### Compile Dependency Updates

* Updated `com.exasol:exasol-test-setup-abstraction-java:2.0.0` to `2.0.2`
* Removed `com.exasol:exasol-testcontainers:6.5.1`
* Updated `io.javalin:javalin:5.3.2` to `5.6.1`

#### Runtime Dependency Updates

* Updated `org.slf4j:slf4j-jdk14:2.0.6` to `2.0.7`

#### Test Dependency Updates

* Updated `io.rest-assured:rest-assured:5.3.0` to `5.3.1`
* Updated `org.junit.jupiter:junit-jupiter-engine:5.9.2` to `5.9.3`
* Updated `org.junit.jupiter:junit-jupiter-params:5.9.2` to `5.9.3`
* Updated `org.mockito:mockito-junit-jupiter:5.1.1` to `5.4.0`

#### Plugin Dependency Updates

* Updated `com.exasol:error-code-crawler-maven-plugin:1.2.1` to `1.2.3`
* Updated `org.apache.maven.plugins:maven-assembly-plugin:3.4.2` to `3.5.0`
* Updated `org.apache.maven.plugins:maven-compiler-plugin:3.10.1` to `3.11.0`
* Updated `org.apache.maven.plugins:maven-enforcer-plugin:3.1.0` to `3.3.0`
* Updated `org.apache.maven.plugins:maven-failsafe-plugin:3.0.0-M7` to `3.0.0`
* Updated `org.apache.maven.plugins:maven-surefire-plugin:3.0.0-M7` to `3.0.0`
* Added `org.basepom.maven:duplicate-finder-maven-plugin:1.5.1`
* Updated `org.codehaus.mojo:flatten-maven-plugin:1.3.0` to `1.4.1`
* Updated `org.codehaus.mojo:versions-maven-plugin:2.13.0` to `2.15.0`
* Updated `org.jacoco:jacoco-maven-plugin:0.8.8` to `0.8.9`

### Go-client

#### Compile Dependency Updates

* Updated `golang:1.18` to `1.19`
* Updated `github.com/exasol/exasol-driver-go:v0.4.6` to `v1.0.0`

#### Test Dependency Updates

* Updated `github.com/antchfx/xmlquery:v1.3.15` to `v1.3.17`
* Updated `github.com/stretchr/testify:v1.8.1` to `v1.8.4`
