#!/usr/bin/env bash
#
# run locally
#

set -o errexit
set -o pipefail
set -o nounset

go run cmd/badger.go svg "../vectorlogozone/**/*.svg" --showTests
#go run cmd/badger.go svg "../vectorlogozone/**/*.svg" --showFiles
#go run cmd/badger.go jpeg --width=700-2000 "../peter/**/*.jpeg" --showTests
#go run cmd/badger.go \
#    ext \
#    --showTests \
#    --allowed=css,eot,ico,jpeg,js,json,html,md,pdf,png,svg,ttf,txt,woff,woff2,xml,yaml,yml \
#    --report \
#    "../vectorlogozone/www/**"

#go run cmd/badger.go svg \
#    --width=120 \
#    --height=60 \
#    --showDetail \
#    --showTests \
#    --showFiles \
#    --output=json \
#	"../vectorlogozone/www/logos/**/*-ar21.svg"

#go run cmd/badger.go svg \
#    --width=64 \
#    --height=64 \
#    --showDetail \
#    --showTests \
#    --showFiles \
#    --output=json \
#	"../vectorlogozone/www/logos/**/*-icon.svg"
