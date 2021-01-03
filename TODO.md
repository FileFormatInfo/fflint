# To Do

## MVP

- [ ] run.sh
- [ ] version command
- [ ] svg command
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
- trailing-newline: on/off/any
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
- standard: 1 line per file with PASS or FAIL
- verbose: 1 line per failing test
- show-passing: 1 line per test

other options:

- progress; show progress bar (on automatically if output is tty)
- md5: print md5 hash for each file

## Other features

- brew
- deb
- github action
- Docker container
- website

## To consider

- slim binary that only does a single format
- error if a file exists (to prevent certs/.env/source code)
- obey .gitignore
- newline: format `none` means no newlines (but handle trailing-newlines:on)

## External Tools

- https://coptr.digipres.org/Category:Validation
- https://libguides.bodleian.ox.ac.uk/digitalpreservation/validation#:~:text=What%20is%20validation%3F,specific%20file%20format%20must%20follow.
