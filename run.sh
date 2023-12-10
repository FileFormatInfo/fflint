#!/usr/bin/env bash
#
# run locally
#

set -o errexit
set -o pipefail
set -o nounset

rm -rf ./fflint
go build -o ./fflint cmd/fflint/main.go
export PATH=$PATH:$(pwd)
fflint version
go test -timeout 30s -run "^Testfflint$" github.com/FileFormatInfo/fflint/cmd/fflint