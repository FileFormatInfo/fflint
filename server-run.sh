#!/usr/bin/env bash


set -o errexit
set -o pipefail
set -o nounset

if [ -f ".env" ]; then
	echo "INFO: loading .env file"
	export $(cat .env)
else
	echo "ERROR: .env file not found"
	exit 1
fi

~/go/bin/air
