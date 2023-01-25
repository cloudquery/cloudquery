package pk

import (
	"github.com/cloudquery/plugin-sdk/schema"
)

type (
	Store map[string]*entry
)

// IsDuplicate will return true only on the 1st duplicate occurrence for a particular table.
func (s Store) IsDuplicate(table *schema.Table, resource []any) bool {
	e, ok := s[table.Name]
	if !ok {
		e = &entry{keys: make(map[string]struct{})}
		s[table.Name] = e
	}

	return e.isDuplicate(table, resource)
}

func NewStore() Store {
	return make(map[string]*entry)
}
