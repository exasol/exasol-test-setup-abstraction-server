#!/bin/bash
set -eo pipefail
pk_version=2.4.4
mvn org.apache.maven.plugins:maven-dependency-plugin:3.1.1:get "-Dartifact=com.exasol:project-keeper-cli:$pk_version:jar"
java -jar "$HOME/.m2/repository/com/exasol/project-keeper-cli/$pk_version/project-keeper-cli-$pk_version.jar" verify
mvn --file server/ --batch-mode install

cd go-client
go test -v -coverprofile=coverage.out ./...
cd ..