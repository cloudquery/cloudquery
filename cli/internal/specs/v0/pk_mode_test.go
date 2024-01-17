package specs

import (
	"testing"

	"github.com/cloudquery/codegen/jsonschema"
	"github.com/stretchr/testify/require"
)

func TestPKModeFromString(t *testing.T) {
	m, err := PKModeFromString("default")
	require.NoError(t, err)
	require.Equal(t, PKModeDefaultKeys, m)

	m, err = PKModeFromString("cq-id-only")
	require.NoError(t, err)
	require.Equal(t, PKModeCQID, m)

	m, err = PKModeFromString("Default")
	require.Error(t, err)
	require.Equal(t, PKModeDefaultKeys, m)

	m, err = PKModeFromString("")
	require.Error(t, err)
	require.Equal(t, PKModeDefaultKeys, m)
}

func TestPKMode_JSONSchemaExtend(t *testing.T) {
	data, err := jsonschema.Generate(new(PKMode))
	require.NoError(t, err)
	jsonschema.TestJSONSchema(t, string(data), []jsonschema.TestCase{
		{
			Name: "empty",
			Err:  true,
			Spec: `""`,
		},
		{
			Name: "null",
			Err:  true,
			Spec: `null`,
		},
		{
			Name: "bad type",
			Err:  true,
			Spec: `123`,
		},
		{
			Name: "bad value",
			Err:  true,
			Spec: `"extra"`,
		},
		{
			Name: "default",
			Spec: `"default"`,
		},
		{
			Name: "cq-id-only",
			Spec: `"cq-id-only"`,
		},
	})
}

func TestPKModeRoundTrip(t *testing.T) {
	for _, pkModeStr := range AllPKModes {
		pkMode, err := PKModeFromString(pkModeStr)
		if err != nil {
			t.Fatal(err)
		}
		if pkModeStr != pkMode.String() {
			t.Fatalf("expected:%s got:%s", pkModeStr, pkMode.String())
		}
	}
}

func TestPKModeMarshalJSON(t *testing.T) {
	pkMode := PKModeCQID
	if pkModeStr, err := pkMode.MarshalJSON(); err != nil {
		t.Fatal(err)
	} else if string(pkModeStr) != `"cq-id-only"` {
		t.Fatalf("expected:\"cq-id\" got:%s", string(pkModeStr))
	}

	pkMode = PKModeDefaultKeys
	if pkModeStr, err := pkMode.MarshalJSON(); err != nil {
		t.Fatal(err)
	} else if string(pkModeStr) != `"default"` {
		t.Fatalf("expected:\"cq-id\" got:%s", string(pkModeStr))
	}
}
