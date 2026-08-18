# Runvil Documentation

The documentation website for the [Runvil ecosystem](https://github.com/runvil). 
This repository is itself a dogfooding showcase: it builds its site through
[mdbind](https://github.com/runvil/mdbind), which is built on the
[Runvil Web Framework](https://github.com/runvil/framework), closing the
dependency chain:

```text
docs  ──►  mdbind  ──►  framework  ──►  libs
```

## Structure

| Path           | Description                              |
| -------------- | ---------------------------------------- |
| `manuscript/`  | The documentation, as ordered Markdown.  |
| `cmd/docs/`    | Builds the site via `book.Build`.        |
| `docs/`        | This repository's specifications (RVD-*).|
| `site/`        | Generated output (not edited by hand).   |

## Build

```sh
go run ./cmd/docs
```

The output is written to `site/` — a complete static website ready to serve
from any static host. Links are generated for the deployment base (`/docs/` by
default); use `--base /` to build a root-relative copy for local preview.

## Deploy

The site publishes to GitHub Pages from the `gh-pages` branch and is served at
`https://runvil.github.io/docs/`, linked from the
[Runvil landing page](https://runvil.github.io/).

## Prerequisites

- Go toolchain 1.22 or newer — see [go.dev/dl](https://go.dev/dl/)

## Specifications

- [RVD-8NQ2K](./docs/specs/RVD-8NQ2K-documentation-site.md) — Documentation Site.
- [RVD-P3TVZ](./docs/specs/RVD-P3TVZ-build-and-ci.md) — Build & CI.

## Related Repositories

- [framework](https://github.com/runvil/framework)
- [libs](https://github.com/runvil/libs)
- [runvil](https://github.com/runvil/runvil)
- [mdbind](https://github.com/runvil/mdbind)

## License

MIT — see [LICENSE](./LICENSE).
