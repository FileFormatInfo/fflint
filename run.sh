#!/usr/bin/env bash
#
# run locally
#

set -o errexit
set -o pipefail
set -o nounset

#go run cmd/badger.go svg "../vectorlogozone/**/*.svg" --showFiles

#go run cmd/badger.go svg \
#    --width=120 \
#    --height=60 \
#    --showDetail \
#    --showTests \
#    --showFiles \
#    --output=json \
#	"../vectorlogozone/www/logos/**/*-ar21.svg"

go run cmd/badger.go svg \
    --width=64 \
    --height=64 \
    --showDetail \
    --showTests \
    --showFiles \
    --output=json \
	"../vectorlogozone/www/logos/**/*-icon.svg"
