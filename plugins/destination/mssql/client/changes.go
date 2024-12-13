package client

import (
	"slices"
	"strings"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

func prettifyChanges(name string, changes []schema.TableColumnChange) string {
	builder := new(strings.Builder)
	builder.WriteString(name + ":")
	for _, change := range changes {
		builder.WriteString("\n")
		builder.WriteString(change.String())
	}
	return builder.String()
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
		// allow changing string to large string without a table drop
		if change.Previous.Type == arrow.BinaryTypes.String && change.Current.Type == arrow.BinaryTypes.LargeString {
			return false
		}
		return true
	default:
		return true
	}
}
