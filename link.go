package search

import (
	"regexp"
	"strings"
)

type Link struct {
	URL   string `json:"url"`
	Title string `json:"title"`
}

var (
	reSpaces = regexp.MustCompile(`\s+`)
)

// makeLink creates a Link object from a URL and title
func makeLink(url, title string) *Link {

	// Sanitize the URL
	const prefix = `/url?q=`
	url = strings.TrimPrefix(url, prefix)
	p := strings.Index(url, "&")
	if p != -1 {
		url = url[:p]
	}

	// Sanitize the title
	const cutset = "\n\r\t\x80"
	for _, c := range cutset {
		title = strings.ReplaceAll(title, string(c), " ")
	}
	title = reSpaces.ReplaceAllString(title, " ")
	title = strings.TrimSpace(title)

	// Done
	return &Link{
		URL:   url,
		Title: title,
	}
}
