package search

import (
	"strings"

	"golang.org/x/net/html"
)

// ---------------------------------------------------------------------
// Type Definitions
// ---------------------------------------------------------------------

type Link struct {
	URL   string
	Title string
}

type HTMLDoc struct {
	HTML      string     // The HTML of the document
	Links     []Link     // The links that the parser finds
	ParsedDoc *html.Node // Root of the parsed document tree
}

// ---------------------------------------------------------------------
// Constuctors
// ---------------------------------------------------------------------

// NewHTMLDoc creates a new HTML document with the specified HTML data,
// parses it for links, and returns a pointer to it
func NewHTMLDoc(input string) (*HTMLDoc, error) {
	p := new(HTMLDoc)
	p.HTML = input
	p.Links = make([]Link, 0)
	d, err := html.Parse(strings.NewReader(p.HTML))
	if err != nil {
		return nil, err
	}
	p.ParsedDoc = d
	return p, nil
}
