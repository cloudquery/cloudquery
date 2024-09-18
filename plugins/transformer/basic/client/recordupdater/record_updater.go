package recordupdater

import (
	"crypto/sha256"
	"fmt"
	"strings"
	"time"

	"github.com/apache/arrow/go/v17/arrow"
	"github.com/apache/arrow/go/v17/arrow/array"
	"github.com/apache/arrow/go/v17/arrow/memory"
	"github.com/cloudquery/cloudquery/plugins/transformer/basic/client/schemaupdater"
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

func (r *RecordUpdater) RemoveColumns(columnNames []string) (arrow.Record, error) {
	colIndices, err := r.colIndicesByNames(columnNames)
	if err != nil {
		return nil, err
	}
	if len(colIndices) == int(r.record.NumCols()) {
		return nil, fmt.Errorf("cannot remove all columns")
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
	newSchema, err := r.schemaUpdater.AddStringColumnAtPos(columnName, position, true)
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

func (r *RecordUpdater) ObfuscateColumns(columnNames []string) (arrow.Record, error) {
	colIndex, err := r.colIndicesByNames(columnNames)
	if err != nil {
		return nil, err
	}

	oldRecord := r.record.Columns()
	newColumns := make([]arrow.Array, 0, len(oldRecord))
	for i, column := range oldRecord {
		if _, ok := colIndex[i]; ok {
			if column.DataType().ID() != arrow.STRING {
				return nil, fmt.Errorf("column %v is not a string column", r.record.ColumnName(i))
			}
			newColumns = append(newColumns, r.obfuscateColumn(column))
		} else {
			newColumns = append(newColumns, column)
		}
	}

	r.record = array.NewRecord(r.record.Schema(), newColumns, r.record.NumRows())
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
		bld.AppendString(fmt.Sprintf("%x", sha256.Sum256([]byte(column.ValueStr(i)))))
	}
	return bld.NewStringArray()
}
