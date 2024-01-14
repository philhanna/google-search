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

func mockWikiNode() *html.Node{

	// Create elements (without yet links to others)
	elemRoot := &html.Node{}
	elemA := &html.Node{
		Type: html.ElementNode,
		Data: `a`,
		Attr: []html.Attribute{
			{
				Key: `href`,
				Val: `/url?q=https://en.wikipedia.org/wiki/Test-driven_development&amp;sa=U`,
			},
			{
				Key: `data-ved`,
				Val: `2ahUKEwjOvI2Xid2DAxV8M1kFHczQAXwQFnoECAoQAg`,
			},
		},
	}
	elemDiv1 := &html.Node{
		Type: html.ElementNode,
		Data: `div`,
		Attr: []html.Attribute{
			{
				Key: `class`,
				Val: `DnJfK`,
			},
		},
	}
	elemDiv2 := &html.Node{
		Type: html.ElementNode,
		Data: `div`,
		Attr: []html.Attribute{
			{
				Key: `class`,
				Val: `j039Wc`,
			},
		},
	}
	elemDiv3 := &html.Node{
		Type: html.ElementNode,
		Data: `div`,
		Attr: []html.Attribute{
			{
				Key: `class`,
				Val: `BNeawe vvjwJb AP7Wnd`,
			},
		},
	}
	elemH3 := &html.Node{
		Type: html.ElementNode,
		Data: `h3`,
		Attr: []html.Attribute{
			{
				Key: `class`,
				Val: `zBAuLc l97dzf`,
			},
		},
	}

	elemRoot.FirstChild = elemA
	elemRoot.LastChild = elemA

	elemA.Parent = elemRoot
	elemA.FirstChild = elemDiv1
	elemA.LastChild = elemDiv1

	elemDiv1.Parent = elemA
	elemDiv1.FirstChild = elemDiv2
	elemDiv2.LastChild = elemDiv2

	elemDiv2.Parent = elemDiv1
	elemDiv2.FirstChild = elemDiv3
	elemDiv2.LastChild = elemDiv3

	elemDiv3.Parent = elemDiv2
	elemDiv3.FirstChild = elemH3
	elemDiv3.LastChild = elemH3

	return elemRoot

}

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
		wantErr bool
	}{
		{
			name: "test driven development",
			html: getTestHTML("tdd.html"),
			want: []Link{
				{
					`https://en.wikipedia.org/wiki/Test-driven_development`,
					`Test-driven development - Wikipedia`,
				},
				{
					`https://www.browserstack.com/guide/what-is-test-driven-development`,
					`What is Test Driven Development (TDD)`,
				},
				{
					`https://martinfowler.com/bliki/TestDrivenDevelopment.html`,
					`Test Driven Development - Martin Fowler`,
				},
				{
					`https://www.agilealliance.org/glossary/tdd/`,
					`What is Test Driven Development (TDD)?`,
				},
				{
					`https://www.spiceworks.com/tech/devops/articles/what-is-tdd/`,
					`What is TDD (Test Driven Development)?`,
				},
				{
					`https://semaphoreci.com/blog/test-driven-development`,
					`Test-Driven Development: A Time-Tested Recipe for Quality ...`,
				},
				{
					`https://www.techtarget.com/searchsoftwarequality/definition/test-driven-development`,
					`What is test-driven development (TDD)? | Definition from ...`,
				},
				{
					`https://www.geeksforgeeks.org/test-driven-development-tdd/`,
					`Test Driven Development (TDD)`,
				},
				{
					`https://www.guru99.com/test-driven-development.html`,
					`What is Test Driven Development (TDD)? Example`,
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
			doc, err := NewHTMLDoc(tt.html)
			if tt.wantErr {
				assert.NotNil(t, err)
				return
			}
			want := tt.want
			have := doc.Links
			assert.Greater(t, len(have), 0)
			for i := 0; i < len(have); i++ {
				assert.Equal(t, want[i], have[i])
			}
		})
	}
}

func Test_isLinkDiv(t *testing.T) {
	tests := []struct {
		name  string
		class string
		want  bool
	}{
		{
			name: "empty",
		},
		{
			name:  `good`,
			class: `<div class="egMi0 kCrYT">`,
			want:  true,
		},
		{
			name:  `bad`,
			class: `<div>`,
		},
		{
			name:  `partial`,
			class: `<div class="kCrYT">`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			want := tt.want
			have := isLinkDiv(tt.class)
			assert.Equal(t, want, have)
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

func Test_getURL(t *testing.T) {
	tests := []struct {
		name string
		node *html.Node
		want string
	}{
		{
			name: "empty",
			node: &html.Node{},
		},
		{
			name: "wikipedia",
			node: mockWikiNode(),
			want: `/url?q=https://en.wikipedia.org/wiki/Test-driven_development&amp;sa=U`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			want := tt.want
			have := getURL(tt.node)
			assert.Equal(t, want, have)
		})
	}
}

func Test_getTitle(t *testing.T) {
	tests := []struct {
		name string
		node *html.Node
		want string
	}{
		{
			name: "empty",
			node: &html.Node{},
		},
		{
			name: "wikipedia",
			node: mockWikiNode(),
			want: `Test-driven development - Wikipedia`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			want := tt.want
			have := getTitle(tt.node)
			assert.Equal(t, want, have)
		})
	}
}
