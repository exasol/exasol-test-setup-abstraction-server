# Exasol Test-Setup Abstraction GO

[![Build Status](https://github.com/exasol/exasol-test-setup-abstraction-server/actions/workflows/ci-build.yml/badge.svg)](https://github.com/exasol/exasol-test-setup-abstraction-server/actions/workflows/ci-build.yml)

[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=com.exasol%3Aexasol-test-setup-abstraction-server&metric=alert_status)](https://sonarcloud.io/dashboard?id=com.exasol%3Aexasol-test-setup-abstraction-server)

[![Security Rating](https://sonarcloud.io/api/project_badges/measure?project=com.exasol%3Aexasol-test-setup-abstraction-server&metric=security_rating)](https://sonarcloud.io/dashboard?id=com.exasol%3Aexasol-test-setup-abstraction-server)
[![Reliability Rating](https://sonarcloud.io/api/project_badges/measure?project=com.exasol%3Aexasol-test-setup-abstraction-server&metric=reliability_rating)](https://sonarcloud.io/dashboard?id=com.exasol%3Aexasol-test-setup-abstraction-server)
[![Maintainability Rating](https://sonarcloud.io/api/project_badges/measure?project=com.exasol%3Aexasol-test-setup-abstraction-server&metric=sqale_rating)](https://sonarcloud.io/dashboard?id=com.exasol%3Aexasol-test-setup-abstraction-server)
[![Technical Debt](https://sonarcloud.io/api/project_badges/measure?project=com.exasol%3Aexasol-test-setup-abstraction-server&metric=sqale_index)](https://sonarcloud.io/dashboard?id=com.exasol%3Aexasol-test-setup-abstraction-server)

[![Code Smells](https://sonarcloud.io/api/project_badges/measure?project=com.exasol%3Aexasol-test-setup-abstraction-server&metric=code_smells)](https://sonarcloud.io/dashboard?id=com.exasol%3Aexasol-test-setup-abstraction-server)
[![Coverage](https://sonarcloud.io/api/project_badges/measure?project=com.exasol%3Aexasol-test-setup-abstraction-server&metric=coverage)](https://sonarcloud.io/dashboard?id=com.exasol%3Aexasol-test-setup-abstraction-server)
[![Duplicated Lines (%)](https://sonarcloud.io/api/project_badges/measure?project=com.exasol%3Aexasol-test-setup-abstraction-server&metric=duplicated_lines_density)](https://sonarcloud.io/dashboard?id=com.exasol%3Aexasol-test-setup-abstraction-server)
[![Lines of Code](https://sonarcloud.io/api/project_badges/measure?project=com.exasol%3Aexasol-test-setup-abstraction-server&metric=ncloc)](https://sonarcloud.io/dashboard?id=com.exasol%3Aexasol-test-setup-abstraction-server)

[![Go Reference](https://pkg.go.dev/badge/github.com/exasol/exasol-test-setup-abstraction-server/go-client.svg)](https://pkg.go.dev/github.com/exasol/exasol-test-setup-abstraction-server/go-client)

This project makes the features
of [exasol-test-setup-abstraction-java](https://github.com/exasol/exasol-test-setup-abstraction-java/) available in the
following languages:

* Go

## Usage

### Go

```go
package test

import testSetupAbstraction "github.com/exasol/exasol-test-setup-abstraction-server/go-client"

func myTest() {
	exasol, err := testSetupAbstraction.New().CloudSetupConfigFilePath("myConfig.json").DockerDbVersion("7.1.15").Start()
	if err != nil {
		panic("failed to start test setup")
	}
	connection := exasol.CreateConnection()
	//...
}
```

## Additional Information

* [Changelog](doc/changes/changelog.md)
* [Dependencies](dependencies.md)
* [Developers Guide](doc/developers_guide/developers_guide.md)
