#!/usr/bin/env bash
#
# build a binary
#

set -o errexit
set -o pipefail
set -o nounset

rm -rf dist
goreleaser build --snapshot --single-target
