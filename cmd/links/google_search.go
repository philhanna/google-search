package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"

	search "github.com/philhanna/google-search"
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
	doc, err := search.Run(query)
	if err != nil {
		log.Fatal(err)
	}
	data, _ := json.MarshalIndent(doc.Links, "", "  ")
	fmt.Println(string(data))
}
