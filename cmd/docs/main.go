// The Runvil documentation site: builds the manuscript into a static website
// through mdbind's public book package, following the dependency chain
// docs -> mdbind -> framework/web -> libs.
package main

import (
	"flag"
	"fmt"
	"os"

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
	})
	if err != nil {
		fmt.Fprintln(os.Stderr, "docs:", err)
		os.Exit(1)
	}
	for _, p := range created {
		fmt.Println("created " + p)
	}
}
