# Server for the Exasol Test-Setup Abstraction 0.2.0, released 2022-??-??

Code name: Fixed failing startup when docker is not running on localhost

## Features

* #7: Support for creating a connection with no autocommit
* #9: Added API for accessing BucketFS

## Bug Fixes

* #5: Fixed failing startup when docker is not running on localhost

## Dependency Updates

### Compile Dependency Updates

* Updated `com.exasol:exasol-test-setup-abstraction-java:0.3.0` to `0.3.2`
* Updated `io.javalin:javalin:4.5.0` to `4.6.0`
* Added `software.amazon.awssdk:cloudformation:2.17.198`

### Test Dependency Updates

* Added `commons-codec:commons-codec:1.15`
* Updated `io.rest-assured:rest-assured:5.0.1` to `5.1.0`
