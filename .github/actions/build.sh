#!/usr/bin/env bash
#
# build fflint as a github action
#

set -o errexit
set -o pipefail
set -o nounset

SCRIPT_HOME="$( cd "$( dirname "$0" )" && pwd )"
APP_HOME=$(realpath "$SCRIPT_HOME/../..")

LASTMOD=$(date -u +%Y-%m-%dT%H:%M:%SZ)
VERSION=$(cat "${APP_HOME}/version.txt")

echo "PWD=$(pwd)"
echo "APP_HOME=${APP_HOME}"
echo "SCRIPT_HOME=${SCRIPT_HOME}"
echo "GITHUB_SHA=${GITHUB_SHA:-(not set)}"

echo "INFO: creating directory"
mkdir -p "${SCRIPT_HOME}/dist"

echo "INFO: building"
go build \
    -a \
    -trimpath \
    -ldflags "-s -w -extldflags '-static' -X main.COMMIT=$GITHUB_SHA -X main.LASTMOD=$LASTMOD -X main.VERSION=$VERSION -X main.BUILTBY=build.sh" \
    -installsuffix cgo \
    -tags netgo \
    -o "${APP_HOME}/dist/fflint" \
    "${APP_HOME}"

echo "INFO: running test"
"${APP_HOME}/dist/fflint" version
