#!/usr/bin/env bash
#
# build badger as a github action
#

set -o errexit
set -o pipefail
set -o nounset

SCRIPT_HOME="$( cd "$( dirname "$0" )" && pwd )"
APP_HOME=$(realpath "$SCRIPT_HOME/../..")

#COMMIT=$(git rev-parse --short HEAD)
COMMIT=gha
LASTMOD=$(date -u +%Y-%m-%dT%H:%M:%SZ)
VERSION=$(cat "${APP_HOME}/version.txt")

echo "PWD=$(pwd)"

go build \
    -a \
    -trimpath \
    -ldflags "-s -w -extldflags '-static' -X github.com/fileformat/badger/cmd.COMMIT=$COMMIT -X github.com/fileformat/badger/cmd.LASTMOD=$LASTMOD -X github.com/fileformat/badger/cmd.VERSION=$VERSION" \
    -installsuffix cgo \
    -tags netgo \
    -o "${SCRIPT_HOME}/badger-gha" \
    "${APP_HOME}"


ls -la "${APP_HOME}"
