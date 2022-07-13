# Exasol Test Setup Abstraction Server 0.2.2, released 2022-07-13

Code name: Upgrade dependencies

## Summary

This release upgrades dependencies to fix the following vulnerabilities:

* Server:
  * org.eclipse.jetty:jetty-server: CVE-2022-2047
  * org.eclipse.jetty:jetty-http: CVE-2022-2047
  * org.eclipse.jetty:jetty-client: CVE-2022-2047
* Go client:
  * github.com/containerd/containerd: CVE-2022-31030
  * github.com/opencontainers/runc: CVE-2022-29162

## Dependency Updates

### Server for the Exasol Test-Setup Abstraction

#### Compile Dependency Updates

* Updated `io.javalin:javalin:4.6.0` to `4.6.4`

#### Test Dependency Updates

* Updated `io.rest-assured:rest-assured:5.1.0` to `5.1.1`
* Updated `org.mockito:mockito-junit-jupiter:4.5.1` to `4.6.1`

### Go-client

#### Compile Dependency Updates

* Updated `github.com/exasol/exasol-driver-go:v0.4.2` to `v0.4.3`

#### Test Dependency Updates

* Updated `github.com/stretchr/testify:v1.7.1` to `v1.8.0`
