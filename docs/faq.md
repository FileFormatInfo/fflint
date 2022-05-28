---
title: FAQ - Badger
h1: Frequently Asked Questions
faq:
  - q: How can I edit just the files with errors?
  - a: |
        Use `--output=filenames` to get a list of files that failed.

        Open the files in an editor (`vi` in this example):
        ```bash
        badger frontmatter --required=title --output=filenames --progress=false "./docs/**/*.md" | xargs vi
        ```

        Optimize every SVG files that has extra namespaces:
        ```bash
        badger svg --namespace --output=filenames --progress=false "**/*.svg" | xargs svgo
        ```

  - q: My company does not allow AGPL licensed software.  How can I use it?
    a: No problem!  See the [pricing](/pricing.html) page for affordable non-AGPL licenses.

later:
  - q: "How can I contact you?"
    a: "[Github issues](), [contact page](), [paid support]()"

---

{% for entry in page.faq %}
<h4>{{entry.q}}</h4>
<p>{{entry.a | markdownify}}</p>
{% endfor %}

