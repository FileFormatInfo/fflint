# Badger [<img alt="badger logo" src="docs/favicon.svg" height="90" align="right" />](https://www.badger.sh/)

[![build](https://github.com/fileformat/badger/actions/workflows/build.yaml/badge.svg)](https://github.com/fileformat/badger/actions/workflows/build.yaml)
[![release](https://github.com/fileformat/badger/actions/workflows/release.yaml/badge.svg)](https://github.com/fileformat/badger/actions/workflows/release.yaml)
[![dogfooding](https://github.com/fileformat/badger/actions/workflows/dogfooding.yaml/badge.svg)](https://github.com/fileformat/badger/actions/workflows/dogfooding.yaml)

Badger is a linter for file formats. Are your files:
* in the correct format?
* with the correct extension?
* with the correct image dimensions?
* properly stripped of revealing metadata?
* not too big or too small?
* have decent names?

Perfect for your CI/CD pipeline to make sure bad files don't get committed.

[**Documentation**](https://www.badger.sh)

## Installation

The [latest releases](https://github.com/fileformat/badger/releases/latest) are available on Github. [Detailed instructions](https://www.badger.sh/install.html).

## Usage

General command syntax is:

```bash
badger CMD [options...] files...
```

* `CMD` is the command to run
* `options...` are command-specific options
* `files...` are the files to check

More:
* Complete documentation is on [**www.badger.sh**](https://www.badger.sh)
* Run `badger help` to see a list of available commands
* Run `--help` for any command to see options specific to that command.  Example: `badger svg --help`

## License

[GNU Affero General Public License v3.0](LICENSE.txt)

For anyone who cannot use AGPL software, an inexpensive commercial license is available.<!-- LATER: link to pricing page on website -->

## Credits

[![bash](https://www.vectorlogo.zone/logos/gnu_bash/gnu_bash-ar21.svg)](https://www.gnu.org/software/bash/ "Scripting")
[![Git](https://www.vectorlogo.zone/logos/git-scm/git-scm-ar21.svg)](https://git-scm.com/ "Version control")
[![Github](https://www.vectorlogo.zone/logos/github/github-ar21.svg)](https://github.com/ "Code hosting")
[![golang](https://www.vectorlogo.zone/logos/golang/golang-ar21.svg)](https://golang.org/ "Programming language")
[![Google Noto Emoji](https://www.vectorlogo.zone/logos/google/google-ar21.svg)](https://github.com/googlefonts/noto-emoji/blob/5628587386c78161f87aa2ca9ddee37c2e8ea212/svg/emoji_u1f9a1.svg "Logo")
[![Jekyll](https://www.vectorlogo.zone/logos/jekyllrb/jekyllrb-ar21.svg)](https://www.jekyllrb.com/ "Website")

* [GoReleaser](https://goreleaser.com/)
* [mathiasbynens/small](https://github.com/mathiasbynens/small) - sample files for testing

* See [`go.mod`](https://github.com/fileformat/badger/blob/main/go.mod) for the GoLang modules used
