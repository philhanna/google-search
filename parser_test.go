package search

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	tests := []struct {
		name    string
		html    string
		want    []Link
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			want := tt.want
			have, err := Parse(tt.html)
			if tt.wantErr {
				assert.NotNil(t, err)
				return
			}
			assert.Equal(t, len(want), len(have))
			for i := 0; i < len(want); i++ {
				assert.Equal(t, want[i], have[i])
			}
		})
	}
}
