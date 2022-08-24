package codegen

import "testing"

func TestAll(t *testing.T) {
	resources := All()
	for _, r := range resources {
		foundID := false
		for _, c := range r.Table.Columns {
			if c.Name == "id" {
				foundID = true
				want := `PathResolver("ID")`
				if c.Resolver != want {
					t.Errorf("table %v: id column resolver = %q, want %q", c.Name, c.Resolver, want)
				}
			}
		}
		if !foundID {
			t.Errorf("did not find id column for table %v", r.Table.Name)
		}
	}
}
