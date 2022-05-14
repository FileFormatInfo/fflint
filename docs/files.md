---
title: Specifying Files
---

There are several different ways to specify the files that you want to check.

## Allow your shell to expand

This works well (depending on which shell you are using), but most shells have a limit on the number of parameters you can pass.

Example (note there are *no quotes*):
```
badger svg *.svg
```

## Use badger's built-in expander

Badger can expand wildcards similar to a shell, but with the addition of `**` to support zero or more directories (see [patterns](https://github.com/bmatcuk/doublestar/tree/v4#patterns) for details).

Example (note the double quotes):
```
badger svg "./**/*.svg"
```

If Badger's built-in expand is causing conflicts, you can use the `--glob` flag to change it:
* `--glob=golang` - use an expander based on Go's `filepath.Glob`
* `--glob=none` - do not do any expansion

## Send a list via stdin

This works well and has the most flexibility of all!

Example:
```
find . -name "*.svg" | badger svg -
```

## Tips

If you have a file with a list of file names, you can use the `stdin` method:

```
cat filelist.txt | badger svg -
```

If you have a directory named `-` (i.e. a single dash), you can force non-stdin mode by prefixing it with `./`:

```
badger svg ./-
```
