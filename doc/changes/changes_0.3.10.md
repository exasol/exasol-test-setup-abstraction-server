# Exasol Test Setup Abstraction Server 0.3.10, released 2024-11-19

Code name: Fix startup with Exasol 8.32.0

## Summary

This release fixes the startup with Exasol 8.32.0 and fixes CVE-2024-8184 in transitive dependency `org.eclipse.jetty:jetty-server`.

The release also upgrades the default Exasol version to 8.32.0.

**Note:** Transitive dependency `org.eclipse.jetty:jetty-http` contains vulnerability CVE-2024-6763. We can't upgrade to a fixed version due to API incompatibility with Javalin.

## Features

* #52: Fixed startup with Exasol 8.32.0

## Dependency Updates

### Server for the Exasol Test-Setup Abstraction

#### Compile Dependency Updates

* Updated `com.exasol:exasol-test-setup-abstraction-java:2.1.4` to `2.1.6`
* Updated `io.javalin:javalin:6.1.6` to `6.3.0`
* Added `org.eclipse.jetty:jetty-server:11.0.24`

#### Runtime Dependency Updates

* Updated `org.slf4j:slf4j-jdk14:2.0.13` to `2.0.16`

#### Test Dependency Updates

* Updated `io.rest-assured:rest-assured:5.4.0` to `5.5.0`
* Updated `org.hamcrest:hamcrest:2.2` to `3.0`
* Updated `org.itsallcode:junit5-system-extensions:1.2.0` to `1.2.2`
* Updated `org.junit.jupiter:junit-jupiter-engine:5.10.2` to `5.11.3`
* Updated `org.junit.jupiter:junit-jupiter-params:5.10.2` to `5.11.3`
* Updated `org.mockito:mockito-junit-jupiter:5.12.0` to `5.14.2`

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
* Updated `org.codehaus.mojo:exec-maven-plugin:3.2.0` to `3.5.0`
* Updated `org.codehaus.mojo:versions-maven-plugin:2.16.2` to `2.17.1`

### Go-client

#### Compile Dependency Updates

* Updated `golang:1.21` to `1.22`
* Updated `github.com/exasol/exasol-driver-go:v1.0.7` to `v1.0.10`

#### Test Dependency Updates

* Updated `github.com/antchfx/xmlquery:v1.4.0` to `v1.4.2`
