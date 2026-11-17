package util

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSanitizeID(t *testing.T) {
	cases := []struct {
		name string
		id   string
		want string
	}{
		{name: "plain", id: "table_name", want: "`table_name`"},
		{name: "single quote", id: "table'name", want: "`table'name`"},
		{name: "double quote", id: `table"name`, want: "`table\"name`"},
		{name: "dot", id: "db.table", want: "`db.table`"},
		{name: "backtick", id: "x`.`y", want: "`x\\`.\\`y`"},
		{name: "backslash", id: `table\name`, want: "`table\\\\name`"},
		{name: "trailing backslash", id: `table\`, want: "`table\\\\`"},
		{name: "backslash before backtick", id: "x\\`y", want: "`x\\\\\\`y`"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := SanitizeID(tc.id)
			require.Equal(t, tc.want, got)
			require.Equal(t, tc.id, UnquoteID(got))
		})
	}
}

func TestSanitizeIDStripsNullByte(t *testing.T) {
	require.Equal(t, "`tablename`", SanitizeID("table\x00name"))
}

func TestUnquoteID(t *testing.T) {
	cases := []struct {
		name string
		id   string
		want string
	}{
		{name: "unquoted", id: "column_name", want: "column_name"},
		{name: "quoted", id: "`column_name`", want: "column_name"},
		{name: "escaped backtick", id: "`x\\`y`", want: "x`y"},
		{name: "escaped trailing backtick", id: "`x\\``", want: "x`"},
		{name: "escaped backslash", id: "`x\\\\y`", want: `x\y`},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			require.Equal(t, tc.want, UnquoteID(tc.id))
		})
	}
}

func TestSanitized(t *testing.T) {
	require.Equal(t, []string{"`a'b`", "`c\\`d`", "`e\\\\f`"}, Sanitized("a'b", "c`d", `e\f`))
}
