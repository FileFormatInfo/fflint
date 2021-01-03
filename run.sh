#!/usr/bin/env bash
#
# run locally
#

set -o errexit
set -o pipefail
set -o nounset


export LASTMOD=$(date -u)
export COMMIT=local

go run main.go version
