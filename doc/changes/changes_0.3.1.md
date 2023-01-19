# Exasol Test Setup Abstraction Server 0.3.1, released 2023-01-18

Code name: Upgrade dependencies

## Summary

In this release we removed a workaround for fixed issue [exasol-test-setup-abstraction-java#46](https://github.com/exasol/exasol-test-setup-abstraction-java/issues/46) and upgraded dependencies.

## Features

* #27: Removed workaround for [exasol-test-setup-abstraction-java#46](https://github.com/exasol/exasol-test-setup-abstraction-java/issues/46)

## Dependency Updates

### Server for the Exasol Test-Setup Abstraction

#### Compile Dependency Updates

* Updated `com.exasol:exasol-test-setup-abstraction-java:1.0.0` to `2.0.0`
* Updated `io.javalin:javalin:5.1.2` to `5.3.1`

#### Test Dependency Updates

* Updated `io.rest-assured:rest-assured:5.2.0` to `5.3.0`
* Updated `org.junit.jupiter:junit-jupiter-engine:5.9.1` to `5.9.2`
* Updated `org.junit.jupiter:junit-jupiter-params:5.9.1` to `5.9.2`
* Updated `org.mockito:mockito-junit-jupiter:4.8.1` to `5.0.0`

#### Plugin Dependency Updates

* Updated `com.exasol:artifact-reference-checker-maven-plugin:0.4.0` to `0.4.2`
* Updated `com.exasol:error-code-crawler-maven-plugin:1.1.2` to `1.2.1`
* Updated `io.github.zlika:reproducible-build-maven-plugin:0.15` to `0.16`
* Updated `org.apache.maven.plugins:maven-assembly-plugin:3.3.0` to `3.4.2`
* Updated `org.apache.maven.plugins:maven-failsafe-plugin:3.0.0-M5` to `3.0.0-M7`
* Updated `org.apache.maven.plugins:maven-jar-plugin:3.2.2` to `3.3.0`
* Updated `org.apache.maven.plugins:maven-surefire-plugin:3.0.0-M5` to `3.0.0-M7`
* Updated `org.codehaus.mojo:exec-maven-plugin:3.0.0` to `3.1.0`
* Updated `org.codehaus.mojo:flatten-maven-plugin:1.2.7` to `1.3.0`
* Updated `org.codehaus.mojo:versions-maven-plugin:2.10.0` to `2.13.0`

### Go-client

#### Test Dependency Updates

* Updated `github.com/antchfx/xmlquery:v1.3.12` to `v1.3.14`
