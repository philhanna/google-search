package search

import (
	// "golang.org/x/net/html"
	// "strings"
)

// ---------------------------------------------------------------------
// Type Definitions
// ---------------------------------------------------------------------

type Link struct {
	URL   string
	Title string
}

type HTMLDoc struct {
	Data  string // The HTML of the document
	Links []Link // The links that the parser finds
}

// ---------------------------------------------------------------------
// Constuctors
// ---------------------------------------------------------------------

// NewHTMLDoc creates a new HTML document with the specified HTML data,
// parses it for links, and returns a pointer to it
func NewHTMLDoc(input string) (*HTMLDoc, error) {
	p := new(HTMLDoc)
	p.Data = input
	p.Links = make([]Link, 0)
	err := p.parse()
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (doc *HTMLDoc) parse() error {
	return nil
}
