# Exasol Test Setup Abstraction Server 0.3.6, released 2024-03-12

Code name: Fix CVE-2024-25710 and CVE-2024-26308 in compile dependency

## Summary

This release fixes vulnerabilities CVE-2024-25710 and CVE-2024-26308 in compile dependency `org.apache.commons:commons-compress`.

**Excluded Vulnerability** We accept vulnerability CVE-2017-10355 (CWE-833: Deadlock) in transitive compile dependency `xerces:xercesImpl:jar:2.12.2` as we assume that we only connect to the known endpoint ExaOperations.

## Security

* #44: Fixed vulnerabilities CVE-2024-25710 and CVE-2024-26308 in compile dependency `org.apache.commons:commons-compress`

## Dependency Updates

### Server for the Exasol Test-Setup Abstraction

#### Compile Dependency Updates

* Updated `com.exasol:exasol-test-setup-abstraction-java:2.1.0` to `2.1.1`
* Updated `io.javalin:javalin:5.6.3` to `6.1.3`

#### Runtime Dependency Updates

* Updated `org.slf4j:slf4j-jdk14:2.0.9` to `2.0.12`

#### Test Dependency Updates

* Updated `io.rest-assured:rest-assured:5.3.2` to `5.4.0`
* Updated `org.junit.jupiter:junit-jupiter-engine:5.10.1` to `5.10.2`
* Updated `org.junit.jupiter:junit-jupiter-params:5.10.1` to `5.10.2`
* Updated `org.mockito:mockito-junit-jupiter:5.7.0` to `5.11.0`

#### Plugin Dependency Updates

* Updated `com.exasol:error-code-crawler-maven-plugin:1.3.1` to `2.0.0`
* Updated `org.apache.maven.plugins:maven-compiler-plugin:3.11.0` to `3.12.1`
* Updated `org.apache.maven.plugins:maven-failsafe-plugin:3.2.2` to `3.2.5`
* Updated `org.apache.maven.plugins:maven-surefire-plugin:3.2.2` to `3.2.5`
* Added `org.apache.maven.plugins:maven-toolchains-plugin:3.1.0`
* Updated `org.codehaus.mojo:flatten-maven-plugin:1.5.0` to `1.6.0`

### Go-client

#### Compile Dependency Updates

* Updated `golang:1.20` to `1.21`
* Updated `github.com/stretchr/testify:v1.8.4` to `v1.9.0`
