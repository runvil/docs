# The Site Builder

[runvil/mdbind](https://github.com/runvil/mdbind) assembles folders of Markdown
into book-shaped static websites: ordered chapters, a table of contents, and
prev/next navigation. It is built on the framework's `web` package and is the
engine behind this very documentation site.

## Command line

```sh
go install github.com/runvil/mdbind/cmd/mdbind@v0.5.0

mdbind init          # scaffold a sample manuscript
mdbind build         # build ./ -> site/
mdbind serve         # preview over HTTP
```

Settings honor the Runvil precedence convention: flags > `MDBIND_*`
environment variables > defaults.

## Library

```go
created, err := book.Build(book.Config{
    Input:  "manuscript",
    Output: "site",
    Title:  "Runvil Documentation",
    Author: "Runvil Contributors",
})
```

Sub-directories in the manuscript become chapters with `N.M` subchapters
served at `/tutorials/{chapter}/{sub}/`; every chapter page carries a
breadcrumb trail back to Home, and the sidebar expands the active section's
subchapters. This documentation site itself dogfoods the builder.

## Specifications

- [RVM-5F9TL](https://github.com/runvil/mdbind/blob/main/docs/specs/RVM-5F9TL-mdbind-site-builder.md)
  — the builder's initial specification.
- [RVM-K2DM8](https://github.com/runvil/mdbind/blob/main/docs/specs/RVM-K2DM8-cli-workflows.md)
  — the CLI and its workflows.
