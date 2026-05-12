package cmd

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidateSpecAgainstSchema_LenientForBuggyPlugins(t *testing.T) {
	// Empty schema (e.g. plugin returned Unimplemented over gRPC) is treated as skip.
	require.NoError(t, validateSpecAgainstSchema("", map[string]any{}))
	// Unparseable schema is logged and skipped (lenient path used for plugin gRPC results).
	require.NoError(t, validateSpecAgainstSchema(`{not-valid-json`, map[string]any{}))
}

func TestValidateSpecAgainstSchemaStrict_FailsOnBadSchemaAndBadSpec(t *testing.T) {
	const goodSchema = `{
		"$schema": "https://json-schema.org/draft/2020-12/schema",
		"type": "object",
		"properties": {"field": {"type": "string"}},
		"required": ["field"],
		"additionalProperties": false
	}`

	// Good spec passes.
	require.NoError(t, validateSpecAgainstSchemaStrict(goodSchema, map[string]any{"field": "ok"}))

	// Bad spec is rejected.
	err := validateSpecAgainstSchemaStrict(goodSchema, map[string]any{"field": 42})
	require.Error(t, err)

	// Corrupt schema is rejected, NOT silently passed.
	err = validateSpecAgainstSchemaStrict(`{not-valid-json`, map[string]any{})
	require.Error(t, err)
	require.Contains(t, err.Error(), "failed to parse JSON schema")

	// Empty schema is rejected, NOT silently passed.
	err = validateSpecAgainstSchemaStrict("", map[string]any{})
	require.Error(t, err)
}
