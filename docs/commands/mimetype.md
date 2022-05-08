---
h1: The mimetype Command
title: 'mimetype: Validate (or report) MIME content types - Badger'
name: badger mimetype
synopsis: Validate (or report) MIME content types
usage: badger mimetype [flags] files...
options:
- name: allowUnknown
  default_value: "true"
  usage: Allow application/octet-stream
- name: help
  shorthand: h
  default_value: "false"
  usage: help for mimetype
- name: report
  default_value: "true"
  usage: Print summary report (default is true)
inherited_options:
- name: config
  usage: config file (default is $HOME/.badger.yaml)
- name: debug
  default_value: "false"
  usage: Debugging output
- name: filesize
  default_value: any
  usage: Range of allowed file size
- name: glob
  usage: |
    Algorithm to use to expanding wildcards in file names [ doublestar | golang | none ]
- name: output
  shorthand: o
  default_value: text
  usage: Output format [ json | text ]
- name: progress
  default_value: "true"
  usage: Show progress bar (default is false when stderr is piped)
- name: showDetail
  default_value: "true"
  usage: Show detailed data about each test
- name: showFiles
  shorthand: f
  default_value: "false"
  usage: |
    Show each file tested (default is false when stderr is piped)
- name: showPassing
  default_value: "false"
  usage: Show passing files and tests
- name: showTests
  shorthand: t
  default_value: "false"
  usage: Show each test performed
- name: showTotal
  default_value: "true"
  usage: Show total files tested, passed and failed
see_also:
- badger - Badgers you if your file formats are invalid
---
{% comment %}NOTE: this file is auto-generated by bin/docgen.sh.  Manual edits will be overwritten!{% endcomment -%}
{% include command.html %}