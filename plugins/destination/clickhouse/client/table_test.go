package client

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSplitKeyExpression(t *testing.T) {
	cases := []struct {
		name string
		key  string
		want []string
	}{
		{name: "empty", key: "", want: []string{}},
		{name: "empty tuple", key: "()", want: []string{}},
		{name: "single column", key: "_cq_id", want: []string{"_cq_id"}},
		{name: "single column in parentheses", key: "(_cq_id)", want: []string{"_cq_id"}},
		{name: "multiple columns", key: "_cq_source_name, _cq_id", want: []string{"_cq_source_name", "_cq_id"}},
		{name: "multiple columns in parentheses", key: "(_cq_source_name, _cq_id)", want: []string{"_cq_source_name", "_cq_id"}},
		{name: "function call", key: "toYYYYMM(_cq_sync_time)", want: []string{"toYYYYMM(_cq_sync_time)"}},
		{name: "function call in parentheses", key: "(toYYYYMM(_cq_sync_time))", want: []string{"toYYYYMM(_cq_sync_time)"}},
		{name: "separate parenthesized expressions", key: "(a) + (b)", want: []string{"(a) + (b)"}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			require.Equal(t, tc.want, splitKeyExpression(tc.key))
		})
	}
}
