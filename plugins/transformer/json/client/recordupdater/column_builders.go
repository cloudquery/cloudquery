package recordupdater

import (
	"fmt"

	"github.com/apache/arrow/go/v17/arrow"
	"github.com/cloudquery/cloudquery/plugins/transformer/json/client/schemaupdater"
	"github.com/cloudquery/plugin-sdk/v4/types"
)

// ColumnBuilder creates separate columns from ONE json column, based on the type schema.
type ColumnBuilder interface {
	AddRow(row map[string]any)
	Build(key string) (arrow.Array, error)
}

type ColumnBuilders struct {
	builders []ColumnBuilder
}

func NewColumnBuilders(typeSchema map[string]string, originalColumn *types.JSONArray) (*ColumnBuilders, error) {
	b := &ColumnBuilders{
		builders: []ColumnBuilder{
			NewInt64ColumnsBuilder(typeSchema, originalColumn),
			NewUTF8ColumnsBuilder(typeSchema, originalColumn),
			NewTimestampColumnsBuilder(typeSchema, originalColumn),
			NewBoolColumnsBuilder(typeSchema, originalColumn),
			NewFloat64ColumnsBuilder(typeSchema, originalColumn),
			NewJSONColumnsBuilder(typeSchema, originalColumn),
		},
	}

	if err := b.requireNoUnknownTypes(typeSchema); err != nil {
		return nil, err
	}
	return b, nil
}

func (b *ColumnBuilders) AddRow(row map[string]any) {
	for _, builder := range b.builders {
		builder.AddRow(row)
	}
}

func (b *ColumnBuilders) Build(key string) (arrow.Array, error) {
	for _, builder := range b.builders {
		col, err := builder.Build(key)
		if err != nil {
			return nil, err
		}
		if col != nil {
			return col, nil
		}
	}
	return nil, fmt.Errorf("column %s not found", key)
}

func (*ColumnBuilders) requireNoUnknownTypes(typeSchema map[string]string) error {
	for key, typ := range typeSchema {
		if typ != schemaupdater.Int64Type &&
			typ != schemaupdater.Float64Type &&
			typ != schemaupdater.JSONType &&
			typ != schemaupdater.UTF8Type &&
			typ != schemaupdater.TimestampType &&
			typ != schemaupdater.BoolType {
			return fmt.Errorf("unsupported type for column [%s]: [%s]", key, typ)
		}
	}
	return nil
}
