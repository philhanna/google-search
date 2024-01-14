package search

import (
	"io"
	"log"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
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
			want := tt.want
			have, err := Parse(tt.html)
			if tt.wantErr {
				assert.NotNil(t, err)
				return
			}
			assert.Greater(t, len(have), 0)
			for i := 0; i < len(have); i++ {
				assert.Equal(t, want[i], have[i])
			}
		})
	}
}
