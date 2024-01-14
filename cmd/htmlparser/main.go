package main

import (
	"fmt"
	"log"
	"strings"

	"golang.org/x/net/html"
)

var level = 0

func main() {
	s := `<p>Links:</p><ul><li><a href="foo">Foo</a><li><a href="/bar/baz">BarBaz</a></ul>`
	doc, err := html.Parse(strings.NewReader(s))
	if err != nil {
		log.Fatal(err)
	}
	printDocumentNode(doc)
}

func indent() string {
	return strings.Repeat("  ", level)
}

func printNode(node *html.Node) {
	// fmt.Printf("node type: %d, data: %q\n", node.Type, node.Data)
	switch node.Type {
	case html.TextNode:
		printTextNode(node)
	case html.DocumentNode:
		printDocumentNode(node)
	case html.ElementNode:
		printElementNode(node)
	}
}

func printTextNode(node *html.Node) {
	fmt.Printf("%s%s\n", indent(), node.Data)
}

func printDocumentNode(node *html.Node) {
	for child := node.FirstChild; child != nil; child = child.NextSibling {
		printNode(child)
	}
}

func printElementNode(node *html.Node) {
	fmt.Printf("%s<%s>\n", indent(), node.Data)
	level++
	for child := node.FirstChild; child != nil; child = child.NextSibling {
		printNode(child)
	}
	level--
	fmt.Printf("%s</%s>\n", indent(), node.Data)
}