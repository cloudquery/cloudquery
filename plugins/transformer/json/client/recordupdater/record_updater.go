package recordupdater

import (
	"encoding/json"
	"fmt"

	"github.com/apache/arrow/go/v17/arrow"
	"github.com/apache/arrow/go/v17/arrow/array"
	"github.com/cloudquery/cloudquery/plugins/transformer/json/client/schemaupdater"
	"github.com/cloudquery/cloudquery/plugins/transformer/json/client/util"
	"github.com/cloudquery/plugin-sdk/v4/caser"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/types"
)

// RecordUpdater takes an `arrow.Record` and knows how to make simple subsequent changes to it.
// It doesn't know which table it belongs to or if the changes make sense.
type RecordUpdater struct {
	record        arrow.Record
	schemaUpdater *schemaupdater.SchemaUpdater
	caser         *caser.Caser
	tableName     string
}

func New(record arrow.Record) *RecordUpdater {
	tableName, _ := record.Schema().Metadata().GetValue(schema.MetadataTableName)

	return &RecordUpdater{
		record:        record,
		schemaUpdater: schemaupdater.New(record.Schema()),
		caser:         caser.New(),
		tableName:     tableName,
	}
}

func (r *RecordUpdater) FlattenJSONFields() (arrow.Record, error) {
	fieldTypeSchemas := map[string]map[string]string{}

	for _, field := range r.record.Schema().Fields() {
		rawTypeSchema, ok := field.Metadata.ToMap()[schema.MetadataTypeSchema]
		if !ok || rawTypeSchema == "" {
			continue
		}
		var typeSchema map[string]string
		if err := json.Unmarshal([]byte(rawTypeSchema), &typeSchema); err != nil {
			// In this case it can be an array
			fmt.Println("failed to unmarshal type schema", rawTypeSchema)
			continue
		}
		typeSchema = preprocessTypeSchema(typeSchema)
		fieldTypeSchemas[field.Name] = typeSchema
	}

	if len(fieldTypeSchemas) == 0 {
		return r.record, nil
	}

	newColumns := make([]arrow.Array, 0, int(r.record.NumCols())+countNewColumns(fieldTypeSchemas))
	for i := 0; i < int(r.record.NumCols()); i++ {
		newColumns = append(newColumns, r.record.Column(i))
	}
	brandNewColumns, err := r.buildAllNewColumns(fieldTypeSchemas)
	if err != nil {
		return nil, err
	}
	newColumns = append(newColumns, brandNewColumns...)

	newSchema, err := r.schemaUpdater.AddJSONFlattenedFields(fieldTypeSchemas)
	if err != nil {
		return nil, err
	}
	r.record = array.NewRecord(newSchema, newColumns, r.record.NumRows())
	return r.record, nil
}

func countNewColumns(fieldTypeSchemas map[string]map[string]string) int {
	count := 0
	for _, typeSchema := range fieldTypeSchemas {
		count += len(typeSchema)
	}
	return count
}

func (r *RecordUpdater) buildAllNewColumns(fieldTypeSchemas map[string]map[string]string) ([]arrow.Array, error) {
	newColumns := make([]arrow.Array, 0, countNewColumns(fieldTypeSchemas))

	// N.B. new columns are ordered lexicographically by key, but we don't have a mapping between
	// the new column indices and these keys
	colNameToIndex := make(map[string]int)
	for i := 0; i < int(r.record.NumCols()); i++ {
		colNameToIndex[r.record.ColumnName(i)] = i
	}

	colNames := util.SortedKeys(fieldTypeSchemas)
	for _, colName := range colNames {
		cols, err := r.buildNewColumnsFromColumn(colName, r.record.Column(colNameToIndex[colName]), fieldTypeSchemas[colName])
		if err != nil {
			return nil, err
		}
		newColumns = append(newColumns, cols...)
	}

	return newColumns, nil
}

func (r *RecordUpdater) buildNewColumnsFromColumn(colName string, col arrow.Array, typeSchema map[string]string) ([]arrow.Array, error) {
	jsonArrayCol, ok := col.(*types.JSONArray)
	if !ok {
		return nil, fmt.Errorf("expected JSONArray column, got %T", col)
	}

	builders, err := NewColumnBuilders(r.tableName, colName, typeSchema, col.(*types.JSONArray))
	if err != nil {
		return nil, err
	}

	for i := 0; i < jsonArrayCol.Len(); i++ {
		row, err := r.preprocessRow(colName, i, jsonArrayCol.Value(i))
		if err != nil {
			return nil, err
		}
		builders.AddRow(row)
	}

	// Order of the new columns must match with the schema.
	// The ordering rule is: (1) by original column existing order, (2) by lexicographical order of subkeys
	newColumns := make([]arrow.Array, 0, len(typeSchema))
	keys := util.SortedKeys(typeSchema)
	for _, key := range keys {
		col, err := builders.Build(key)
		if err != nil {
			return nil, err
		}
		newColumns = append(newColumns, col)
	}

	return newColumns, nil
}

func (r *RecordUpdater) preprocessRow(colName string, rowIndex int, rawRow any) (map[string]any, error) {
	var row map[string]any

	switch typedRawRow := rawRow.(type) {
	case *types.JSONType:
		bs, err := typedRawRow.MarshalJSON()
		if err != nil {
			return nil, fmt.Errorf("failed to marshal JSON at row %d: %v", rowIndex, err)
		}
		if err := json.Unmarshal(bs, &row); err != nil {
			return nil, fmt.Errorf("failed to unmarshal JSON at row %d: %v", rowIndex, err)
		}
	case map[string]any:
		row = typedRawRow
	case nil:
		row = map[string]any{}
	default:
		return nil, fmt.Errorf("expected JSON column for column %s, got %T", colName, rawRow)
	}
	row = r.snakeCaseKeys(row)

	return row, nil
}

func (r *RecordUpdater) snakeCaseKeys(data map[string]any) map[string]any {
	newData := make(map[string]any)
	for key, value := range data {
		newData[r.caser.ToSnake(key)] = value
	}
	return newData
}

func preprocessTypeSchema(typeSchema map[string]string) map[string]string {
	for key, typ := range typeSchema {
		if typ == "any" || typ == "binary" {
			typeSchema[key] = schemaupdater.JSONType
		}
	}
	return typeSchema
}
