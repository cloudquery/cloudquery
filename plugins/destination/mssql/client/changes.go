package client

import (
	"strings"

	"github.com/cloudquery/plugin-sdk/v3/schema"
	"golang.org/x/exp/slices"
)

func prettifyChanges(byTable map[string][]schema.TableColumnChange) string {
	builder := new(strings.Builder)
	for name, changes := range byTable {
		builder.WriteString(name + ":")
		for _, change := range changes {
			builder.WriteString("\n")
			builder.WriteString(change.String())
		}
	}
	return builder.String()
}

func unsafeSchemaChanges(have, want schema.Tables) map[string][]schema.TableColumnChange {
	result := make(map[string][]schema.TableColumnChange)
	for _, w := range want {
		current := have.Get(w.Name)
		if current == nil {
			continue
		}
		unsafe := unsafeChanges(w.GetChanges(current))
		if len(unsafe) > 0 {
			result[w.Name] = unsafe
		}
	}
	return result
}

func unsafeChanges(changes []schema.TableColumnChange) []schema.TableColumnChange {
	unsafe := make([]schema.TableColumnChange, 0, len(changes))
	for _, c := range changes {
		if needsTableDrop(c) {
			unsafe = append(unsafe, c)
		}
	}
	return slices.Clip(unsafe)
}

func needsTableDrop(change schema.TableColumnChange) bool {
	switch change.Type {
	case schema.TableColumnChangeTypeAdd:
		return change.Current.NotNull || change.Current.PrimaryKey
	case schema.TableColumnChangeTypeRemove:
		return change.Previous.NotNull || change.Previous.PrimaryKey
	case schema.TableColumnChangeTypeUpdate:
		return true
	default:
		return true
	}
}
