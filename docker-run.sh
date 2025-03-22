#!/usr/bin/env bash
#
# run via Dockerfile
#

set -o errexit
set -o pipefail
set -o nounset


APP_NAME="fflint-server"

ENVFILE=${1:-.env}

if [ ! -r "${ENVFILE}" ]
then
    echo "ERROR: no .env file '${ENVFILE}'!"
    exit 1
fi

echo "INFO: building docker image..."
docker build \
    --build-arg COMMIT=local@$(git rev-parse --short HEAD) \
    --build-arg LASTMOD=$(date -u +%Y-%m-%dT%H:%M:%SZ) \
    --progress=plain \
    --tag "${APP_NAME}" \
    .

docker run \
    --env-file "./.env" \
	--env ADDRESS=0.0.0.0 \
    --interactive \
    --name "${APP_NAME}" \
    --publish "4000:4000" \
    --rm \
    --tty \
    "${APP_NAME}"

