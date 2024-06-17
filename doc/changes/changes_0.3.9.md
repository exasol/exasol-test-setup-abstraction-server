# Exasol Test Setup Abstraction Server 0.3.9, released 2024-06-17

Code name: Fix startup with DockerDB 8.27.0

## Summary

This release fixes startup of DockerDB version 8.27.0 and later.

## Bugfixes

* #50: Fixed startup of DockerDB 8.27.0

## Dependency Updates

### Server for the Exasol Test-Setup Abstraction

#### Compile Dependency Updates

* Updated `com.exasol:exasol-test-setup-abstraction-java:2.1.3` to `2.1.4`
* Updated `io.javalin:javalin:6.1.3` to `6.1.6`

#### Test Dependency Updates

* Updated `org.mockito:mockito-junit-jupiter:5.11.0` to `5.12.0`

#### Plugin Dependency Updates

* Updated `com.exasol:error-code-crawler-maven-plugin:2.0.2` to `2.0.3`
* Updated `org.apache.maven.plugins:maven-enforcer-plugin:3.4.1` to `3.5.0`
* Updated `org.apache.maven.plugins:maven-jar-plugin:3.3.0` to `3.4.1`
* Updated `org.apache.maven.plugins:maven-toolchains-plugin:3.1.0` to `3.2.0`
* Updated `org.sonarsource.scanner.maven:sonar-maven-plugin:3.11.0.3922` to `4.0.0.4121`

### Go-client

#### Compile Dependency Updates

* Updated `github.com/exasol/exasol-driver-go:v1.0.6` to `v1.0.7`
