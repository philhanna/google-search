package search

import (
	"io"
	"log"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/net/html"
)

var testdataCache = make(map[string]string)

func getTestHTML(filename string) string {
	if data, OK := testdataCache[filename]; OK {
		return data
	}
	path := filepath.Join("testdata", filename)
	fp, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer fp.Close()
	body, err := io.ReadAll(fp)
	if err != nil {
		log.Fatal(err)
	}
	data := string(body)
	testdataCache[filename] = data
	return data
}

func TestParse(t *testing.T) {
	tests := []struct {
		name    string
		html    string
		want    []Link
	}{
		{
			name: "test driven development",
			html: getTestHTML("a.html"),
			want: []Link{
				{
					`https://en.wikipedia.org/wiki/Test-driven_development`,
					`Test-driven development - Wikipedia`,
				},
				{
					`https://www.browserstack.com/guide/what-is-test-driven-development`,
					`What is Test Driven Development (TDD) ? | BrowserStack`,
				},
				{
					`https://martinfowler.com/bliki/TestDrivenDevelopment.html`,
					`Test Driven Development - Martin Fowler`,
				},
				{
					`https://semaphoreci.com/blog/test-driven-development`,
					`Test-Driven Development: A Time-Tested Recipe for Quality Software`,
				},
				{
					`https://www.geeksforgeeks.org/test-driven-development-tdd/`,
					`Test Driven Development (TDD) - GeeksforGeeks`,
				},
				{
					`https://www.guru99.com/test-driven-development.html`,
					`What is Test Driven Development (TDD)? Example - Guru99`,
				},
				{
					`https://www.agilealliance.org/glossary/tdd/`,
					`What is Test Driven Development (TDD)? - Agile Alliance`,
				},
				{
					`https://www.spiceworks.com/tech/devops/articles/what-is-tdd/`,
					`What is TDD (Test Driven Development)? - Spiceworks`,
				},
				{
					`https://www.techtarget.com/searchsoftwarequality/definition/test-driven-development`,
					`What is test-driven development (TDD)? | Definition from TechTarget`,
				},
				{
					`https://www.amazon.com/Test-Driven-Development-Kent-Beck/dp/0321146530`,
					`Test Driven Development: By Example: Beck, Kent`,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			doc := NewHTMLDoc(tt.html)
			want := tt.want
			have := doc.Links
			assert.Greater(t, len(have), 0)
			for i := 0; i < len(have); i++ {
				assert.Equal(t, want[i], have[i])
			}
		})
	}
}

func Test_getAttribute(t *testing.T) {
	tests := []struct {
		name string
		node *html.Node
		key  string
		want string
	}{
		{
			name: "empty",
			node: &html.Node{},
		},
		{
			name: "Good",
			node: &html.Node{
				Data: "a",
				Attr: []html.Attribute{
					{Key: "href", Val: "foo"},
				},
			},
			key:  "href",
			want: "foo",
		},
		{
			name: "2nd attribute",
			node: &html.Node{
				Data: "a",
				Attr: []html.Attribute{
					{Key: "class", Val: "something"},
					{Key: "href", Val: "foo"},
				},
			},
			key:  "href",
			want: "foo",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			want := tt.want
			have := getAttribute(tt.node, tt.key)
			assert.Equal(t, want, have)
		})
	}
}

func TestNewHTMLDoc(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    *HTMLDoc
	}{
		{
			name: "empty",
		},
		{
			name:  "Good",
			input: getTestHTML("a.html"),
		},
		{
			name:  "Bogus",
			input: getTestHTML("bogus.html"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			doc := NewHTMLDoc(tt.input)
			assert.NotNil(t, doc)
		})
	}
}
