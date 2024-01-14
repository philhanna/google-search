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
		Level: slog.LevelInfo,
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
	printNode(doc)
}

func printNode(node *html.Node) {
	slog.Info("printNode: Entry", "type", node.Type)
	switch node.Type {
	case html.DocumentNode:
		printDocumentNode(node)
	case html.ElementNode:
		printElementNode(node)
	case html.TextNode:
		printTextNode(node)
	}
	slog.Info("printNode: Exit")
}

func printDocumentNode(node *html.Node) {
	slog.Info("printDocumentNode: Entry", "data", node.Data)
	for child := node.FirstChild; child != nil; child = child.NextSibling {
		printNode(child)
	}
	slog.Info("printDocumentNode: Exit")
}

func printElementNode(node *html.Node) {
	slog.Info("printElementNode: Entry", "data", node.Data)
	startTag := "<"
	startTag += node.Data
	for _, attr := range node.Attr {
		startTag += fmt.Sprintf(" %s=%q", attr.Key, attr.Val)
	}
	startTag += ">"
	fmt.Println(startTag)
	if node.Data != "script" {
		for child := node.FirstChild; child != nil; child = child.NextSibling {
			printNode(child)
		}
	}
	fmt.Printf("</%s>\n", node.Data)
	slog.Info("printElementNode: Exit")

}

func printTextNode(node *html.Node) {
	slog.Info("printTextNode: Entry", "text", node.Data)
	fmt.Printf("%s\n", node.Data)
	slog.Info("printTextNode: Exit")
}
