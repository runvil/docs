# Specification — Runvil Documentation Site

| Field       | Value                                       |
| ----------- | ------------------------------------------- |
| SpecID      | RVD-8NQ2K                                   |
| Title       | Runvil Documentation Site — Initial Specification |
| Status      | Draft                                       |
| Version     | 0.4.0                                       |
| Date        | 2026-08-18                                  |
| Author      | Runvil Contributors                         |
| Domain      | Documentation                              |

## 1. Context

**runvil/docs** hosts the Runvil ecosystem documentation as a single,
book-shaped website. Its content is a Markdown manuscript assembled by the
mdbind site builder and driven by `runvil build`, making this repository the
dogfooding consumer that closes the dependency chain
`runvil/docs -> runvil -> mdbind + framework/web + framework/ui -> libs`.

## 2. Problem Statement

Runvil's knowledge is spread across multiple repositories, README files, and
specifications. There is no single place a contributor or user can read the
whole story — philosophy, tooling, libraries, framework, and builder — in
reading order. The documentation site solves that by presenting the ecosystem
as one book, authored as plain Markdown and generated with Runvil's own tools.

## 3. Goals

- G1 — Provide a single documentation book covering the whole ecosystem.
- G2 — Author content as an ordered Markdown manuscript with subchapter sections.
- G3 — Generate the site through `runvil build`, which delegates to mdbind's public `book` package.
- G4 — Dogfood the full dependency chain end to end.

## 4. Non-Goals

- NG1 — Search index, translations, or versioned documentation in this phase.
- NG2 — A custom content model beyond mdbind's manuscript conventions.

## 5. Requirements

### 5.1 Content

| ID          | Requirement                                                   | Priority |
| ----------- | ------------------------------------------------------------- | -------- |
| DOC-SITE-001 | Content lives in `manuscript/**/*.md` in reading order, with directories forming `N.M` subchapters. | Must     |
| DOC-SITE-002 | The book covers framework, libs, runvil, mdbind, docs, and tutorials. | Must     |
| DOC-SITE-003 | In-book links use site paths under the deployment base (`/docs/chapters/...`). | Must     |
| DOC-SITE-004 | Cross-repository links use GitHub URLs.                       | Must     |

### 5.2 Generation

| ID          | Requirement                                                   | Priority |
| ----------- | ------------------------------------------------------------- | -------- |
| DOC-SITE-005 | The site is generated through `runvil build`, which delegates to `book.Build`; no project-specific `cmd` exists. | Must |
| DOC-SITE-006 | Site settings (title, author, base, nav, footer, palette) live in `runvil.yaml`. | Must |
| DOC-SITE-007 | The generated site includes a table of contents, subchapter sections, and breadcrumbs. | Must |
| DOC-SITE-008 | Generation is deterministic for identical manuscripts.        | Must |
| DOC-SITE-009 | Generated links resolve under a configurable base path, deploying at `runvil.github.io/docs/`. | Must |
| DOC-SITE-010 | Pages render mdbind's navbar, sidebar, and footer chrome.         | Must |
| DOC-SITE-011 | Pages render the light/dark theme switcher (system scheme by default). | Must |
| DOC-SITE-012 | The site overrides palette tokens (primary/accent) per mode while the rest fall back to defaults. | Must |

## 6. Non-Functional Requirements

- NFR1 — **Memory safety.** The `unsafe` package must not be used.
- NFR2 — **Portability.** Linux, macOS, and Windows.
- NFR3 — **Quality gates.** `gofmt`, `go vet ./...`, and `go test ./...` must pass.

## 7. Success Criteria

- S1 — `runvil build` produces `site/index.html` and one page per chapter and subchapter.
- S2 — Every ecosystem repository is covered by a chapter, plus a tutorials section.
- S3 — The build pipeline imports no project-specific builder.

## 8. Related Specifications

| SpecID    | Title                                           |
| --------- | ----------------------------------------------- |
| [RVD-P3TVZ](./RVD-P3TVZ-build-and-ci.md)          | Documentation Build & CI           |
| [RVM-5F9TL](https://github.com/runvil/mdbind/blob/main/docs/specs/RVM-5F9TL-mdbind-site-builder.md) | mdbind Site Builder |
| [RVN-MPLQ4](https://github.com/runvil/runvil/blob/main/docs/specs/RVN-MPLQ4-project-building.md) | Project Building |

## 9. References

- [RVF-8G3WQ](https://github.com/runvil/framework/blob/main/docs/specs/RVF-8G3WQ-runvil-web-framework.md) — Runvil Web Framework.
- [RVF-QOFJK](https://github.com/runvil/framework/blob/main/docs/specs/RVF-QOFJK-runvil-meta-framework.md) — Runvil Framework initial specification.
- [RVL-4Y8UP](https://github.com/runvil/libs/blob/main/docs/specs/RVL-4Y8UP-runvil-libraries.md) — Runvil Libraries initial specification.