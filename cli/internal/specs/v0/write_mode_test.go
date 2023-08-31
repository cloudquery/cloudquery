package specs

import (
	"testing"
)

func TestWriteModeFromString(t *testing.T) {
	var writeMode WriteMode
	if err := writeMode.UnmarshalJSON([]byte(`"append"`)); err != nil {
		t.Fatal(err)
	}
	if writeMode != WriteModeAppend {
		t.Fatalf("expected WriteModeAppend, got %v", writeMode)
	}
	if err := writeMode.UnmarshalJSON([]byte(`"overwrite"`)); err != nil {
		t.Fatal(err)
	}
	if writeMode != WriteModeOverwrite {
		t.Fatalf("expected WriteModeOverwrite, got %v", writeMode)
	}
}

func TestWriteMode(t *testing.T) {
	for _, writeModeStr := range writeModeStrings {
		writeMode, err := WriteModeFromString(writeModeStr)
		if err != nil {
			t.Fatal(err)
		}
		if writeModeStr != writeMode.String() {
			t.Fatalf("expected:%s got:%s", writeModeStr, writeMode.String())
		}
	}
}
