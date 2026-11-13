package client

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidateTableNameAcceptsBigQueryTableIDs(t *testing.T) {
	names := []string{
		"aws_ec2_instances",
		"table1",
		"_cq_internal",
		"MixedCase",
		"with-dash",
		"with space",
		"tablé_名前",
		strings.Repeat("a", maxTableNameBytes),
	}
	for _, name := range names {
		t.Run(name, func(t *testing.T) {
			require.NoError(t, validateTableName(name))
		})
	}
}

func TestValidateTableNameRejectsUnsafeNames(t *testing.T) {
	names := map[string]string{
		"empty":             "",
		"too long":          strings.Repeat("a", maxTableNameBytes+1),
		"path traversal":    "../../datasets/other/tables/secret",
		"slash":             "a/b",
		"backslash":         `a\b`,
		"dot":               "dataset.table",
		"backtick":          "a`b",
		"question mark":     "a?b",
		"hash":              "a#b",
		"percent":           "a%2e%2e",
		"colon":             "project:table",
		"wildcard":          "table_*",
		"newline":           "a\nb",
		"nul":               "a\x00b",
		"invalid utf8":      "a\xffb",
		"unicode separator": "a\u0085b",
	}
	for description, name := range names {
		t.Run(description, func(t *testing.T) {
			err := validateTableName(name)
			require.ErrorContains(t, err, "invalid table name")
		})
	}
}

func TestValidateTableNameErrorNamesTheTable(t *testing.T) {
	require.ErrorContains(t, validateTableName("bad/name"), `"bad/name"`)
}

func TestValidateUnquotedColumnName(t *testing.T) {
	for _, name := range []string{"_cq_id", "name", "Column1"} {
		require.NoError(t, validateUnquotedColumnName(name), name)
	}
	for _, name := range []string{"", "1abc", "a b", "a`b", "a, (SELECT 1)", "a -- b", "a.b"} {
		require.Error(t, validateUnquotedColumnName(name), name)
	}
}

func TestValidateQuotedIdentifier(t *testing.T) {
	for _, value := range []string{"my-project", "example.com:my-project", "my_dataset"} {
		require.NoError(t, validateQuotedIdentifier("project_id", value), value)
	}
	for _, value := range []string{"", "a`b", `a\b`, "a\nb"} {
		require.ErrorContains(t, validateQuotedIdentifier("project_id", value), "invalid project_id", value)
	}
}
