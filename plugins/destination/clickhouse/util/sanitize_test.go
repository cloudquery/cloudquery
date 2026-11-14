package util

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSanitizeIDEscapesBackticksAndBackslashes(t *testing.T) {
	require.Equal(t, "`db``name`", SanitizeID("db`name"))
	require.Equal(t, "`table\\\\name`", SanitizeID(`table\name`))
	require.Equal(t, "`db\\\\``name`", SanitizeID("db\\`name"))
}

func TestUnquoteIDDecodesEscapes(t *testing.T) {
	require.Equal(t, "db`name", UnquoteID("`db``name`"))
	require.Equal(t, `table\name`, UnquoteID("`table\\\\name`"))
	require.Equal(t, `db\`+"`"+`name`, UnquoteID("`db\\\\``name`"))
}

func TestSanitizeIDPreservesEmbeddedDoubleQuotes(t *testing.T) {
	require.Equal(t, "`db\"name`", SanitizeID(`db"name`))
}
