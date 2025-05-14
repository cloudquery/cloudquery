package recordupdater

import (
	"crypto/sha256"
	"encoding/json"
	"errors"
	"fmt"
	"slices"
	"strings"
	"time"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
	"github.com/cloudquery/cloudquery/plugins/transformer/basic/client/schemaupdater"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// RecordUpdater takes an `arrow.Record` and knows how to make simple subsequent changes to it.
// It doesn't know which table it belongs to or if the changes make sense.
type RecordUpdater struct {
	record        arrow.Record
	schemaUpdater *schemaupdater.SchemaUpdater
}

func New(record arrow.Record) *RecordUpdater {
	return &RecordUpdater{
		record:        record,
		schemaUpdater: schemaupdater.New(record.Schema()),
	}
}

const redactedByCQMessage = "Redacted by CloudQuery |"
const redactedByCQJSONName = "redacted_by_cloudquery"

func (r *RecordUpdater) RemoveColumns(columnNames []string) (arrow.Record, error) {
	plainCols, jsonCols := r.splitJSONColumns(columnNames)

	if len(plainCols) > 0 {
		colIndices, err := r.colIndicesByNames(plainCols)
		if err != nil {
			return nil, err
		}
		if len(colIndices) == int(r.record.NumCols()) {
			return nil, errors.New("cannot remove all columns")
		}

		oldRecord := r.record.Columns()
		newColumns := make([]arrow.Array, 0, len(oldRecord)-len(colIndices))
		for i, column := range oldRecord {
			if _, ok := colIndices[i]; ok {
				continue
			}
			newColumns = append(newColumns, column)
		}

		r.record = array.NewRecord(r.schemaUpdater.RemoveColumnIndices(colIndices), newColumns, r.record.NumRows())
	}

	if len(jsonCols) > 0 {
		for i, jcs := range r.jsonColIndicesByNames(jsonCols) {
			bld := types.NewJSONBuilder(memory.NewGoAllocator())
			for j := 0; j < r.record.Column(i).Len(); j++ {
				valStr := r.record.Column(i).ValueStr(j)
				if gjson.Valid(valStr) {
					for _, jc := range jcs {
						if val, err := sjson.Delete(valStr, jc.columnPath); err == nil {
							valStr = val
						}
					}
				}
				bld.AppendBytes([]byte(valStr))
			}

			rec, err := r.record.SetColumn(i, bld.NewJSONArray())
			if err != nil {
				return nil, err
			}
			r.record = rec
		}
	}
	return r.record, nil
}

func (r *RecordUpdater) AddLiteralStringColumn(columnName, columnValue string, position int) (arrow.Record, error) {
	if position == -1 {
		position = int(r.record.NumCols())
	}
	if position < 0 || position > int(r.record.NumCols()) {
		return nil, fmt.Errorf("invalid position %v", position)
	}

	newColumns := make([]arrow.Array, 0, int(r.record.NumCols())+1)
	for i := 0; i < int(r.record.NumCols()); i++ {
		if i == position {
			newColumns = append(newColumns, r.buildStringColumn(columnValue, int(r.record.NumRows())))
		}
		newColumns = append(newColumns, r.record.Column(i))
	}
	if position == int(r.record.NumCols()) {
		newColumns = append(newColumns, r.buildStringColumn(columnValue, int(r.record.NumRows())))
	}
	newSchema, err := r.schemaUpdater.AddStringColumnAtPos(columnName, position, false)
	if err != nil {
		return nil, err
	}
	r.record = array.NewRecord(newSchema, newColumns, r.record.NumRows())
	return r.record, nil
}

func (r *RecordUpdater) AddTimestampColumn(columnName string, position int) (arrow.Record, error) {
	if position == -1 {
		position = int(r.record.NumCols())
	}
	if position < 0 || position > int(r.record.NumCols()) {
		return nil, fmt.Errorf("invalid position %v", position)
	}
	timeVal := time.Now()

	newColumns := make([]arrow.Array, 0, int(r.record.NumCols())+1)
	for i := 0; i < int(r.record.NumCols()); i++ {
		if i == position {
			newColumns = append(newColumns, r.buildCurrentTimestampColumn(timeVal, int(r.record.NumRows())))
		}
		newColumns = append(newColumns, r.record.Column(i))
	}
	if position == int(r.record.NumCols()) {
		newColumns = append(newColumns, r.buildCurrentTimestampColumn(timeVal, int(r.record.NumRows())))
	}
	newSchema, err := r.schemaUpdater.AddTimestampColumnAtPos(columnName, position, true)
	if err != nil {
		return nil, err
	}
	r.record = array.NewRecord(newSchema, newColumns, r.record.NumRows())
	return r.record, nil
}

func (r *RecordUpdater) ObfuscateSensitiveColumns() (arrow.Record, error) {
	if r.record.Schema() == nil {
		return nil, errors.New("record schema is nil")
	}
	s, ok := r.record.Schema().Metadata().GetValue(schema.MetadataTableSensitiveColumns)
	if !ok {
		return r.record, nil
	}
	var sensitiveColumnsArr []string
	err := json.Unmarshal([]byte(s), &sensitiveColumnsArr)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal sensitive columns: %w", err)
	}
	if len(sensitiveColumnsArr) == 0 {
		return r.record, nil
	}
	return r.ObfuscateColumns(sensitiveColumnsArr)
}

