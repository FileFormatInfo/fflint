#!/usr/bin/env bash
#
# run via Dockerfile
#

set -o errexit
set -o pipefail
set -o nounset


docker build \
	--build-arg COMMIT=$(git rev-parse --short HEAD) \
	--build-arg LASTMOD=$(date -u +%Y-%m-%dT%H:%M:%SZ) \
	--build-arg BUILTBY=docker-run \
    --build-arg VERSION=local \
	-t fflint-online .

docker run -it -p 4000:4000 fflint-online
