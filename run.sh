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
go run main.go svg --min 5000 --max 20000 --size=:20000 --svgWidth=0: "/Users/andrew/Downloads/*.svg"
