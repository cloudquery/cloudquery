package history

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDSNElement(t *testing.T) {
	tbl := []struct {
		input    string
		mod      map[string]string
		expected string
	}{
		{
			input:    "postgres://a:b@c.d?x=y&z=f",
			mod:      map[string]string{"ADD": "THIS"},
			expected: "postgres://a:b@c.d?x=y&z=f&ADD=THIS",
		},
		{
			input:    "host=localhost user=postgres password=pass database=postgres port=5432 sslmode=disable",
			mod:      map[string]string{"ADD": "THIS"},
			expected: "postgres://postgres:pass@localhost:5432/postgres?ADD=THIS&sslmode=disable",
		},
		{
			input:    "tsdb://a:b@c.d?x=y&z=f",
			mod:      map[string]string{"ADD": "THIS"},
			expected: "tsdb://a:b@c.d?x=y&z=f&ADD=THIS",
		},
	}
	for _, tc := range tbl {
		out, err := setDsnElement(tc.input, tc.mod)
		assert.NoError(t, err)
		u1, err := url.Parse(tc.expected)
		assert.NoError(t, err)
		u2, err := url.Parse(out)
		assert.NoError(t, err)
		assert.EqualValues(t, u1.Scheme, u2.Scheme)
		assert.EqualValues(t, u1.Host, u2.Host)
		assert.EqualValues(t, u1.Path, u2.Path)
		assert.EqualValues(t, u1.Query(), u2.Query())
	}
}
