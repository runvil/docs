# Getting Started

The fastest way into Runvil is the `runvil` developer tool, which scaffolds a
project that already composes the ecosystem.

## Prerequisites

- Go toolchain 1.22 or newer — see [go.dev/dl](https://go.dev/dl/)

## Install the developer tool

```sh
go install github.com/runvil/runvil/cmd/runvil@v0.1.0
```

## Create a project

```sh
runvil new hello --module example.com/hello
cd hello
runvil test
```

The scaffolded project depends on `github.com/runvil/framework` and
`github.com/runvil/libs`, and `runvil test` runs its tests through the Go
toolchain.

## What you can run

| Command      | Purpose                                        |
| ------------ | ---------------------------------------------- |
| `runvil version` | Print the tool and framework versions.     |
| `runvil info`    | Print environment and project information. |
| `runvil new`     | Scaffold a new Runvil project.             |
| `runvil test`    | Run the current project's tests.           |

## Next steps

- Read how the [framework](/chapters/framework) composes its packages.
- Explore the [shared libraries](/chapters/libraries).
- Publish documentation with the [site builder](/chapters/site-builder).