func (r *RecordUpdater) ObfuscateColumns(columnNames []string) (arrow.Record, error) {
	plainCols, jsonCols := r.splitJSONColumns(columnNames)

	plainColIndex, err := r.colIndicesByNames(plainCols)
	if err != nil {
		return nil, err
	}
	jsonColIndex := r.jsonColIndicesByNames(jsonCols)

	oldRecord := r.record.Columns()
	newColumns := make([]arrow.Array, 0, len(oldRecord))
	for i, column := range oldRecord {
		if _, ok := plainColIndex[i]; ok {
			if column.DataType().ID() == arrow.STRING {
				newColumns = append(newColumns, r.obfuscateColumn(column))
				continue
			}
			if _, ok := column.DataType().(*types.JSONType); ok {
				newColumns = append(newColumns, r.obfuscateEntireJSONColumn(column))
				continue
			}
			return nil, fmt.Errorf("column %v is not a string or JSON column", r.record.ColumnName(i))
		}

		jcs, ok := jsonColIndex[i]
		if !ok {
			newColumns = append(newColumns, column)
			continue
		}

		if _, ok := column.DataType().(*types.JSONType); !ok {
			return nil, fmt.Errorf("column %v is not a JSON column", r.record.ColumnName(i))
		}

		newColumns = append(newColumns, r.obfuscateJSONColumns(column, jcs))
	}

	r.record = array.NewRecord(r.record.Schema(), newColumns, r.record.NumRows())

	return r.record, nil
}

func (r *RecordUpdater) AddPrimaryKeys(columnNames []string) (arrow.Record, error) {
	newSchema, err := r.schemaUpdater.AddPrimaryKeys(columnNames)
	if err != nil {
		return nil, err
	}
	r.record = array.NewRecord(newSchema, r.record.Columns(), r.record.NumRows())
	return r.record, nil
}

func (r *RecordUpdater) ChangeTableName(newTableNamePattern string) (arrow.Record, error) {
	newSchema, err := r.schemaUpdater.ChangeTableName(newTableNamePattern)
	if err != nil {
		return nil, err
	}
	r.record = array.NewRecord(newSchema, r.record.Columns(), r.record.NumRows())
	return r.record, nil
}

func (r *RecordUpdater) RenameColumn(oldName, newName string) (arrow.Record, error) {
	newSchema, err := r.schemaUpdater.RenameColumn(oldName, newName)
	if err != nil {
		return nil, err
	}
	r.record = array.NewRecord(newSchema, r.record.Columns(), r.record.NumRows())
	return r.record, nil
}

func (r *RecordUpdater) colIndicesByNames(columnNames []string) (map[int]struct{}, error) {
	colNameMap := make(map[string]struct{})
	for _, columnName := range columnNames {
		colNameMap[columnName] = struct{}{}
	}

	colIndexes := make(map[int]struct{})
	for i := 0; i < int(r.record.NumCols()); i++ {
		colName := r.record.ColumnName(i)
		if _, ok := colNameMap[colName]; ok {
			colIndexes[i] = struct{}{}
			delete(colNameMap, colName)
		}
	}
	if len(colNameMap) > 0 {
		missingColumns := make([]string, 0, len(colNameMap))
		for colName := range colNameMap {
			missingColumns = append(missingColumns, colName)
		}
		return nil, fmt.Errorf("columns %v not found", strings.Join(missingColumns, ", "))
	}

	return colIndexes, nil
}

type jsonColumn struct {
	columnName string
	columnPath string
}

