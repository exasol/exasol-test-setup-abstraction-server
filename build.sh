#!/bin/bash
set -eo pipefail

base_dir="$( cd "$(dirname "$0")" >/dev/null 2>&1 ; pwd -P )"
readonly base_dir

cd "$base_dir"

mvn --batch-mode --file server/ clean install

cd "$base_dir/go-client"
go test -count 1 -p 1 -v -coverprofile=coverage.out ./...
