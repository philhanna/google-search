package search

import (
	"strings"

	"golang.org/x/net/html"
)

// ---------------------------------------------------------------------
// Type Definitions
// ---------------------------------------------------------------------

type HTMLDoc struct {
	HTML  string     // The HTML of the document
	Links []Link     // The links that the parser finds
	Root  *html.Node // Root of the parsed document tree
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
	p.Root = d
	return p, nil
}

// ---------------------------------------------------------------------
// Methods
// ---------------------------------------------------------------------

// getH3s is an iterator over the H3 elements in the document.
func (doc *HTMLDoc) getH3s() chan *html.Node {
	ch := make(chan *html.Node)
	go func() {
		defer close(ch)
		for x := range walk(doc.Root) {
			if x.Type == html.ElementNode {
				if x.Data == `h3` {
					ch <- x
				}
			}
		}
	}()
	return ch
}

// ---------------------------------------------------------------------
// Functions
// ---------------------------------------------------------------------

// getAttribute returns the value of the specified key in the specified
// node
func getAttribute(node *html.Node, name string) string {
	for _, attr := range node.Attr {
		if attr.Key == name {
			return attr.Val
		}
	}
	return ""
}
/*
// getChildren returns a slice of all the direct children of the
// specified node
func getChildren(node *html.Node) []*html.Node {
	children := make([]*html.Node, 0)
	for child := node.FirstChild; child != nil; child = child.NextSibling {
		children = append(children, child)
	}
	return children
}
*/

func getDescendants(node *html.Node, ch chan *html.Node) {
	for child := node.FirstChild; child != nil; child = child.NextSibling {
		ch <- child
		getDescendants(child, ch)
	}
}

// walk is an iterator that walks through all the nodes in the
// document, depth first
func walk(node *html.Node) chan *html.Node {
	ch := make(chan *html.Node)
	go func(node *html.Node) {
		defer close(ch)
		getDescendants(node, ch)
	}(node)
	return ch
}
