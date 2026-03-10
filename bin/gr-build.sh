#!/usr/bin/env bash
#
# build a binary
#

set -o errexit
set -o pipefail
set -o nounset

rm -rf dist
rm -f ./fflint
goreleaser build --snapshot --single-target --output=./fflint
