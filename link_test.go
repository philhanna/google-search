package search

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_makeLink(t *testing.T) {
	tests := []struct {
		name    string
		url     string
		title   string
		want    *Link
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Good",
			url:  `/url?q=https://en.wikipedia.org/wiki/Test-driven_development&amp;sa=U&amp;ved=2ahUKEwjOvI2Xid2DAxV8M1kFHczQAXwQFnoECAoQAg&amp;usg=AOvVaw2r6YYBeWCTs0Zv2VznWfbr`,
			title: `
			Test-driven development - Wikipedia
		`,
			want: &Link{
				URL:   `https://en.wikipedia.org/wiki/Test-driven_development`,
				Title: `Test-driven development - Wikipedia`,
			},
		},
		{
			name: "No sanitization required",
			url:  `http://www.philhanna.net`,
			title: `JSP 2.0 - The Complete Reference`,
			want: &Link{
				URL:   `http://www.philhanna.net`,
				Title: `JSP 2.0 - The Complete Reference`,
			},
		},
		{
			name:    "empty",
			want:    &Link{URL: "", Title: ""},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			linkPtr := makeLink(tt.url, tt.title)
			if tt.wantErr {
				assert.Nil(t, linkPtr)
				return
			}
			assert.NotNil(t, linkPtr)
			assert.Equal(t, *tt.want, *linkPtr)
		})
	}
}
