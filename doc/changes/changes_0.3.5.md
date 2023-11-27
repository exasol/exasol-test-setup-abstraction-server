# Exasol Test Setup Abstraction Server 0.3.5, released 2023-11-27

Code name: Update default Exasol to 8.23.1

## Summary

This release updates the default Exasol DB version to 8.23.1.

## Features

* #42: Updated default Exasol DB version to 8.23.1

## Dependency Updates

### Server for the Exasol Test-Setup Abstraction

#### Compile Dependency Updates

* Updated `com.exasol:exasol-test-setup-abstraction-java:2.0.4` to `2.1.0`

#### Test Dependency Updates

* Updated `org.junit.jupiter:junit-jupiter-engine:5.10.0` to `5.10.1`
* Updated `org.junit.jupiter:junit-jupiter-params:5.10.0` to `5.10.1`
* Updated `org.mockito:mockito-junit-jupiter:5.6.0` to `5.7.0`

#### Plugin Dependency Updates

* Updated `com.exasol:error-code-crawler-maven-plugin:1.3.0` to `1.3.1`
* Updated `org.apache.maven.plugins:maven-enforcer-plugin:3.4.0` to `3.4.1`
* Updated `org.apache.maven.plugins:maven-failsafe-plugin:3.1.2` to `3.2.2`
* Updated `org.apache.maven.plugins:maven-surefire-plugin:3.1.2` to `3.2.2`
* Updated `org.codehaus.mojo:versions-maven-plugin:2.16.0` to `2.16.2`
* Updated `org.jacoco:jacoco-maven-plugin:0.8.10` to `0.8.11`
* Updated `org.sonarsource.scanner.maven:sonar-maven-plugin:3.9.1.2184` to `3.10.0.2594`

### Go-client

#### Compile Dependency Updates

* Updated `github.com/exasol/exasol-driver-go:v1.0.3` to `v1.0.4`
