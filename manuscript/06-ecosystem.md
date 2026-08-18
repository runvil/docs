# The Ecosystem

Runvil is organized as an ecosystem of focused repositories that depend on one
another in a single, explicit chain.

## Repository map

| Repository | Role | Documentation |
| ---------- | ---- | ------------- |
| [runvil/framework](https://github.com/runvil/framework) | Meta-framework (`cli`, `web`). | [`framework/docs`](https://github.com/runvil/framework/tree/main/docs) |
| [runvil/libs](https://github.com/runvil/libs) | Shared libraries (`core`, `term`). | [`libs/docs`](https://github.com/runvil/libs/tree/main/docs) |
| [runvil/runvil](https://github.com/runvil/runvil) | Developer tool. | [`runvil/docs`](https://github.com/runvil/runvil/tree/main/docs) |
| [runvil/mdbind](https://github.com/runvil/mdbind) | Site builder. | [`mdbind/docs`](https://github.com/runvil/mdbind/tree/main/docs) |
| [runvil/docs](https://github.com/runvil/docs) | This documentation site. | — |

## Dependency graph

```text
docs  ──►  mdbind  ──►  framework  ──►  libs
runvil ──►  framework ──►  libs
```

Each repository is an independent public Go module with a green CI on `main`
and a versioned release. Specification conventions are shared across the
ecosystem:

- **SpecIDs** — 5-character alphanumeric codes (e.g., `RVF-8G3WQ`).
- **Requirement IDs** — prefixed per domain (`FRK-*`, `RVL-*`, `RND-*`, `MDB-*`).
- **Cross-repository links** — specifications reference each other via GitHub URLs.

## Contributing

1. Read the relevant specifications in `docs/specs/`.
2. Follow the conventions in this book.
3. Open a pull request; CI runs `gofmt`, `go vet ./...`, and `go test ./...`.

## License

The Runvil ecosystem is available under the MIT license. See each
repository's `LICENSE` file.
