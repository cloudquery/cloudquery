package specs

import (
	"testing"

	"github.com/cloudquery/codegen/jsonschema"
	"github.com/stretchr/testify/require"
)

func TestMigrateModeFromString(t *testing.T) {
	m, err := MigrateModeFromString("safe")
	require.NoError(t, err)
	require.Equal(t, MigrateModeSafe, m)

	m, err = MigrateModeFromString("forced")
	require.NoError(t, err)
	require.Equal(t, MigrateModeForced, m)

	m, err = MigrateModeFromString("Forced")
	require.Error(t, err)
	require.Equal(t, MigrateModeSafe, m)

	m, err = MigrateModeFromString("")
	require.Error(t, err)
	require.Equal(t, MigrateModeSafe, m)
}

func TestMigrateMode_JSONSchemaExtend(t *testing.T) {
	data, err := jsonschema.Generate(new(MigrateMode))
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
			Name: "safe",
			Spec: `"safe"`,
		},
		{
			Name: "forced",
			Spec: `"forced"`,
		},
	})
}

func TestMigrateModeRoundTrip(t *testing.T) {
	for _, migrateModeStr := range AllMigrateModes {
		migrateMode, err := MigrateModeFromString(migrateModeStr)
		if err != nil {
			t.Fatal(err)
		}
		if migrateModeStr != migrateMode.String() {
			t.Fatalf("expected:%s got:%s", migrateMode, migrateMode.String())
		}
	}
}
