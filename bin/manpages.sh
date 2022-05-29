#!/usr/bin/env bash
#
# generate completions for the various shells
#

set -o errexit
set -o pipefail
set -o nounset

SCRIPT_HOME="$( cd "$( dirname "$0" )" && pwd )"
REPO_HOME=$(realpath "${SCRIPT_HOME}/..")
MANPAGES_DIR="${REPO_HOME}/manpages"

echo "INFO: manpages starting at $(date -u +%Y-%m-%dT%H:%M:%SZ)"

rm -rf "${MANPAGES_DIR}"
mkdir "${MANPAGES_DIR}"

echo "INFO: generating manpages"
go run "${REPO_HOME}/cmd/mangen" man
#| gzip -c -9 >"${MANPAGES_DIR}/badger.1.gz"

echo "INFO: manpages complete at $(date -u +%Y-%m-%dT%H:%M:%SZ)"
