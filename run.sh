#!/usr/bin/env bash
#
# run locally
#

set -o errexit
set -o pipefail
set -o nounset

rm -rf ./badger
go build -o ./badger cmd/badger/main.go
export PATH=$PATH:$(pwd)
badger version
go test -timeout 30s -run "^TestBadger$" github.com/fileformat/badger/cmd/badger