package recordupdater

import (
	"encoding/json"
	"fmt"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/cloudquery/cloudquery/plugins/transformer/jsonflattener/client/schemaupdater"
	"github.com/cloudquery/cloudquery/plugins/transformer/jsonflattener/client/util"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/types"
	"github.com/rs/zerolog"
)

// RecordUpdater takes an `arrow.Record` and knows how to make simple subsequent changes to it.
// It doesn't know which table it belongs to or if the changes make sense.
type RecordUpdater struct {
	logger        zerolog.Logger
	record        arrow.Record
	schemaUpdater *schemaupdater.SchemaUpdater
	tableName     string
}

func New(logger zerolog.Logger, record arrow.Record) *RecordUpdater {
	tableName, _ := record.Schema().Metadata().GetValue(schema.MetadataTableName)

	return &RecordUpdater{
		logger:        logger,
		record:        record,
		schemaUpdater: schemaupdater.New(record.Schema()),
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
		var unprocessedTypeSchema any
		if err := json.Unmarshal([]byte(rawTypeSchema), &unprocessedTypeSchema); err != nil {
			r.logger.Error().Err(err).Msg("failed to unmarshal type schema")
			continue
		}
		switch s := unprocessedTypeSchema.(type) {
		case map[string]any:
			typeSchema := preprocessTypeSchema(s)
			fieldTypeSchemas[field.Name] = typeSchema
		default:
			r.logger.Info().Msgf("skipping unsupported type schema: %T", unprocessedTypeSchema)
		}
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

	builders, err := newColumnBuilders(r.tableName, colName, typeSchema, col.(*types.JSONArray))
	if err != nil {
		return nil, err
	}

	for i := 0; i < jsonArrayCol.Len(); i++ {
		row, err := r.preprocessRow(colName, i, jsonArrayCol.Value(i))
		if err != nil {
			return nil, err
		}
		builders.addRow(row)
	}

	// Order of the new columns must match with the schema.
	// The ordering rule is: (1) by original column existing order, (2) by lexicographical order of subkeys
	newColumns := make([]arrow.Array, 0, len(typeSchema))
	keys := util.SortedKeys(typeSchema)
	for _, key := range keys {
		col, err := builders.build(key)
		if err != nil {
			return nil, err
		}
		newColumns = append(newColumns, col)
	}

	return newColumns, nil
}

func (*RecordUpdater) preprocessRow(colName string, rowIndex int, rawRow any) (map[string]any, error) {
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

	return row, nil
}

func preprocessTypeSchema(unprocessedTypeSchema map[string]any) map[string]string {
	typeSchema := make(map[string]string)
	for key, typ := range unprocessedTypeSchema {
		// Edge case: if the key is utf8, we don't process it, because utf8 is a special
		// string that means that there can be many keys with any name.
		if key == "utf8" {
			continue
		}

		// If the type of a given key is not string, we consider it as a JSON type
		// so that we don't flatten deeper than the first level.
		if _, ok := typ.(string); !ok {
			typeSchema[key] = schemaupdater.JSONType
			continue
		}
		strTyp := typ.(string)
		// Binary and any are treated as JSON types, that is, we don't process them.
		if strTyp == "any" || strTyp == "binary" {
			strTyp = schemaupdater.JSONType
		}
		typeSchema[key] = strTyp
	}
	return typeSchema
}
