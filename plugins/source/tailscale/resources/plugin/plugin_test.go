package plugin

import (
	"testing"
)

func TestTailscale(t *testing.T) {
	// Note: this test is simple, but serves as a smoke test.
	// The Tailscale() call below also catches duplicate columns and other issues
	// that may have been missed if mock tests are incomplete.
	p := Tailscale()
	name := p.Name()
	if name != "tailscale" {
		t.Errorf("Name() = %q, want %q", name, "tailscale")
	}
}
