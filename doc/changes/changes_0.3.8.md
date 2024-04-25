# Exasol Test Setup Abstraction Server 0.3.8, released 2024-04-25

Code name: Fix SSH connection to Exasol Docker container 8.25.0 and later

## Summary

This release fixes error `JSchException: Algorithm negotiation fail` when starting an Exasol Docker container 8.25.0 and later. It also updates the default Exasol version to 8.26.0.

## Bugfixes

* #48: Fixed SSH connection to Exasol Docker container 8.25.0 and later

## Dependency Updates

### Server for the Exasol Test-Setup Abstraction

#### Compile Dependency Updates

* Updated `com.exasol:exasol-test-setup-abstraction-java:2.1.2` to `2.1.3`

#### Runtime Dependency Updates

* Updated `org.slf4j:slf4j-jdk14:2.0.12` to `2.0.13`
