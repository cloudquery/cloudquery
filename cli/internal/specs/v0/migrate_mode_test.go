package specs

import (
	"testing"
)

func TestMigrateModeFromString(t *testing.T) {
	var migrateMode MigrateMode
	if err := migrateMode.UnmarshalJSON([]byte(`"forced"`)); err != nil {
		t.Fatal(err)
	}
	if migrateMode != MigrateModeForced {
		t.Fatalf("expected MigrateModeForced, got %v", migrateMode)
	}
	if err := migrateMode.UnmarshalJSON([]byte(`"safe"`)); err != nil {
		t.Fatal(err)
	}
	if migrateMode != MigrateModeSafe {
		t.Fatalf("expected MigrateModeSafe, got %v", migrateMode)
	}
}

func TestMigrateMode(t *testing.T) {
	for _, migrateModeStr := range migrateModeStrings {
		migrateMode, err := MigrateModeFromString(migrateModeStr)
		if err != nil {
			t.Fatal(err)
		}
		if migrateModeStr != migrateMode.String() {
			t.Fatalf("expected:%s got:%s", migrateMode, migrateMode.String())
		}
	}
}
