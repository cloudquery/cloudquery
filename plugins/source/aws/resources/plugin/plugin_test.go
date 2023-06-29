package plugin

import (
	"testing"
)

func TestAWS(t *testing.T) {
	// Note: this test is simple, but serves as a smoke test.
	// The AWS() call below also catches duplicate columns and other issues
	// that may have been missed if mock tests are incomplete.
	p := AWS()
	name := p.Name()
	if name != "aws" {
		t.Errorf("Name() = %q, want %q", name, "aws")
	}
}

// This test ensures that all tables have a unique description.
func TestAWSTableDescriptions(t *testing.T) {
	descriptions := make(map[string]bool)
	p := AWS()
	tables := p.Tables()
	for _, table := range tables {
		if table.Description == "" {
			t.Errorf("Table %q has no description", table.Name)
		}
		if _, ok := descriptions[table.Description]; ok {
			t.Fatalf("description for table exists %s", table.Name)
		} else {
			descriptions[table.Description] = true
		}
	}
}
