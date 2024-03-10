#!/usr/bin/env bash
#
# run tests locally
#

set -o errexit
set -o pipefail
set -o nounset

if [ -f "./fflint" ]; then
  echo "INFO: removing old build of fflint"
  rm ./fflint
fi

echo "INFO: building new fflint"
go build -o ./fflint cmd/fflint/main.go
if [ ! -f "./fflint" ]; then
  echo "ERROR: failed to build fflint"
  exit 1
fi


export PATH=$(pwd):$PATH
echo "INFO: running fflint version $(fflint version)"

echo "INFO: running tests"
go test -timeout 30s -run "^TestFflint$" github.com/FileFormatInfo/fflint/cmd/fflint

echo "INFO: complete at $(date -u +%Y-%m-%dT%H:%M:%SZ)"