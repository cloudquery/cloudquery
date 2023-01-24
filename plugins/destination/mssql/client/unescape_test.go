package client

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUnescape(t *testing.T) {
	type testCase struct {
		have, want string
	}
	for _, tc := range []testCase{
		{
			have: `{"BillTo":"P\u0026T"}`,
			want: `{"BillTo":"P&T"}`,
		},
		{
			have: `{"BillTo":"P&T"}`,
			want: `{"BillTo":"P&T"}`,
		},
	} {
		require.Equal(t, tc.want, unescape(tc.have))
	}
}
