# Build a Landing Page

A landing page is more custom than a book: hero, features, a repository grid,
and a footer. The framework's static site generator (`framework/web/ssg`)
composes it from components with scoped styles, and `runvil build` runs the
project's generator through the `cmd/site` convention.

## 1. Create the entrypoint

Create `cmd/site/main.go`. It imports the SSG, registers components and a
layout, and writes the site to the output directory:

```go
package main

import (
	"os"

	"github.com/runvil/framework/web/ssg"
)

func main() {
	out := "site"
	if len(os.Args) > 1 {
		out = os.Args[1]
	}
	_, err := buildSite().Build(out)
	if err != nil {
		panic(err)
	}
}

func buildSite() *ssg.Site {
	return ssg.New().
		Component(ssg.Component{
			Name: "hero",
			Body: `<section><h1>Hello, Runvil.</h1></section>`,
		}).
		Layout(ssg.Layout{Name: "site", Body: "<html><body>{{component \"hero\" .Data}}</body></html>"}).
		Page(ssg.Page{Path: "/", Layout: "site", Root: "hero"})
}
```

## 2. Add the theme

The layout can render the ui framework's theme switcher and toggle styles.
The static assets `assets/style.css` and `assets/theme.css` are collected and
linked automatically:

```go
Theme: &ui.Theme{},
...
Asset("assets/theme.css", ui.ThemeModeVarsCSS+"\n"+ui.ThemeToggleCSS)
```

## 3. Build with runvil

No `manuscript/` directory means `runvil build` falls back to the `cmd/site`
shape and runs `go run ./cmd/site` with the output directory:

```sh
runvil build --output .
```

The site lands in the repository root — ready for GitHub Pages to serve as-is.