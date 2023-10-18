# Exasol Test Setup Abstraction Server 0.3.4, released 2023-10-19

Code name: Configure startup timeout

## Summary

This release allows configuring the server startup timeout for the Go client. This is useful when the default of 10 minutes is not long enough, e.g. on slow GitHub runners.

The release also upgrades the default Exasol version to 7.1.23.

**Note**: This release contains the following third party libraries with vulnerabilities:
* `io.netty:netty-handler`: CVE-2023-4586 CWE-300: Channel Accessible by Non-Endpoint ('Man-in-the-Middle') (6.5);
* `fr.turri:aXMLRPC`: CVE-2020-36641 CWE-611: Improper Restriction of XML External Entity Reference ('XXE') (9.8);

We assume that these vulnerabilities are not exploitable.

## Features

* #41: Added configuration for server startup timeout

## Dependency Updates

### Server for the Exasol Test-Setup Abstraction

#### Compile Dependency Updates

* Updated `com.exasol:exasol-test-setup-abstraction-java:2.0.2` to `2.0.4`
* Updated `io.javalin:javalin:5.6.1` to `5.6.3`

#### Runtime Dependency Updates

* Updated `org.slf4j:slf4j-jdk14:2.0.7` to `2.0.9`

#### Test Dependency Updates

* Updated `io.rest-assured:rest-assured:5.3.1` to `5.3.2`
* Updated `org.junit.jupiter:junit-jupiter-engine:5.9.3` to `5.10.0`
* Updated `org.junit.jupiter:junit-jupiter-params:5.9.3` to `5.10.0`
* Updated `org.mockito:mockito-junit-jupiter:5.4.0` to `5.6.0`

#### Plugin Dependency Updates

* Updated `com.exasol:error-code-crawler-maven-plugin:1.2.3` to `1.3.0`
* Updated `org.apache.maven.plugins:maven-assembly-plugin:3.5.0` to `3.6.0`
* Updated `org.apache.maven.plugins:maven-enforcer-plugin:3.3.0` to `3.4.0`
* Updated `org.apache.maven.plugins:maven-failsafe-plugin:3.0.0` to `3.1.2`
* Updated `org.apache.maven.plugins:maven-surefire-plugin:3.0.0` to `3.1.2`
* Updated `org.basepom.maven:duplicate-finder-maven-plugin:1.5.1` to `2.0.1`
* Updated `org.codehaus.mojo:flatten-maven-plugin:1.4.1` to `1.5.0`
* Updated `org.codehaus.mojo:versions-maven-plugin:2.15.0` to `2.16.0`
* Updated `org.jacoco:jacoco-maven-plugin:0.8.9` to `0.8.10`

### Go-client

#### Compile Dependency Updates

* Updated `golang:1.19` to `1.20`
* Updated `github.com/exasol/exasol-driver-go:v1.0.0` to `v1.0.3`

#### Test Dependency Updates

* Updated `github.com/antchfx/xmlquery:v1.3.17` to `v1.3.18`
