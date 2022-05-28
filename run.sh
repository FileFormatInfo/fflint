#!/usr/bin/env bash
#
# run locally
#

set -o errexit
set -o pipefail
set -o nounset

rm -rf dist
goreleaser build --snapshot --single-target
#rm -rf badger
#ln -s dist/badger_darwin_amd64_v1/badger badger
export PATH=$PATH:$(pwd)/dist/badger_darwin_amd64_v1
badger version
go test -timeout 30s -run "^TestBadger$" github.com/fileformat/badger/cmd/badger