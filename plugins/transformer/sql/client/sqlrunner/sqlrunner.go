package sqlrunner

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"io"

	"github.com/apache/arrow-adbc/go/adbc"
	"github.com/apache/arrow-adbc/go/adbc/drivermgr"
	"github.com/apache/arrow/go/v17/arrow"
	"github.com/apache/arrow/go/v17/arrow/array"
	"github.com/apache/arrow/go/v17/arrow/ipc"
	"github.com/apache/arrow/go/v17/arrow/memory"
	"github.com/cloudquery/plugin-sdk/v4/types"
	"github.com/google/uuid"
)

const tempTable = "source_table"

type DuckDBSQLRunner struct {
	ctx  context.Context
	conn adbc.Connection
	db   adbc.Database
}

func New(ctx context.Context) (*DuckDBSQLRunner, error) {
	var drv drivermgr.Driver
	db, err := drv.NewDatabase(map[string]string{
		"driver":     "duckdb",
		"entrypoint": "duckdb_adbc_init",
		"path":       ":memory:",
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create new in-memory DuckDB database: %w", err)
	}

	conn, err := db.Open(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to open connection to new in-memory DuckDB database: %w", err)
	}
	return &DuckDBSQLRunner{ctx: ctx, conn: conn, db: db}, nil
}

func serializeRecord(record arrow.Record) (io.Reader, error) {
	buf := new(bytes.Buffer)
	wr := ipc.NewWriter(buf, ipc.WithSchema(record.Schema()))

	if err := wr.Write(record); err != nil {
		return nil, fmt.Errorf("failed to write record: %w", err)
	}

	if err := wr.Close(); err != nil {
		return nil, fmt.Errorf("failed to close writer: %w", err)
	}

	return buf, nil
}

func (r *DuckDBSQLRunner) importRecord(sr io.Reader) error {
	//TEMP
	stmt, err := r.conn.NewStatement()
	if err != nil {
		return fmt.Errorf("failed to create new statement: %w", err)
	}
	defer stmt.Close()

	if err := stmt.SetSqlQuery(fmt.Sprintf(`
		CREATE TABLE %s (
		    _cq_sync_time timestamp without time zone,
			_cq_source_name text,
			_cq_id uuid UNIQUE,
			_cq_parent_id uuid,
			id text PRIMARY KEY,
			fields text
		);
	`, tempTable)); err != nil {
		return fmt.Errorf("failed to set SQL query: %w", err)
	}
	_, err = stmt.ExecuteUpdate(r.ctx)
	if err != nil {
		return fmt.Errorf("failed to execute query: %w", err)
	}

	rdr, err := ipc.NewReader(sr)
	if err != nil {
		return fmt.Errorf("failed to create IPC reader: %w", err)
	}
	defer rdr.Release()

	stmt, err = r.conn.NewStatement()
	if err != nil {
		return fmt.Errorf("failed to create new statement: %w", err)
	}

	if err := stmt.SetOption(adbc.OptionKeyIngestMode, adbc.OptionValueIngestModeAppend); err != nil {
		return fmt.Errorf("failed to set ingest mode: %w", err)
	}
	// duckdb hasn't implemented temp table ingest yet unfortunately, would be good to update this!
	// stmt.SetOption(adbc.OptionValueIngestTemporary, adbc.OptionValueEnabled)
	// optional!
	// stmt.SetOption(adbc.OptionValueIngestTargetCatalog, "catalog")
	if err := stmt.SetOption(adbc.OptionKeyIngestTargetTable, tempTable); err != nil {
		return fmt.Errorf("failed to set ingest target table: %w", err)
	}

	if err := stmt.BindStream(r.ctx, rdr); err != nil {
		return fmt.Errorf("failed to bind stream: %w", err)
	}

	if _, err := stmt.ExecuteUpdate(r.ctx); err != nil {
		return fmt.Errorf("failed to execute update: %w", err)
	}

	return stmt.Close()
}

func parseSQL(sql string) (string, error) {
	tmpl, err := template.New("sql").Parse(sql)
	if err != nil {
		return "", err
	}

	var sqlBuffer bytes.Buffer
	err = tmpl.Execute(&sqlBuffer, map[string]string{"Table": tempTable})
	if err != nil {
		return "", err
	}
	return sqlBuffer.String(), nil
}

func (r *DuckDBSQLRunner) runSQL(sql string, ignoreOutput bool) ([]arrow.Record, error) {
	stmt, err := r.conn.NewStatement()
	if err != nil {
		return nil, fmt.Errorf("failed to create new statement: %w", err)
	}
	defer stmt.Close()

	sql, err = parseSQL(sql)
	if err != nil {
		return nil, fmt.Errorf("failed to parse SQL: %w", err)
	}

	if err := stmt.SetSqlQuery(sql); err != nil {
		return nil, fmt.Errorf("failed to set SQL query: %w", err)
	}
	out, n, err := stmt.ExecuteQuery(r.ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	defer out.Release()
	if ignoreOutput {
		return nil, nil
	}

	result := make([]arrow.Record, 0, n)
	for out.Next() {
		rec := out.Record()
		rec.Retain() // .Next() will release the record, so we need to retain it
		result = append(result, rec)
	}
	if out.Err() != nil {
		return nil, out.Err()
	}
	return result, nil
}

func (r *DuckDBSQLRunner) RunSQLOnRecord(record arrow.Record, sqls ...string) ([]arrow.Record, error) {
	if len(sqls) == 0 {
		return nil, errors.New("no SQL statement provided")
	}
	serializedRecord, err := serializeRecord(record)
	if err != nil {
		return nil, fmt.Errorf("failed to serialize record: %w", err)
	}
	if err := r.importRecord(serializedRecord); err != nil {
		return nil, fmt.Errorf("failed to import record: %w", err)
	}
	var result []arrow.Record
	for i, sql := range sqls {
		result, err = r.runSQL(sql, i != len(sqls)-1)
		if err != nil {
			return nil, fmt.Errorf("failed to run SQL: %w", err)
		}
	}

	fmt.Println("about to repair metadata and columns")
	repairedResults, err := repairMetadataAndColumns(record, result)
	if err != nil {
		return nil, fmt.Errorf("failed to repair metadata and columns: %w", err)
	}
	fmt.Println("finished repairing metadata and columns")

	if _, err := r.runSQL("DROP TABLE "+tempTable, true); err != nil {
		return nil, fmt.Errorf("failed to drop temp table after running query: %w", err)
	}
	return repairedResults, nil
}

// Unfortunately, at the moment, DuckDB's ADBC implementation swallows metadata, so we need to repair it.
//
// Repairing is slightly magic, because there's no way to know how the query has changed the schema.
//
// The strategy is to keep a map of the old record's field metadata (keyed by name) and to put it back in
// if the new record has a field with the same name. Thus, new fields will not have metadata.
//
// It also converts extension types to their storage layer types, so we must convert them too.
func repairMetadataAndColumns(oldRecord arrow.Record, newRecords []arrow.Record) ([]arrow.Record, error) {
	if len(newRecords) == 0 {
		return newRecords, nil
	}

	oldRecordMd := oldRecord.Schema().Metadata()
	oldRecordFieldsToTypes := make(map[string]arrow.DataType)
	for _, oldField := range oldRecord.Schema().Fields() {
		oldRecordFieldsToTypes[oldField.Name] = oldField.Type
	}

	oldFieldMetadata := make(map[string]arrow.Metadata)
	for _, oldField := range oldRecord.Schema().Fields() {
		if !oldField.HasMetadata() {
			continue
		}
		oldFieldMetadata[oldField.Name] = oldField.Metadata
	}

	repairedFields := make([]arrow.Field, len(newRecords[0].Schema().Fields()))
	for i, newField := range newRecords[0].Schema().Fields() {
		md, ok := oldFieldMetadata[newField.Name]
		if !ok {
			repairedFields[i] = newField
			continue
		}
		repairedFields[i] = arrow.Field{
			Name:     newField.Name,
			Type:     repairExtensionType(oldRecord.Schema().Fields()[i].Type, newField.Type),
			Nullable: newField.Nullable,
			Metadata: md,
		}
	}

	repairedSchema := arrow.NewSchema(repairedFields, &oldRecordMd)
	repairedRecords := make([]arrow.Record, 0, len(newRecords))
	for _, newRecord := range newRecords {
		repairedColumns := make([]arrow.Array, 0, newRecord.NumCols())
		for i, newColumn := range newRecord.Columns() {
			newField := newRecord.Schema().Field(i)
			oldFieldType := oldRecordFieldsToTypes[newField.Name]
			repairedColumn, err := repairColumn(newColumn, oldFieldType, newField.Type)
			if err != nil {
				return nil, fmt.Errorf("failed to repair column: %w", err)
			}
			repairedColumns = append(repairedColumns, repairedColumn)
		}
		repairedRecord := array.NewRecord(repairedSchema, repairedColumns, newRecord.NumRows())
		repairedRecords = append(repairedRecords, repairedRecord)
		newRecord.Release()
	}

	oldRecord.Release()

	return repairedRecords, nil
}

func repairColumn(oldColumn arrow.Array, oldType, newType arrow.DataType) (arrow.Array, error) {
	if oldType == nil {
		return oldColumn, nil
	}
	switch {
	case getTypeNameOrExtensionName(oldType) == "json" && getTypeNameOrExtensionName(newType) == "binary":
		fmt.Println("json")
		return repairBinaryColumnToJSON(oldColumn), nil
	case getTypeNameOrExtensionName(oldType) == "uuid" && getTypeNameOrExtensionName(newType) == "binary":
		fmt.Println("uuid")
		repairedColumn, err := repairUUIDColumnToJSON(oldColumn)
		return repairedColumn, err
	default:
		return oldColumn, nil
	}
}

func repairExtensionType(oldType, newType arrow.DataType) arrow.DataType {
	switch {
	case getTypeNameOrExtensionName(oldType) == "json" && getTypeNameOrExtensionName(newType) == "binary":
		return types.NewJSONType()
	case getTypeNameOrExtensionName(oldType) == "uuid" && getTypeNameOrExtensionName(newType) == "binary":
		return types.NewUUIDType()
	default:
		return oldType
	}
}

func getTypeNameOrExtensionName(dataType arrow.DataType) string {
	if extType, ok := dataType.(arrow.ExtensionType); ok {
		return extType.ExtensionName()
	}
	return dataType.Name()
}

func repairBinaryColumnToJSON(arr arrow.Array) arrow.Array {
	defer arr.Release()
	jsonBuilder := types.NewJSONBuilder(array.NewExtensionBuilder(memory.DefaultAllocator, types.NewJSONType()))
	binArr := arr.(*array.Binary)
	for i := 0; i < binArr.Len(); i++ {
		data := binArr.GetOneForMarshal(i)
		if data == nil {
			jsonBuilder.AppendNull()
			continue
		}
		if len(data.([]byte)) == 0 {
			jsonBuilder.AppendNull()
			continue
		}
		jsonBuilder.Append(json.RawMessage(data.([]byte)))
	}
	return jsonBuilder.NewArray()
}

func repairUUIDColumnToJSON(arr arrow.Array) (arrow.Array, error) {
	defer arr.Release()
	uuidBuilder := types.NewUUIDBuilder(array.NewExtensionBuilder(memory.DefaultAllocator, types.NewUUIDType()))
	binArr := arr.(*array.Binary)
	for i := 0; i < binArr.Len(); i++ {
		uid, err := uuidFromBytes(binArr.Value(i))
		if err != nil {
			return nil, fmt.Errorf("failed to convert binary to UUID: %w", err)
		}
		if uid == uuid.Nil {
			uuidBuilder.AppendNull()
		} else {
			uuidBuilder.Append(uid)
		}
	}
	return uuidBuilder.NewArray(), nil
}

func uuidFromBytes(b []byte) (uuid.UUID, error) {
	if len(b) == 0 {
		return uuid.Nil, nil
	}
	return uuid.FromBytes(b)
}

func (r *DuckDBSQLRunner) Close() {
	r.conn.Close()
	r.db.Close()
}
