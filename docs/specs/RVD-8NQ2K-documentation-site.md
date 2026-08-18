# Specification — Runvil Documentation Site

| Field       | Value                                       |
| ----------- | ------------------------------------------- |
| SpecID      | RVD-8NQ2K                                   |
| Title       | Runvil Documentation Site — Initial Specification |
| Status      | Draft                                       |
| Version     | 0.1.0                                       |
| Date        | 2026-08-18                                  |
| Author      | Runvil Contributors                         |
| Domain      | Documentation                              |

## 1. Context

**runvil/docs** hosts the Runvil ecosystem documentation as a single,
book-shaped website. Its content is a Markdown manuscript assembled by the
mdbind site builder, making this repository the dogfooding consumer that closes
the dependency chain `docs -> mdbind -> framework/web -> libs`.

## 2. Problem Statement

Runvil's knowledge is spread across multiple repositories, README files, and
specifications. There is no single place a contributor or user can read the
whole story — philosophy, tooling, libraries, framework, and builder — in
reading order. The documentation site solves that by presenting the ecosystem
as one book, authored as plain Markdown and generated with Runvil's own tools.

## 3. Goals

- G1 — Provide a single documentation book covering the whole ecosystem.
- G2 — Author content as an ordered Markdown manuscript.
- G3 — Generate the site through mdbind's public `book` package.
- G4 — Dogfood the full dependency chain end to end.

## 4. Non-Goals

- NG1 — Search index, translations, or versioned documentation in this phase.
- NG2 — A custom content model beyond mdbind's manuscript conventions.

## 5. Requirements

### 5.1 Content

| ID          | Requirement                                                   | Priority |
| ----------- | ------------------------------------------------------------- | -------- |
| DOC-SITE-001 | Content lives in `manuscript/*.md` in reading order.          | Must     |
| DOC-SITE-002 | The book covers framework, libs, runvil, mdbind, and docs.    | Must     |
| DOC-SITE-003 | In-book links use site paths (`/chapters/...`).               | Must     |
| DOC-SITE-004 | Cross-repository links use GitHub URLs.                       | Must     |

### 5.2 Generation

| ID          | Requirement                                                   | Priority |
| ----------- | ------------------------------------------------------------- | -------- |
| DOC-SITE-005 | The site is generated through `book.Build`, not internal APIs.| Must     |
| DOC-SITE-006 | The generated site includes a table of contents and chapters.| Must     |
| DOC-SITE-007 | Generation is deterministic for identical manuscripts.        | Must     |

## 6. Non-Functional Requirements

- NFR1 — **Memory safety.** The `unsafe` package must not be used.
- NFR2 — **Portability.** Linux, macOS, and Windows.
- NFR3 — **Quality gates.** `gofmt`, `go vet ./...`, and `go test ./...` must pass.

## 7. Success Criteria

- S1 — `go run ./cmd/docs` produces `site/index.html` and one page per chapter.
- S2 — Every ecosystem repository is covered by a chapter.
- S3 — The build imports only public mdbind packages.

## 8. Related Specifications

| SpecID    | Title                                           |
| --------- | ----------------------------------------------- |
| [RVD-P3TVZ](./RVD-P3TVZ-build-and-ci.md)          | Documentation Build & CI           |
| [RVM-5F9TL](https://github.com/runvil/mdbind/blob/main/docs/specs/RVM-5F9TL-mdbind-site-builder.md) | mdbind Site Builder |

## 9. References

- [RVF-8G3WQ](https://github.com/runvil/framework/blob/main/docs/specs/RVF-8G3WQ-runvil-web-framework.md) — Runvil Web Framework.
- [RVF-QOFJK](https://github.com/runvil/framework/blob/main/docs/specs/RVF-QOFJK-runvil-meta-framework.md) — Runvil Framework initial specification.
- [RVL-4Y8UP](https://github.com/runvil/libs/blob/main/docs/specs/RVL-4Y8UP-runvil-libraries.md) — Runvil Libraries initial specification.