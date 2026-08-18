# The Framework

[runvil/framework](https://github.com/runvil/framework) is the meta-framework:
a Go module monorepo hosting the framework's orchestration packages. These
compose and unify components that live outside the repository — such as the
libraries maintained in Runvil Libraries — into a single, consistent
experience.

## The `cli` package

An integrated command-line application model built on `flag`, `log/slog`, and
the `term` library. Applications register sub-commands with `About`, `Register`
(flag definitions), and `Run` (execution returning a canonical exit code).

```go
app := cli.NewApp("hello", "v0.1.0").
    Command(cli.NewCommand("greet", "say hello", register, run))
os.Exit(int(app.Run(os.Args[1:])))
```

Conventions are specified in the framework's CLI specifications:
[application model](https://github.com/runvil/framework/blob/main/docs/specs/RVF-M8SSR-cli-application-model.md),
[configuration](https://github.com/runvil/framework/blob/main/docs/specs/RVF-UUQ3X-cli-configuration.md),
and [errors & diagnostics](https://github.com/runvil/framework/blob/main/docs/specs/RVF-QZTY2-cli-errors-diagnostics.md).

## The `web` package

A routing and rendering layer on `net/http` and `html/template`. It provides a
`Router` with `{param}` route segments, static file serving, template sets, and
a static-site export (`web.Export`) that writes each page as
`<path>/index.html` plus assets.

```go
router := web.NewRouter()
router.Get("/{slug}", handler)
router.Static("/assets", "assets")
```

Its specification is
[RVF-8G3WQ](https://github.com/runvil/framework/blob/main/docs/specs/RVF-8G3WQ-runvil-web-framework.md).

## Framework specifications

All framework requirements are recorded under
[`framework/docs/specs/`](https://github.com/runvil/framework/tree/main/docs/specs)
with `RVF-*` SpecIDs, and cross-referenced with the libraries they depend on.
