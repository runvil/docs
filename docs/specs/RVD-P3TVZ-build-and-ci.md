# Specification — Documentation Build & CI

| Field       | Value                                       |
| ----------- | ------------------------------------------- |
| SpecID      | RVD-P3TVZ                                   |
| Title       | Documentation Build & CI                   |
| Status      | Draft                                       |
| Version     | 0.1.0                                       |
| Date        | 2026-08-18                                  |
| Author      | Runvil Contributors                         |
| Domain      | Documentation — Build                       |

## 1. Context

The documentation site is generated from its manuscript by `cmd/docs`, which
calls mdbind's public `book.Build`. This specification defines how the site is
built and how CI keeps it correct, reproducible, and verified on every change.

## 2. Problem Statement

A generated site can silently drift from its manuscript: chapters renamed,
links broken, or generation failing only on clean checkouts. Without a defined
build entry point and a CI gate, the published documentation is not guaranteed
to reflect the source of truth.

## 3. Goals

- G1 — Provide a single build entry point: `go run ./cmd/docs`.
- G2 — Keep the site build deterministic and CI-verified.
- G3 — Restrict imports to public mdbind packages.

## 4. Non-Goals

- NG1 — Deployment or hosting automation in this phase.
- NG2 — Pull-request preview environments.

## 5. Requirements

### 5.1 Build

| ID         | Requirement                                                       | Priority |
| ---------- | ----------------------------------------------------------------- | -------- |
| DOC-BLD-001 | `go run ./cmd/docs` builds the manuscript into `./site`.          | Must     |
| DOC-BLD-002 | The site build uses only mdbind's public package surface.         | Must     |
| DOC-BLD-003 | The build reports created files deterministically.                | Must     |

### 5.2 CI

| ID         | Requirement                                                       | Priority |
| ---------- | ----------------------------------------------------------------- | -------- |
| DOC-BLD-004 | CI runs `gofmt`, `go vet ./...`, and `go test ./...`.             | Must     |
| DOC-BLD-005 | CI builds the site and verifies `site/index.html` and chapter pages exist. | Must |

## 6. Non-Functional Requirements

- NFR1 — **Memory safety.** The `unsafe` package must not be used.
- NFR2 — **Reproducibility.** A clean checkout builds byte-identical output.
- NFR3 — **Portability.** Linux, macOS, and Windows.
- NFR4 — **Quality gates.** `gofmt`, `go vet ./...`, and `go test ./...` must pass.

## 7. Success Criteria

- S1 — A clean checkout followed by CI leaves a complete `site/`.
- S2 — No import of `github.com/runvil/mdbind/internal/...`.
- S3 — The `site/` directory is treated as generated output.

## 8. Related Specifications

| SpecID    | Title                                           |
| --------- | ----------------------------------------------- |
| [RVD-8NQ2K](./RVD-8NQ2K-documentation-site.md)  | Runvil Documentation Site — Initial Specification |
| [RVM-K2DM8](https://github.com/runvil/mdbind/blob/main/docs/specs/RVM-K2DM8-cli-workflows.md) | mdbind CLI & Workflows |

## 9. References

- [RVF-UUQ3X](https://github.com/runvil/framework/blob/main/docs/specs/RVF-UUQ3X-cli-configuration.md) — CLI Configuration.
- [RVF-QZTY2](https://github.com/runvil/framework/blob/main/docs/specs/RVF-QZTY2-cli-errors-diagnostics.md) — CLI Errors & Diagnostics.