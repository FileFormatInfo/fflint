---
title: FAQ - FFLint
h1: Frequently Asked Questions
faq:
  - q: How can I edit just the files with errors?
  - a: |
        Use `--output=filenames` to get a list of files that failed.

        ##### Some examples

        Open the files in an editor (`vi` in this example):
        ```bash
        fflint frontmatter --required=title --output=filenames --progress=false "./docs/**/*.md" | xargs vi
        ```

        Optimize every SVG files that has extra namespaces:
        ```bash
        fflint svg --namespace --output=filenames --progress=false "**/*.svg" | xargs svgo
        ```

  - q: "How can I contact you?"
    a: "[Github issues](https://github.com/FileFormatInfo/fflint/issues)"

later:
  - add paid support
---

{% for entry in page.faq %}
<h4>{{entry.q}}</h4>
<div class="ms-3 mb-3">{{entry.a | markdownify}}</div>
{% endfor %}

