package search

import (
	"io"
	"net/http"
	urn "net/url"
	"strings"

	"golang.org/x/net/html"
)

const (
	GOOGLE_URL = "https://www.google.com/search?q="
	UA         = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36"
)

// ---------------------------------------------------------------------
// Type Definitions
// ---------------------------------------------------------------------

type HTMLDoc struct {
	HTML  string     // The HTML of the document
	Links []Link     // The links that the parser finds
	Root  *html.Node // Root of the parsed document tree
}

var (
	// Downloader is a variable pointing to a function that performs the
	// HTTP GET request and returns a string of HTML. This variable can
	// be overridden with another function pointer for unit testing
	// purposes.
	Downloader = func(url string) (string, error) {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return "", err
		}
		req.Header.Set("User-Agent", UA)
		client := new(http.Client)
		resp, err := client.Do(req)
		if err != nil {
			return "", err
		}
		byteData, _ := io.ReadAll(resp.Body)
		return string(byteData), nil
	}

	// DefaultDownloader makes it possible to restore the original
	// Downloader variable
	DefaultDownloader = Downloader
)

// ---------------------------------------------------------------------
// Constuctors
// ---------------------------------------------------------------------

// Download accepts a query and passes it to Google search, and returns
// the HTML created by Google.  The function that performs the HTTP Get
// is pointed to by the Downloader variable, so it is possible to mock
// it with an object that supplies HTML from a local source.
func Download(query string) (*HTMLDoc, error) {
	url := GOOGLE_URL
	url += urn.QueryEscape(query)
	inputHTML, err := Downloader(url)
	if err != nil {
		return nil, err
	}
	doc := NewHTMLDoc(inputHTML)
	return doc, nil
}

// NewHTMLDoc creates a new HTML document with the specified HTML data,
// parses it for links, and returns a pointer to it
func NewHTMLDoc(input string) *HTMLDoc {
	p := new(HTMLDoc)
	p.HTML = input
	p.Links = make([]Link, 0)
	d, _ := html.Parse(strings.NewReader(p.HTML))
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
	return p
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
					if getAttribute(x, "aria-hidden") == "" {
						ch <- x
					}
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
