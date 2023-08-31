package client

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestListStr(t *testing.T) {
	cases := []struct {
		have string
		want []string
	}{
		{
			have: `[
  "2023-04-24",
  null,
  "2023-04-24"
]`,
			want: []string{"2023-04-24", "null", "2023-04-24"},
		},
	}

	for _, tc := range cases {
		require.Equal(t, tc.want, snowflakeStrToArray(tc.have))
	}
}