func (r *RecordUpdater) jsonColIndicesByNames(columns map[string]jsonColumn) map[int][]jsonColumn {
	colNameMap := make(map[int][]jsonColumn)
	for i := 0; i < int(r.record.NumCols()); i++ {
		for _, jc := range columns {
			if jc.columnName == r.record.ColumnName(i) {
				if _, ok := colNameMap[i]; !ok {
					colNameMap[i] = []jsonColumn{jc}
				} else {
					colNameMap[i] = append(colNameMap[i], jc)
				}
			}
		}
	}
	return colNameMap
}

func (*RecordUpdater) buildStringColumn(literalValue string, numRows int) arrow.Array {
	bld := array.NewStringBuilder(memory.DefaultAllocator)
	for i := 0; i < numRows; i++ {
		bld.AppendString(literalValue)
	}
	return bld.NewStringArray()
}

func (*RecordUpdater) buildCurrentTimestampColumn(t time.Time, numRows int) arrow.Array {
	ts, _ := arrow.TimestampFromTime(t, arrow.Microsecond)
	syncTimeBldr := array.NewTimestampBuilder(memory.DefaultAllocator, &arrow.TimestampType{Unit: arrow.Microsecond, TimeZone: "UTC"})
	for i := 0; i < numRows; i++ {
		syncTimeBldr.Append(ts)
	}
	return syncTimeBldr.NewArray()
}

func (*RecordUpdater) obfuscateColumn(column arrow.Array) arrow.Array {
	bld := array.NewStringBuilder(memory.DefaultAllocator)
	for i := 0; i < column.Len(); i++ {
		if !column.IsValid(i) {
			bld.AppendNull()
			continue
		}
		bld.AppendString(fmt.Sprintf("%s %x", redactedByCQMessage, sha256.Sum256([]byte(column.ValueStr(i)))))
	}
	return bld.NewStringArray()
}

func (*RecordUpdater) obfuscateJSONColumns(column arrow.Array, jcs []jsonColumn) arrow.Array {
	bld := types.NewJSONBuilder(memory.NewGoAllocator())
	for i := 0; i < column.Len(); i++ {
		if !column.IsValid(i) {
			bld.AppendNull()
			continue
		}

		str := column.ValueStr(i)
		for _, jc := range jcs {
			val := gjson.Get(column.ValueStr(i), jc.columnPath)
			if val.Exists() && val.Type == gjson.String {
				if modified, err := sjson.Set(str, jc.columnPath, fmt.Sprintf("%s %x", redactedByCQMessage, sha256.Sum256([]byte(val.Str)))); err == nil {
					str = modified
					continue
				}
			}
		}
		bld.AppendBytes([]byte(str))
	}
	return bld.NewJSONArray()
}

func (*RecordUpdater) obfuscateEntireJSONColumn(column arrow.Array) arrow.Array {
	bld := types.NewJSONBuilder(memory.NewGoAllocator())
	for i := 0; i < column.Len(); i++ {
		if !column.IsValid(i) {
			bld.AppendNull()
			continue
		}

		str := column.ValueStr(i)
		newStr := "{}"

		if modified, err := sjson.Set(newStr, redactedByCQJSONName, fmt.Sprintf("%x", sha256.Sum256([]byte(str)))); err == nil {
			str = modified
			bld.AppendBytes([]byte(str))
		}
	}
	return bld.NewJSONArray()
}

func (r *RecordUpdater) splitJSONColumns(columnNames []string) (plainCols []string, jsonCols map[string]jsonColumn) {
	plainColMap := make(map[string]struct{})
	jsonColMap := make(map[string]jsonColumn)
	for _, columnName := range columnNames {
		if idx := strings.Index(columnName, "."); idx > -1 {
			jsonColMap[columnName] = jsonColumn{columnName: columnName[:idx], columnPath: columnName[idx+1:]}
		} else {
			plainColMap[columnName] = struct{}{}
		}
	}

	for i := 0; i < int(r.record.NumCols()); i++ {
		if _, ok := plainColMap[r.record.ColumnName(i)]; ok {
			plainCols = append(plainCols, r.record.ColumnName(i))
			continue
		}
	}

	for k, jc := range jsonColMap {
		if slices.Contains(plainCols, jc.columnName) {
			delete(jsonColMap, k)
		}
	}
	jsonCols = make(map[string]jsonColumn)
	for i := 0; i < int(r.record.NumCols()); i++ {
		for _, jc := range jsonColMap {
			if jc.columnName == r.record.ColumnName(i) {
				jsonCols[jc.columnName+"."+jc.columnPath] = jc
			}
		}
	}

	return plainCols, jsonCols
}
