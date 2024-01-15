package main

import (
	"fmt"
	"io"
	"log/slog"
	"os"
	"path/filepath"
	"strings"

	"github.com/philhanna/google_search"
	"golang.org/x/net/html"
)

func init() {
	logger := slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
		Level: slog.LevelWarn,
	}))
	slog.SetDefault(logger)
}

func main() {
	filename := filepath.Join(search.ProjectRoot, "testdata", "tdd.html")
	fp, err := os.Open(filename)
	if err != nil {
		errmsg := fmt.Sprintf("%v", err)
		slog.Error(errmsg)
		os.Exit(1)
	}
	defer fp.Close()

	body, err := io.ReadAll(fp)
	if err != nil {
		errmsg := fmt.Sprintf("%v", err)
		slog.Error(errmsg)
		os.Exit(1)
	}

	sbody := string(body)
	reader := strings.NewReader(sbody)
	doc, err := html.Parse(reader)
	if err != nil {
		errmsg := fmt.Sprintf("%v", err)
		slog.Error(errmsg)
		os.Exit(1)
	}
	handleElementNode(doc)
}

func handleElementNode(node *html.Node) {
	if node.Data == "div" {
		class := getAttribute(node, "class")
		if isLinkDiv(class) {
			elemA := node.FirstChild
			if hasURL(elemA) {
				href := getAttribute(elemA, "href")
				fmt.Printf("DEBUG: %v\n", href)
			}
		}
	}
	for child := node.FirstChild; child != nil; child = child.NextSibling {
		switch child.Type {
		case html.ElementNode:
			handleElementNode(child)
		}
	}
}

func isLinkDiv(class string) bool {
	return strings.Contains(class, "egMi0") && strings.Contains(class, "kCrYT")
}

func hasURL(node *html.Node) bool {
	return node.Data == "a"
}

func getAttribute(node *html.Node, key string) string {
	for _, attr := range node.Attr {
		if attr.Key == key {
			return attr.Val
		}
	}
	return ""
}
