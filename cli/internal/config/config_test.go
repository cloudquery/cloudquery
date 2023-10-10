package config

import (
	"testing"

	"github.com/adrg/xdg"
)

func TestSetGetValue(t *testing.T) {
	t.Setenv("XDG_CONFIG_HOME", t.TempDir())
	xdg.Reload()

	err := SetValue("team", "my-team")
	if err != nil {
		t.Fatal(err)
	}

	got, err := GetValue("team")
	if err != nil {
		t.Fatal(err)
	}
	if got != "my-team" {
		t.Fatalf("expected %q, got %q", "my-team", got)
	}

	err = UnsetValue("team")
	if err != nil {
		t.Fatal(err)
	}

	got, err = GetValue("team")
	if err != nil {
		t.Fatal(err)
	}
	if got != "" {
		t.Fatalf("expected %q, got %q", "", got)
	}
}
