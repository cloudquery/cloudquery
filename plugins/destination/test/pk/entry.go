package pk

import (
	"github.com/cloudquery/plugin-sdk/schema"
)

type entry struct {
	reported bool
	keys     map[string]struct{}
}

// isDuplicate checks:
// 1. entry already reported -> nop
// 2. new resource has duplicate pk -> report
// 3. record the pk
func (e *entry) isDuplicate(table *schema.Table, resource []any) bool {
	if e.reported {
		return false
	}

	pk := Convert(table, resource)
	if _, ok := e.keys[pk]; ok {
		e.reported = true
		return true
	}

	e.keys[pk] = struct{}{}
	return false
}
