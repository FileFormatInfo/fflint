# fflint [<img alt="fflint logo" src="docs/favicon.svg" height="90" align="right" />](https://www.fflint.org/)

[![build](https://github.com/FileFormatInfo/fflint/actions/workflows/build.yaml/badge.svg)](https://github.com/FileFormatInfo/fflint/actions/workflows/build.yaml)
[![release](https://github.com/FileFormatInfo/fflintactions/workflows/release.yaml/badge.svg)](https://github.com/FileFormatInfo/fflint/actions/workflows/release.yaml)
[![dogfooding](https://github.com/FileFormatInfo/fflint/actions/workflows/dogfooding.yaml/badge.svg)](https://github.com/FileFormatInfo/fflint/actions/workflows/dogfooding.yaml)

fflint is a linter for file formats. Are your files:
* in the correct format?
* with the correct extension?
* with the correct image dimensions?
* properly stripped of revealing metadata?
* not too big or too small?
* have decent names?

Perfect for your CI/CD pipeline to make sure bad files don't get committed.

[**Documentation**](https://www.fflint.org/)

## Installation

The [latest releases](https://github.com/FileFormatInfo/fflint/releases/latest) are available on Github. [Detailed instructions](https://www.fflint.org/install.html).

## Usage

General command syntax is:

```bash
fflint CMD [options...] files...
```

* `CMD` is the command to run
* `options...` are command-specific options
* `files...` are the files to check

More:
* Complete documentation is on [**www.fflint.org**](https://www.fflint.org)
* Run `fflint help` to see a list of available commands
* Run `--help` for any command to see options specific to that command.  Example: `fflint svg --help`

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
* [svgrepo](https:svgrepo.com) - [flint hand axe icon](https://www.svgrepo.com/svg/156483/hand-axe)
* IconFinder alternate icon: [flint hammer](https://www.iconfinder.com/icons/3286710/caveman_flint_hammer_prehistoric_stone_tribal_weapon_icon) (purchased)
* IconFinder alternate icon: [flint axe](https://www.iconfinder.com/icons/11293805/stone_weapon_age_ancient_ax_icon) (purchased)
* See [`go.mod`](https://github.com/FileFormatInfo/fflint/blob/main/go.mod) for the GoLang modules used
