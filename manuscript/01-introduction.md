# Introduction

**Runvil** is a meta-framework written in [Go](https://go.dev/). Instead of a
conventional framework built inside a single, self-contained ecosystem, Runvil
composes modules sourced across multiple ecosystems and repositories and
orchestrates them into one cohesive, high-performance foundation for building a
wide range of applications — from web services and CLI tools to background
workers and desktop applications.

## Design philosophy

- **Meta-framework.** Runvil does not re-implement what other ecosystems do
  well; it composes and unifies them under one consistent experience.
- **Stdlib-first.** Where the Go standard library is sufficient — argument
  parsing (`flag`), logging (`log/slog`), configuration (`os`) — Runvil uses it.
- **Safe by design.** Go's memory safety, no `unsafe`, no manual memory
  management.
- **Spec-driven.** Every component ships with a formal specification
  (e.g., `RVF-*`, `RVL-*`, `RVN-*`, `RVM-*`) describing requirements and
  conventions before behavior is added.
- **Dogfooded.** Runvil's own tooling is built on Runvil.

## The ecosystem at a glance

| Repository | Purpose |
| ---------- | ------- |
| [runvil/framework](https://github.com/runvil/framework) | The meta-framework: `cli`, `web`, and orchestration packages. |
| [runvil/libs](https://github.com/runvil/libs) | Shared, reusable libraries: `core`, `term`. |
| [runvil/runvil](https://github.com/runvil/runvil) | The developer tool: scaffolding, testing, project information. |
| [runvil/mdbind](https://github.com/runvil/mdbind) | The site builder: Markdown folders into book-shaped websites. |

Continue to [Getting Started](/getting-started) to build your first
Runvil application.
