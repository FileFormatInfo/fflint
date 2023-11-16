# To Do

- [ ] doc: [svg role and title](https://www.smashingmagazine.com/2021/05/accessible-svg-patterns-comparison/)
- [ ] doc and README: share links
- [ ] workflow: run fflint on the fflint.org website
- [ ] workflow: go test
- [ ] test without ./main.go
- [ ] shared/ImageFlags.go
- [ ] credits: goatcounter, sass, nodeping, bootstrap
- [ ] doc strings in Cobra: Long, Examples
- [ ] doc: installation
- [ ] doc: pricing/license
- [ ] doc: single font (i.e. header font)
- [ ] doc: responsive navbar
- [ ] doc: upgrade bootstrap
- [ ] frontmatter: report
- [ ] til frontmatter: title,noindex,layout,draft,date,tags,created
- [ ] github actions working
- [ ] OptionalBool flag type: required/any/forbidden
- [ ] Dimensions flag (ranges): WxH || N (=square)
- [ ] OptionalXxx flags
- [ ] SignedRange (i.e. handle negative numbers: -10-10, -20--10, --10, -10-)
- [ ] Range: take a comma-separated list of acceptable values

## Online

- [ ] logging (not per-page, since that is done by CloudRun)
- [ ] disclaimer in all responses
- [ ] verbose flag
- [ ] pass flags to fflint
- [ ] forms (in docs)
- [ ] api.html (in docs)
- [ ] links to forms from /index.html

## Docs

- [ ] better description on home page
- [ ] local contact page
- [ ] command `Long` should include links to file format specifications
- [ ] examples
- [ ] installation
- [ ] quick start
- [ ] using in pre-commit hook, CI pipeline(s), etc.
- [ ] examples: everything in an online sitemap.xml
- [ ] add link to [ignorefile specs](https://git-scm.com/docs/gitignore)

## General work

- [ ] make `go test` part of build.yaml
- [ ] all commands/types: tests
- [ ] Dockerfile to run locally, docker-run.sh
- [ ] nicer formatting of numbers in text output [x/text/message](https://pkg.go.dev/golang.org/x/text/message) or [go-humanize](https://github.com/dustin/go-humanize)
- [ ] include a man page: [mango](https://github.com/muesli/mango)
- [ ] shell completion
- [ ] cache ReadFile results (--nocache=bytes to control memory usage)

## GoLang

- [ ] multithreading
- [ ] why .go file in repo root directory
- [ ] generics for [Decimal|Integer]Range, [Decimal|Integer]Ratio, Optional[*], [*]List
- [ ] generics for contains(haystack, needle)
- [ ] how to share code for image tests
- [ ] ability to use plugins? [general discussion](https://eli.thegreenplace.net/2021/plugins-in-go/), [go-plugin](https://github.com/hashicorp/go-plugin), [pie](https://github.com/natefinch/pie), [wazero](https://github.com/tetratelabs/wazero)

## Additional formats

- [ ] [yaml](https://pkg.go.dev/gopkg.in/yaml.v2)
- [ ] RSS and atom feeds [best practices](https://kevincox.ca/2022/05/06/rss-feed-best-practices/)
- [ ] sitemap.xml
- [ ] jsonlines (should share a lot/all of code with json)
- [ ] MS Excel [xlsx](https://github.com/tealeg/xlsx) [excelize](https://github.com/360EntSecGroup-Skylar/excelize)
- [ ] [bmp](https://pkg.go.dev/golang.org/x/image@v0.0.0-20201208152932-35266b937fa6/bmp)
- [ ] [hcl](https://github.com/hashicorp/hcl)
- [ ] [toml](github.com/BurntSushi/toml)
- [ ] [webp](https://github.com/kolesa-team/go-webp)
- [ ] [shell scripts](https://github.com/mvdan/sh) ([existing gha](https://github.com/luizm/action-sh-checker))
- [ ] xhtml
- [ ] OPML
- [ ] [gif](https://pkg.go.dev/image/gif)
- [ ] webfonts: eot, woff, woff2, ttf
- [ ] pdf (# of pages, page size, page orientation)
- [ ] js: https://github.com/tdewolff/parse
- [ ] css: https://github.com/tdewolff/parse
- [ ] [GoLang encoding](https://golang.org/pkg/encoding/) formats
- [ ] [TAP](https://testanything.org/tap-version-13-specification.html), [awesome TAP](https://github.com/sindresorhus/awesome-tap)

- compressed formats (bz2, gz, tar, tgz, zip)
- config files (env, ini)
- crypto formats (crt, csr, der, key, p12, pem)
- data dumps (csv, tsv, Xsv)
- html scripting: handlebars/php/jsp
- markup (asciidoc, markdown, reStructuredText)
- MSOffice formats
- music formats
- OpenOffice format
- video formats
- SQLite

## Mixed/Auto mode

- single pass through file list
- different tests based on filename
- config file with multiple sections
- each section: name, filename_regex, command, flags
- special 'skip' command for files that should not be checked
- default case: nothing or error or warn or ???
- profiles for different scenarios: website, backend codebase...

## To consider

- [ ] progress: fix counter display during globbing
- [ ] progress: move terminal cursor to keep line in constant location if showFiles or showTests
- [ ] progress: if stderr is redirected to file, show stats every n seconds
- [ ] output yaml: similar to JSON, but grouped by file, etc
- [ ] output md5: print md5 hash for each file (maybe as part of debug mode?)
- [ ] web mode: scan files on a website instead of a file system (or just example with `wget`)

## Globbing improvements

- [ ] glob: alternative globber [mattn/go-zglob](https://github.com/mattn/go-zglob)
- [ ] glob: --recursion flag and handle directories
- [ ] glob: case insensitive sort of files before processing
- [ ] support for filenames (on stdin) separated by 0x00 byte
- [ ] glob: alternate sorts (date created/modified, size)

- newline: format `none` means no newlines (but handle trailing-newlines:on)
- check for file modes (i.e. executable, read-only, etc)
- support [`--version`](https://github.com/spf13/cobra#version-flag)
- [localization](https://pkg.go.dev/golang.org/x/text@v0.3.5/message)
- minimum file sizes
- https://github.com/zabawaba99/go-gitignore
- NO: alternative globber: [godo](https://github.com/go-godo/godo/blob/master/glob.go)
- https://www.client9.com/golang-globs-and-the-double-star-glob-operator/
- https://golang.org/pkg/path/filepath/
- https://github.com/danwakefield/fnmatch
- https://github.com/syncthing/syncthing/blob/v0.12.0-rc3/lib/fnmatch/fnmatch.go

## Distribution

- brew [example](https://github.com/yudai/homebrew-gotty)
- deb
- rpm
- github action
- Docker container
- website

## Probably not

- [ ] showFiles: print at end if progress
- [ ] html: support for charsets besides utf8
- [ ] flag to enable colorized output
- [ ] progress: option to calc percentage by file count (vs bytes)
- error if a file exists (to prevent certs/.env/source code)
- slim binaries that only do a single format
- file locking
- each --verbose increments a --showXxx level
- GUI version

## External Tools

- https://coptr.digipres.org/Category:Validation
- https://libguides.bodleian.ox.ac.uk/digitalpreservation/validation#:~:text=What%20is%20validation%3F,specific%20file%20format%20must%20follow.

## Bad file names

- separate program in Rust (namelint)
- takes directory names only (possibly multiple, default to current directory)
- recursively (optional) scans the directory looking for bad file names
- option: specify list of codepoints allowed
- profiles: preset lists of codepoints
- profile: strict (ASCII letters, numbers, dot)
- profile: unicode (UTF-8, no shell chars, newlines, lookalikes)
- profile: loose (UTF-8)
- maybe: regex/camel/kebab/pascal/snake/lowercase/urlsafe/none/etc
- maybe: profile for a given language
- maybe: note hidden files, leading dot files, bad modes, etc
- separate utility program to list all codepoints used (with first few files containing each)

https://github.com/jhspetersson/fselect
https://github.com/github/super-linter#supported-linters
https://htmlhint.com/docs/user-guide/list-rules
https://stackoverflow.com/questions/29838185/how-to-detect-additional-mime-type-in-golang
text/tabwriter