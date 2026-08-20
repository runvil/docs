# Build a Documentation Site

This tutorial turns a folder of Markdown into a book-shaped documentation
site, shaped with `runvil.yaml` so no project-specific command is needed.

## 1. Scaffold the manuscript

`runvil new` scaffolds a Go module, but a site needs content. Create a
`manuscript/` directory at the module root:

```sh
mkdir manuscript
```

## 2. Author chapters

Each Markdown file is a chapter, ordered by numeric filename prefix:

```text
manuscript/
  01-introduction.md
  02-getting-started.md
  03-reference.md
```

Longer sections become chapters with subchapters: put a directory in
manuscript/ and mdbind numbers its files as `N.M` sections, served at
`/tutorials/{chapter}/{sub}/`. An optional `index.md` inside the directory
becomes the chapter page; without one, the chapter page lists its sections
automatically.

## 3. Shape the site with runvil.yaml

Configure title, navigation, footer, base path, and the theme palette in a
single file — no `cmd/` entrypoint needed:

```yaml
# runvil.yaml
title: My Docs
author: Runvil Contributors
base: /docs/
nav:
  - text: Runvil
    url: /
theme:
  light:
    primary: "#7c3aed"
  dark:
    primary: "#a78bfa"
```

## 4. Build

`runvil build` detects `manuscript/`, reads `runvil.yaml`, and delegates to the
mdbind site builder:

```sh
runvil build
```

The command prints every created path, produces a table of contents with
subchapter sections, breadcrumbs, and prev/next navigation, and ships a
light/dark theme switcher driven entirely by the configured palette.