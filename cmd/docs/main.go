// The Runvil documentation site: builds the manuscript into a static website
// through mdbind's public book package, following the dependency chain
// docs -> mdbind -> framework/web -> libs.
package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/runvil/framework/web"
	"github.com/runvil/mdbind/book"
)

func main() {
	input := flag.String("input", "manuscript", "manuscript directory")
	output := flag.String("output", "site", "output directory")
	base := flag.String("base", "", "URL base the site is served under (default DOCS_BASE or /)")
	flag.Parse()

	basePath := *base
	if basePath == "" {
		basePath = os.Getenv("DOCS_BASE")
	}
	if basePath == "" {
		basePath = "/"
	}

	created, err := book.Build(book.Config{
		Input:    *input,
		Output:   *output,
		Title:    "Runvil Documentation",
		Author:   "Runvil Contributors",
		BasePath: basePath,
		NavLinks: []book.Link{
			{Text: "Runvil", URL: "/"},
			{Text: "Ecosystem", URL: "https://github.com/runvil"},
		},
		FooterText: "Runvil Documentation — MIT License",
		Theme: &web.Theme{
			Light: web.Palette{
				Primary:        "#7c3aed",
				PrimaryContent: "#ffffff",
				Accent:         "#b45309",
			},
			Dark: web.Palette{
				Primary:        "#a78bfa",
				PrimaryContent: "#1e1b4b",
				Accent:         "#f0b429",
			},
		},
	})
	if err != nil {
		fmt.Fprintln(os.Stderr, "docs:", err)
		os.Exit(1)
	}
	for _, p := range created {
		fmt.Println("created " + p)
	}
}
