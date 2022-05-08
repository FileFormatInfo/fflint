#!/bin/bash
#
# flush the CloudFlare CDN with cURL.
#

set -o errexit
set -o pipefail
set -o nounset

SCRIPT_HOME="$( cd "$( dirname "$0" )" && pwd )"
REPO_HOME=$(realpath "${SCRIPT_HOME}/..")


echo "INFO: docgen starting at $(date -u +%Y-%m-%dT%H:%M:%SZ)"

cd "${REPO_HOME}"
go run cmd/docgen/main.go

echo "INFO: docgen complete at $(date -u +%Y-%m-%dT%H:%M:%SZ)"
