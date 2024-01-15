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

	// Parse the document for links
	for h3 := range p.getH3s() {
		
		// Get the URL
		urlPtr := getURL(h3)
		if urlPtr == nil {
			continue
		}
		
		// Get the title associated with the <h3>
		title := h3.FirstChild.Data

		link := makeLink(*urlPtr, title)
		if link != nil {
			p.Links = append(p.Links, *link)
		}
	}
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

func getDescendants(node *html.Node, ch chan *html.Node) {
	for child := node.FirstChild; child != nil; child = child.NextSibling {
		ch <- child
		getDescendants(child, ch)
	}
}

// getURL finds the URL for the specified <h3>
//
// Once an <h3> has been found, the title will be in the first child that is a text node.
// To find the corresponding URL, the following algorithm is used:
//
//  1. Find the first ancestor that is an HTML element with a tag of `a`.
//  2. If no such ancestor exists, return nil
//  3. Return the value of the `href` attribute
func getURL(node *html.Node) *string {
	for {
		if node == nil {
			return nil
		}
		if node.Type == html.ElementNode && node.Data == `a` {
			href := getAttribute(node, `href`)
			if href == "" {
				return nil
			}
			return &href
		}
		node = node.Parent
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
