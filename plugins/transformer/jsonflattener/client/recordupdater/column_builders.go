package recordupdater

import (
	"fmt"

	"github.com/apache/arrow/go/v17/arrow"
	"github.com/cloudquery/cloudquery/plugins/transformer/jsonflattener/client/schemaupdater"
	"github.com/cloudquery/plugin-sdk/v4/types"
)

// columnBuilder creates separate columns from ONE json column, based on the type schema.
type columnBuilder interface {
	addRow(row map[string]any)
	build(key string) (arrow.Array, error)
}

type columnBuilders struct {
	tableName string
	colName   string
	builders  []columnBuilder
}

func newColumnBuilders(tableName string, colName string, typeSchema map[string]string, originalColumn *types.JSONArray) (*columnBuilders, error) {
	b := &columnBuilders{
		tableName: tableName,
		colName:   colName,
		builders: []columnBuilder{
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

func (b *columnBuilders) addRow(row map[string]any) {
	for _, builder := range b.builders {
		builder.addRow(row)
	}
}

func (b *columnBuilders) build(key string) (arrow.Array, error) {
	for _, builder := range b.builders {
		col, err := builder.build(key)
		if err != nil {
			return nil, err
		}
		if col != nil {
			return col, nil
		}
	}
	return nil, fmt.Errorf("column %s not found", key)
}

func (b *columnBuilders) requireNoUnknownTypes(typeSchema map[string]string) error {
	for key, typ := range typeSchema {
		if typ != schemaupdater.Int64Type &&
			typ != schemaupdater.Float64Type &&
			typ != schemaupdater.JSONType &&
			typ != schemaupdater.UTF8Type &&
			typ != schemaupdater.TimestampType &&
			typ != schemaupdater.BoolType {
			return fmt.Errorf("unsupported type for column [%s] on original column [%s.%s]: [%s]", key, b.tableName, b.colName, typ)
		}
	}
	return nil
}
