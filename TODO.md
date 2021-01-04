# To Do

## MVP

- [ ] svg command
- [ ] expand globs
- [ ] loop through files
- [ ] load each file
- [ ] tests for each file

## Formats

Each format will have a list of extensions and mime-types

- csv
- der
- gif
- html
- ini
- jpeg
- json
- pdf
- pem
- png
- svg
- toml
- tsv
- xml
- yaml
- zip

- Compressed formats (zip, gz, tgz, etc)
- MSOffice formats
- Music formats
- OpenOffice format
- Video formats

## General options

- charset:ascii|utf-8
- trailing-newline: on/off/any/only
- newline format: cr/crlf/lf/any (or dos/unix/mac?)
- indent: tab/spaces/any

## Format specific options

- html:noscript
- html:nocss
- html:tags=list,of,allowed,tags
- json:canonical
- json:schema (with optional url of schema)
- json:lines
- jpeg/png:metadata required/optional/none
- jpeg/png:colorprofile
- svg:width/height/viewBox
- svg:bitmap none/embedded/linked/any
- svg:font
- svg:meta
- svg:optimized
- pem/der:password:required/optional/none

## Output

- JSON-lines
- text
- yaml: similar to JSON, but grouped by file

levels:

- silent: just error code
- quiet (default): 1 line per failing file
- verbose: 1 line per failing test
- show-passing: include passing tests for quiet/verbose

other options:

- progress: on/off -  show progress bar (on automatically if output is tty)
- md5 - print md5 hash for each file

structured output:

- md5sum
- filename
- result: pass/fail
- if per-test output:
  - test: testid
  - result: pass/fail
  - detail: blob of detailed results (line #, etc)

## Distribution

- brew
- deb
- github action
- Docker container
- website

## To consider

- error if a file exists (to prevent certs/.env/source code)
- internal/alternate glob algorithms (or disable internal globbing)
- obey .gitignore when globbing
- newline: format `none` means no newlines (but handle trailing-newlines:on)
- check for file modes (i.e. executable, read-only, etc)
- slim binaries that only do a single format

## External Tools

- https://coptr.digipres.org/Category:Validation
- https://libguides.bodleian.ox.ac.uk/digitalpreservation/validation#:~:text=What%20is%20validation%3F,specific%20file%20format%20must%20follow.

* https://golang.org/pkg/path/filepath/

* https://github.com/JoshVarga/svgparser
* https://github.com/tealeg/xlsx
* https://golang.org/pkg/encoding/csv/
* https://golang.org/pkg/encoding/pem/
* https://golang.org/pkg/encoding/xml/
* https://golang.org/pkg/image/jpeg/
* https://pkg.go.dev/golang.org/x/image@v0.0.0-20201208152932-35266b937fa6/bmp