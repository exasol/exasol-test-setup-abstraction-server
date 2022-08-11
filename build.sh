#!/bin/bash
set -eo pipefail

readonly pk_version=2.6.1

base_dir="$( cd "$(dirname "$0")" >/dev/null 2>&1 ; pwd -P )"
readonly base_dir

cd "$base_dir"
mvn --batch-mode org.apache.maven.plugins:maven-dependency-plugin:3.1.1:get "-Dartifact=com.exasol:project-keeper-cli:$pk_version:jar"
java -jar "$HOME/.m2/repository/com/exasol/project-keeper-cli/$pk_version/project-keeper-cli-$pk_version.jar" verify

mvn --batch-mode --file server/ --batch-mode install

cd "$base_dir/go-client"
go test -v -coverprofile=coverage.out ./...
