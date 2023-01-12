package queries

import (
	"github.com/cloudquery/plugin-sdk/schema"
)

type TableDefinition struct {
	Name    string
	Columns ColumnDefinitions
}

func (t *TableDefinition) GetAddedColumns(base *TableDefinition) ColumnDefinitions {
	var res ColumnDefinitions

	for _, col := range t.Columns {
		if base.Columns.Get(col.Name) == nil {
			res = append(res, col)
		}
	}

	return res
}

func (t *TableDefinition) GetChangedColumns(base *TableDefinition) ColumnDefinitions {
	var res ColumnDefinitions

	for _, col := range t.Columns {
		other := base.Columns.Get(col.Name)
		if other == nil || other.Type == col.Type {
			continue
		}

		res = append(res, col)
	}

	return res
}

type TableDefinitions map[string]*TableDefinition

// GetTableDefinitions works with flattened table list (so scans only the top level).
func GetTableDefinitions(flattened schema.Tables) TableDefinitions {
	defs := make(TableDefinitions, len(flattened))

	for _, tbl := range flattened {
		defs[tbl.Name] = getTableDefinition(tbl)
	}

	return defs
}

func getTableDefinition(table *schema.Table) *TableDefinition {
	columns := make(ColumnDefinitions, len(table.Columns))

	for i, col := range table.Columns {
		if col.Name == schema.CqIDColumn.Name {
			// _cq_id has to be marked not null
			col.CreationOptions.NotNull = true
		}
		columns[i] = getColumnDefinition(&col)
	}

	return &TableDefinition{Name: table.Name, Columns: columns}
}
