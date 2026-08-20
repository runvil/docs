# The Libraries

[runvil/libs](https://github.com/runvil/libs) is the monorepo of shared,
reusable libraries for the Runvil ecosystem. It is one of the ecosystems Runvil
*integrates* rather than re-implements.

## `core`

Shared primitives: the common error type pairing a message with a canonical
exit code, and the `ExitCode` type itself.

```go
// ExitCode maps to process exit codes used across Runvil applications.
const (
    ExitCodeSuccess ExitCode = 0
    ExitCodeFailure ExitCode = 1
    ExitCodeUsage   ExitCode = 2
)
```

See the specification:
[RVL-W0J2X](https://github.com/runvil/libs/blob/main/docs/specs/RVL-CORE-W0J2X-errors-exit-codes.md).

## `term`

Terminal I/O and rendering: ANSI styling, colors, and lightweight terminal
interaction used by the framework's CLI and developer tooling.

See the specification:
[RVL-R934Y](https://github.com/runvil/libs/blob/main/docs/specs/RVL-TERM-R934Y-terminal-io-rendering.md).

## Using the libraries

```sh
go get github.com/runvil/libs
```

```go
import "github.com/runvil/libs/core"

err := core.FailureError("something broke")
os.Exit(int(err.ExitCode()))
```

All library requirements are recorded under
[`libs/docs/specs/`](https://github.com/runvil/libs/tree/main/docs/specs) with
`RVL-*` SpecIDs.
