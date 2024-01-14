package search

import "os"

var DocumentRoot string

func init() {
	DocumentRoot, _ = os.Getwd()
}