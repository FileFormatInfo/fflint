#!/usr/bin/env bash
#
# run badger as a github action
#

set -o errexit
set -o pipefail
set -o nounset

# NOTE: deliberately no quotes around options, since they are not specified individually
/bin/badger-gha ${INPUT_CMD} ${INPUT_OPTIONS:-} "${INPUT_FILES}"
