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

// ---------------------------------------------------------------------
// Methods
// ---------------------------------------------------------------------

// getLink starts from a <div> statement and extracts the Link
// it contains, if any
func (doc *HTMLDoc) getLink(node *html.Node) *Link {
	return nil
}

// handleElementNode extracts a link from the specified node, if it has
// one, then recursively applies the same function to all its
// descendants
func (doc *HTMLDoc) handleElementNode(elem *html.Node) error {
	if elem.Data == "div" {
		class := getAttribute(elem, "class")
		if isLinkDiv(class) {
			link := doc.getLink(elem)
			if link != nil {
				doc.Links = append(doc.Links, *link)	
			}
		}
	}
	for child := elem.FirstChild; child != nil; child = child.NextSibling {
		switch child.Type {
		case html.ElementNode:
			doc.handleElementNode(child)
		}
	}
	return nil
}

// parse is an internal method called by the constructor that builds the
// list of links in this document
func (doc *HTMLDoc) parse() error {
	elemRoot, err := html.Parse(strings.NewReader(doc.HTML))
	if err != nil {
		return err
	}
	err = doc.handleElementNode(elemRoot)
	if err != nil {
		return err
	}
	return nil
}

// ---------------------------------------------------------------------
// Functions
// ---------------------------------------------------------------------

// getAttribute returns the value of the specified attribute in a node
func getAttribute(node *html.Node, key string) string {
	for _, attr := range node.Attr {
		if attr.Key == key {
			return attr.Val
		}
	}
	return ""
}

// getTitle returns the title of the link associated with this <div>, if
// there is one.
func getTitle(node *html.Node) string {
	var title string
	return title
}

// getURL returns the URL associated with this link <div>, if there is
// one
func getURL(node *html.Node) string {
	var url string
	elemA := node.FirstChild
	if elemA != nil {
		url = getAttribute(elemA, "href")
	}
	return url
}

// isLinkDiv returns true if the specified class string indicates that
// this is a <div> that contains a link
func isLinkDiv(class string) bool {
	return strings.Contains(class, "egMi0") && strings.Contains(class, "kCrYT")
}
