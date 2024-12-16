package queries

import (
	"testing"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

func TestRead(t *testing.T) {
	ensureContents(t, Read(&schema.Table{
		Name: "table_name",
		Columns: schema.ColumnList{
			schema.CqIDColumn,
			schema.CqParentIDColumn,
			schema.CqSourceNameColumn,
			schema.CqSyncTimeColumn,
			schema.Column{Name: "extra_col", Type: arrow.PrimitiveTypes.Float64},
		},
	}), "read.sql")
}
