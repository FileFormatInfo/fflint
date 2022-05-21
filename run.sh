#!/usr/bin/env bash
#
# run locally
#

set -o errexit
set -o pipefail
set -o nounset

go run cmd/badger.go ico ../*/favicon.ico ../*/docs/favicon.ico --showFiles --showPassing --debug
#go run cmd/badger.go ico ../*/favicon.ico ../*/docs/favicon.ico --showFiles --showDetail --showTests
#go run cmd/badger.go ico ../*/favicon.ico ../*/docs/favicon.ico --showFiles --showPassing --showDetail --showTests
#go run cmd/badger.go frontmatter  --required=title --output=filenames --progress=false "../vectorlogozone/**/*.md"
#go run cmd/badger.go svg "../vectorlogozone/www/**/*.svg" --namespace --namespaces=cc,dc,inkscape,rdf,sketch,sodipodi,svg,xlink --progress=false --output=filenames
#go run cmd/badger.go frontmatter "../vectorlogozone/www/logos/**/*.md" \
#    --required=logohandle,sort,title,website \
#    --strict \
#    --sorted \
#    --optional colors,font,images,keywords,tags \
#    --optional=blog,facebook,git,github,googleplus,guide,instagram,linkedin,pinterest,tumblr,twitter,wikipedia,youtube \
##    --optional=gitlab \
#    --optional=dribbble,flickr,gitter,reddit,slack,slideshare,snapchat,soundcloud,stackexchange,stackoverflow,vimeo,vine,weibo,xing \
#    --optional=noindex,other,redirect_from \
#    --showFiles --showDetail --showTests --debug
#    --showFiles
#go run cmd/badger.go svg "../vectorlogozone/**/*.svg" --showTests
#go run cmd/badger.go svg "../vectorlogozone/**/*.svg" --showFiles --glob=doublestar --debug
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
