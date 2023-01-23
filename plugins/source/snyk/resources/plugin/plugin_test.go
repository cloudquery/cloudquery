package plugin

import (
	"testing"
)

func TestSnyk(t *testing.T) {
	// Note: this test is simple, but serves as a smoke test.
	// The Snyk() call below also catches duplicate columns and other issues
	// that may have been missed if mock tests are incomplete.
	p := Snyk()
	name := p.Name()
	if name != "snyk" {
		t.Errorf("Name() = %q, want %q", name, "snyk")
	}
}
