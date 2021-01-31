# To Do

## MVP

- [ ] https://github.com/mrusme/geld/blob/6d3f507ce5bd333930eb4a69dbf2638c2ee35cb0/.github/workflows/release.yaml
- [ ] use tags for versioning
- [ ] ratio type: decimal or x:y or decimal:decimal
- [ ] aspectRatio test: all image types
- [ ] gif: https://pkg.go.dev/image/gif
- [ ] txt
- [ ] webfonts: eot, woff, woff2, ttf
- [ ] pdf
- [ ] js: https://github.com/tdewolff/parse
- [ ] css: https://github.com/tdewolff/parse
- [ ] github actions working

## MVP 2

- [ ] new format: yaml
- [ ] html: strict flag (default=false)
- [ ] ico: sizes flag (array of ints)
- [ ] workflow: test
- [ ] all commands/types: tests
- [ ] handlers for [go encoding](https://golang.org/pkg/encoding/) formats
- [ ] Dockerfile to run locally, docker-run.sh
- [ ] workflow: release (on version.txt change)
- [ ] text handler: charset, newline, trailingNewLine
- [ ] glob: handle files via [stdin](https://www.socketloop.com/tutorials/golang-check-if-os-stdin-input-data-is-piped-or-from-terminal)
- [ ] nicer formatting of numbers in text output [x/text/message](https://pkg.go.dev/golang.org/x/text/message)

## MVP 3

- [ ] cache ReadFile results
- [ ] config file to support multiple sections, each with its own action and glob

## Documentation/Repo

- [ ] docs/flags/range.md
- [ ] docs/flags/ratio.md
- [ ] docs/flags/index.md
- [ ] docs/files.md - globbing
- [ ] docs/tests/index.md
- [ ] docs/commands/*.md
- [ ] docs/newformat.md - checklist for adding a new format
- [ ] docs/pricing.md
- [ ] [docgen command](https://github.com/spf13/cobra/blob/master/doc/md_docs.md)

## Formats

Each format will have a list of extensions and mime-types

- binary
- csv
- tsv
- txt

- config files (env, ini, toml)
- raster image formats (bmp, gif)
- crypto formats (crt, csr, der, key, p12, pem)
- compressed formats (bz2, gz, tar, tgz, zip)
- html scripting: handlebars/php/jsp
- markup (asciidoc, markdown, reStructuredText)
- MSOffice formats
- music formats
- OpenOffice format
- video formats

## Text tests (should work with other text formats)

- charset:ascii|utf-8
- trailing-newline: on/off/any/only
- newline format: cr/crlf/lf/any (or dos/unix/mac?)
- indent: tab/spaces/any
- contains/doesnotcontain: specific text (license declaration/etc)
- unicode: list of unicode character ranges allowed

## Format specific options

- html:noscript
- html:nocss
- html:tags=list,of,allowed,tags (or * or ones with html atoms)
- json:canonical
- json:schema (with optional url of schema)
- jpeg/png:metadata required/optional/none
- jpeg/png:colorprofile
- svg:bitmap none/embedded/linked/any
- svg:foreignObject
- svg:heightUnits true/false/list
- svg:widthUnits true/false/list
- svg:font
- svg:text
- svg:meta
- svg:optimized
- pem/der:password required/optional/none
- image formats:aspect ratio
- pdf (and others?):# of pages
- svg: viewBoxWidth check: range
- svg: viewBoxHeight check: range
- svg: [width|height]AllowUnits: [commalist|none|*|+], default=*

## Distribution

- github release
- brew [example](https://github.com/yudai/homebrew-gotty)
- deb
- github action
- Docker container
- website

## Cache

- only makes sense when multiple commands are run
- cache command (hidden)
- same globbing
- filesize range becomes cache load parameters
- max cache size parameter


## To consider

- [ ] new output format: [TAP](https://testanything.org/tap-version-13-specification.html), [awesome TAP](https://github.com/sindresorhus/awesome-tap)
- [ ] jsonlines: new format
- [ ] xhtml: new format
- [ ] html: support for charsets besides utf8
- [ ] html: alternate parser [tdewolff/parse](https://github.com/tdewolff/parse)
- [ ] progress: fix counter display during globbing
- [ ] progress: move terminal cursor to keep line in constant location if showFiles or showTests
- [ ] progress: if stderr is redirected to file, show stats every n seconds
- [ ] basic command: just the basic tests
- [ ] glob: alternative globber [mattn/go-zglob](https://github.com/mattn/go-zglob)
- [ ] glob: flag to specify an ignore file (i.e. .gitignore or .dockerignore)
- [ ] glob: --recursion flag and handle directories
- [ ] glob: case insensitive sort of files before processing
- [ ] filename test: regex/camel/kebab/pascal/snake/lowercase/urlsafe/none/etc (or should this be its own command?)
- [ ] find command: find files by extension or filetype
- [ ] mimetype: alterate library [h2non/filetype](https://github.com/h2non/filetype)
- [ ] mimetype: alterate library [gabriel-vasile/mimetype](https://github.com/gabriel-vasile/mimetype)
- [ ] png/jpeg/svg: meta check: none, contains
- [ ] output yaml: similar to JSON, but grouped by file, etc
- [ ] output md5: print md5 hash for each file
- `auto` mode which figures out the command based on the file extension (or contents?)
- shell completion
- generate documentation website
- [ ] directory/package re-org (internal, internal/formats, internal/commands, cmd/badger) [std](https://github.com/golang-standards/project-layout) - only when someone smarter than me can help

- newline: format `none` means no newlines (but handle trailing-newlines:on)
- check for file modes (i.e. executable, read-only, etc)
- support [`--version`](https://github.com/spf13/cobra#version-flag)
- [localization](https://pkg.go.dev/golang.org/x/text@v0.3.5/message)
- minimum file sizes

## Probably not

- [ ] showFiles: print at end if progress
- [ ] glob: alternate sorts (new first, largest first, etc)
- [ ] flag to enable colorized output
- [ ] progress: option to calc percentage by file count (vs bytes)
- error if a file exists (to prevent certs/.env/source code)
- slim binaries that only do a single format
- file locking
- each --verbose increments a --showXxx level
- glob: alternative globber: [godo](https://github.com/go-godo/godo/blob/master/glob.go)

## Links

https://www.client9.com/golang-globs-and-the-double-star-glob-operator/

## External Tools

- https://coptr.digipres.org/Category:Validation
- https://libguides.bodleian.ox.ac.uk/digitalpreservation/validation#:~:text=What%20is%20validation%3F,specific%20file%20format%20must%20follow.

## Go general libraries

- https://github.com/dustin/go-humanize
- https://golang.org/pkg/path/filepath/
- https://github.com/zabawaba99/go-gitignore
- https://github.com/danwakefield/fnmatch
- https://github.com/syncthing/syncthing/blob/v0.12.0-rc3/lib/fnmatch/fnmatch.go
- https://golang.org/pkg/net/http/#DetectContentType

## Go format libraries

* https://github.com/tealeg/xlsx
* https://pkg.go.dev/gopkg.in/yaml.v2
* https://github.com/dsoprea/go-exif or https://github.com/rwcarlsen/goexif
* https://pkg.go.dev/golang.org/x/image@v0.0.0-20201208152932-35266b937fa6/bmp
* https://github.com/xeipuuv/gojsonschema
* https://github.com/hashicorp/hcl
* https://github.com/360EntSecGroup-Skylar/excelize

## domains

- badger.sh
- badger-ci.com/badgerci.com
- badger.ci
