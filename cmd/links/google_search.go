package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	search "github.com/philhanna/google_search"
)
const USAGE = `usage: google-search [QUERY]

Performs a Google search with the specified query. Returns a JSON array of the links found.

positional arguments:
  query          The query to be performed

`
func main() {
	flag.Usage = func() {
		fmt.Fprint(os.Stderr, USAGE)
	}
	flag.Parse()
	if flag.NArg() != 1 {
		log.Fatal("No query specified")
	}
	query := flag.Arg(0)
	doc, err := search.Download(query)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println ("[")
	for i, link := range doc.Links {
		comma := ","
		if i == len(doc.Links) - 1 {
			comma = ""
		}
		fmt.Println("  {")
		fmt.Printf("    %q:%q,\n", "title", link.Title)
		fmt.Printf("    %q:%q\n", "url", link.URL)
		fmt.Printf("  }%s\n", comma)
	}
	fmt.Println ("]")
}
