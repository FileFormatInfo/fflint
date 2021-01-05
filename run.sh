#!/usr/bin/env bash
#
# run locally
#

set -o errexit
set -o pipefail
set -o nounset


export LASTMOD=$(date -u)
export COMMIT=local

#go run main.go --max 20000 --min 5000 svg "../vectorlogozone/**/*.svg"
go run main.go svg "/Users/andrew/Downloads/*.svg"
