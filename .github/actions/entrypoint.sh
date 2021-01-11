#!/usr/bin/env bash
#
# run badger as a github action
#

set -o errexit
set -o pipefail
set -o nounset

SCRIPT_HOME="$( cd "$( dirname "$0" )" && pwd )"

# NOTE: deliberately no quotes around options, since they are not specified individually
"${SCRIPT_HOME}/badger-gha" ${INPUT_CMD:-version} ${INPUT_OPTIONS:-} "${INPUT_FILES:-}"
