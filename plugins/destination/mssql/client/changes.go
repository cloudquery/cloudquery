package client

import (
	"strings"

	"github.com/cloudquery/plugin-sdk/v3/schema"
	"golang.org/x/exp/slices"
)

func prettifyChanges(byTable map[string]schema.FieldChanges) string {
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

func unsafeSchemaChanges(have, want schema.Schemas) map[string]schema.FieldChanges {
	result := make(map[string]schema.FieldChanges)
	for _, w := range want {
		current := have.SchemaByName(schema.TableName(w))
		if current == nil {
			continue
		}
		unsafe := unsafeChanges(schema.GetSchemaChanges(w, current))
		if len(unsafe) > 0 {
			result[schema.TableName(w)] = unsafe
		}
	}
	return result
}

func unsafeChanges(changes []schema.FieldChange) schema.FieldChanges {
	unsafe := make([]schema.FieldChange, 0, len(changes))
	for _, c := range changes {
		if needsTableDrop(c) {
			unsafe = append(unsafe, c)
		}
	}
	return slices.Clip(unsafe)
}

func needsTableDrop(change schema.FieldChange) bool {
	switch change.Type {
	case schema.TableColumnChangeTypeAdd:
		return !change.Current.Nullable || schema.IsPk(change.Current)
	case schema.TableColumnChangeTypeRemove:
		return !change.Previous.Nullable || schema.IsPk(change.Previous)
	case schema.TableColumnChangeTypeUpdate:
		return true
	default:
		return true
	}
}
