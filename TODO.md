# To Do

## MVP

- [ ] png cmd
- [ ] jpeg cmd
- [ ] decent output
- [ ] meta check: none, contains
- [ ] viewBox check: range,range,range,range
- [ ] viewBoxWidth check: range
- [ ] viewBoxHeight check: range
- [ ] range: decimal version
- [ ] workflows
- [ ] charset check

## Documentation/Repo

- [ ] CONTRIBUTING.md
- [ ] .github/ISSUE_TEMPLATE - new format
- [ ] .github/ISSUE_TEMPLATE - bug
- [ ] .github/ISSUE_TEMPLATE - new feature or test
- [ ] .github/ISSUE_TEMPLATE/config.yaml
- [ ] .github/ISSUE_TEMPLATE - new format
- [ ] .github/PULL_REQUEST_TEMPLATE/newpr.md
- [ ] docs/range.md
- [ ] docs/glob.md


## Formats

Each format will have a list of extensions and mime-types

- binary
- csv
- html
- json
- pdf
- svg
- tsv
- txt
- xml
- yaml

- config files (env, ini, toml)
- raster image formats (bmp, gif, ico, jpeg, png)
- crypto formats (crt, csr, der, key, p12, pem)
- compressed formats (bz2, gz, tar, tgz, zip)
- html scripting: handlebars/php/jsp
- markup (asciidoc, markdown, reStructuredText)
- MSOffice formats
- music formats
- OpenOffice format
- video formats

## General options

- charset:ascii|utf-8
- trailing-newline: on/off/any/only
- newline format: cr/crlf/lf/any (or dos/unix/mac?)
- indent: tab/spaces/any
- contains/doesnotcontain: specific text (license declaration/etc)
- filename rules (snake/camel/lowercase/etc)

## Format specific options

- html:noscript
- html:nocss
- html:tags=list,of,allowed,tags
- json:canonical
- json:schema (with optional url of schema)
- json:lines
- jpeg/png:metadata required/optional/none
- jpeg/png:colorprofile
- svg:viewBox
- svg:bitmap none/embedded/linked/any
- svg:foreignObject
- svg:heightUnits true/false/list
- svg:widthUnits true/false/list
- svg:font
- svg:text
- svg:meta
- svg:optimized
- pem/der:password:required/optional/none
- image formats:aspect ratio
- pdf (and others?):# of pages


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

- `auto` mode which figures out the command based on the file extension (or contents?)
- shell completion
- generate documentation website
- internal/alternate glob algorithms (or disable internal globbing)
- obey .gitignore when globbing
- newline: format `none` means no newlines (but handle trailing-newlines:on)
- check for file modes (i.e. executable, read-only, etc)

## Probably not
- error if a file exists (to prevent certs/.env/source code)
- slim binaries that only do a single format

## External Tools

- https://coptr.digipres.org/Category:Validation
- https://libguides.bodleian.ox.ac.uk/digitalpreservation/validation#:~:text=What%20is%20validation%3F,specific%20file%20format%20must%20follow.

## Go general libraries

- https://golang.org/pkg/path/filepath/
- https://github.com/zabawaba99/go-gitignore
- https://github.com/danwakefield/fnmatch
- https://github.com/syncthing/syncthing/blob/v0.12.0-rc3/lib/fnmatch/fnmatch.go

## Go format libraries

* https://github.com/JoshVarga/svgparser
* https://github.com/tealeg/xlsx
* https://golang.org/pkg/encoding/csv/
* https://golang.org/pkg/encoding/pem/
* https://golang.org/pkg/encoding/xml/
* https://golang.org/pkg/image/jpeg/
* https://pkg.go.dev/golang.org/x/image@v0.0.0-20201208152932-35266b937fa6/bmp
