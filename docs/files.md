---
title: Specifying Files
---

There are several different ways to specify the files that you want to check.

## Allow your shell to expand

This works well (depending on which shell you are using), but most shells have a limit on the number of parameters you can pass.

Example (note there are *no quotes*):
```
fflint svg *.svg
```

## Use fflint's built-in expander

fflint can expand wildcards similar to a shell, but with the addition of `**` to support zero or more directories (see [patterns](https://github.com/bmatcuk/doublestar/tree/v4#patterns) for details).

Example (note the double quotes):
```
fflint svg "./**/*.svg"
```

If fflint's built-in expand is causing conflicts, you can use the `--glob` flag to change it:
* `--glob=golang` - use an expander based on Go's `filepath.Glob`
* `--glob=none` - do not do any expansion

## Send a list via stdin

This works well and has the most flexibility of all, since you can use `find` (or a pipe or anything else) to make the list.

Example:
```
find . -name "*.svg" | fflint svg @-
```

## Send a single file via stdin

This is useful if you are downloading or generating a file and do not need to store it locally.

Example:
```
curl --silent https://www.fflint.dev/favicon.ico | fflint ico -
```

If you have a directory named `-` (i.e. a single dash), you can force non-stdin mode by prefixing it with `./`:

```
fflint svg ./-
```
