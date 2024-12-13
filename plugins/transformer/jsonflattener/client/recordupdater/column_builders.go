package recordupdater

import (
	"fmt"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/cloudquery/cloudquery/plugins/transformer/jsonflattener/client/schemaupdater"
	"github.com/cloudquery/plugin-sdk/v4/caser"
	"github.com/cloudquery/plugin-sdk/v4/types"
)

// columnBuilder creates separate columns from ONE json column, based on the type schema.
type columnBuilder interface {
	addRow(row map[string]any)
	build(key string) (arrow.Array, error)
}

type columnBuilders struct {
	tableName           string
	colName             string
	typeSchema          map[string]string
	preprocessRowKeysFn func(row map[string]any) map[string]any
	builders            []columnBuilder
	caser               *caser.Caser
}

func newColumnBuilders(tableName string, colName string, typeSchema map[string]string, originalColumn *types.JSONArray) (*columnBuilders, error) {
	b := &columnBuilders{
		tableName:  tableName,
		colName:    colName,
		typeSchema: typeSchema,
		caser:      caser.New(),
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
	row = b.preprocessRowKeys(row)
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

func (b *columnBuilders) preprocessRowKeys(row map[string]any) map[string]any {
	if b.preprocessRowKeysFn == nil {
		var isPerfectMatch bool
		b.preprocessRowKeysFn, isPerfectMatch = b.choosePreprocessRowKeysFn(row)

		// If the match is not perfect, choose again on the next row.
		if !isPerfectMatch {
			defer func() { b.preprocessRowKeysFn = nil }()
		}
	}
	return b.preprocessRowKeysFn(row)
}

// choosePreprocessRowKeysFn tries different key preprocess functions try to match the typeSchema.
// It returns true if the match was perfect.
func (b *columnBuilders) choosePreprocessRowKeysFn(row map[string]any) (func(map[string]any) map[string]any, bool) {
	preprocessFns := []struct {
		fn      func(map[string]any) map[string]any
		matches int
	}{
		{b.preprocessRowKeysFnIdentity, 0},
		{b.preprocessRowKaysFnSnakeCase, 0},
	}

	// Run each function and count key matches
	for i := range preprocessFns {
		processedRow := preprocessFns[i].fn(row)
		for key := range b.typeSchema {
			if _, ok := processedRow[key]; ok {
				preprocessFns[i].matches++
			}
		}
	}

	// Find function with most key matches
	maxMatches := -1
	var bestFn func(map[string]any) map[string]any
	for _, p := range preprocessFns {
		if p.matches > maxMatches {
			maxMatches = p.matches
			bestFn = p.fn
		}
	}

	return bestFn, maxMatches == len(b.typeSchema)
}

func (*columnBuilders) preprocessRowKeysFnIdentity(row map[string]any) map[string]any {
	return row
}

func (b *columnBuilders) preprocessRowKaysFnSnakeCase(row map[string]any) map[string]any {
	newRow := make(map[string]any)
	for key, value := range row {
		newRow[b.caser.ToSnake(key)] = value
	}
	return newRow
}
