package queries

import (
	"github.com/cloudquery/plugin-sdk/schema"
)

type (
	ColumnDefinition struct {
		Name string
		Type string
	}
	ColumnDefinitions []*ColumnDefinition
)

func (c ColumnDefinitions) Get(name string) *ColumnDefinition {
	for _, col := range c {
		if col.Name == name {
			return col
		}
	}

	return nil
}

func getColumnDefinition(column *schema.Column) *ColumnDefinition {
	return &ColumnDefinition{Name: column.Name, Type: chType(column)}
}
