# Runvil Documentation — Repository Hub

This directory is the documentation hub for the Runvil documentation site. The
website itself is generated from the book in [`manuscript/`](../manuscript) by
mdbind.

## Contents

- [Specifications](./specs/index.md) — Formal specifications for this repository.
- [Initial specification](./specs/RVD-8NQ2K-documentation-site.md) — scope and design.
- [Build & CI specification](./specs/RVD-P3TVZ-build-and-ci.md) — generation and verification.

## Conventions

- The generated site lives in `site/` and must not be edited by hand.
- The source of truth is the manuscript under `manuscript/`.
- Site settings live in `runvil.yaml`; build with `runvil build`.