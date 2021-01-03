#!/usr/bin/env bash
#
# build a binary
#

set -o errexit
set -o pipefail
set -o nounset

mkdir -p dist

export COMMIT=local
export LASTMOD=$(date -u +%Y-%m-%dT%H:%M:%SZ)
export VERSION=$(cat version.txt)

go build \
    -ldflags "-X github.com/FileFormatInfo/badger/cmd.COMMIT=$COMMIT -X github.com/FileFormatInfo/badger/cmd.LASTMOD=$LASTMOD -X github.com/FileFormatInfo/badger/cmd.VERSION=$VERSION" \
    -o dist/badger \
    main.go

dist/badger version