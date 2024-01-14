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
	HTML  string // The HTML of the document
	Links []Link // The links that the parser finds
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
	err := p.parse()
	return p, err
}

func (doc *HTMLDoc) parse() error {
	elemRoot, err := html.Parse(strings.NewReader(doc.HTML))
	if err != nil {
		return err
	}
	err = handleElementNode(elemRoot)
	if err != nil {
		return err
	}
	return nil
}

func handleElementNode(elem *html.Node) error {
	if elem.Data == "div" {
		class := getAttribute(elem, "class")
		if isLinkDiv(class) {
			_ = class
		}
	}
	for child := elem.FirstChild; child != nil; child = child.NextSibling {
		switch child.Type {
		case html.ElementNode:
			handleElementNode(child)
		}
	}
	return nil
}

func getAttribute(node *html.Node, key string) string {
	for _, attr := range node.Attr {
		if attr.Key == key {
			return attr.Val
		}
	}
	return ""
}

func isLinkDiv(class string) bool {
	return strings.Contains(class, "egMi0") && strings.Contains(class, "kCrYT")
}

