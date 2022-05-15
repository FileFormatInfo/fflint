---
title: Examples
---

### Working on files with errors

Use `--output=filenames` to get a list of files that failed.

Open the files in an editor (`vi` in this example):
```bash
badger frontmatter  --required=title --output=filenames --progress=false "./docs/**/*.md" | xargs vi
```

Optimize every SVG files that has extra namespaces:
```bash
badger svg --namespace --output=filenames --progress=false "**/*.svg" | xargs svgo
```