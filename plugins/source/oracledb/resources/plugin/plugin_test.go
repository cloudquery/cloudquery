package plugin

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"sort"
	"strings"
	"testing"
	"time"

	"github.com/apache/arrow/go/v14/arrow"
	"github.com/cloudquery/cloudquery/plugins/source/oracledb/client"
	"github.com/cloudquery/plugin-sdk/v4/message"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/require"
	"golang.org/x/sync/errgroup"

	"github.com/sijms/go-ora/v2/network"
)

func getTestDB(ctx context.Context) (*sql.DB, error) {
	db, err := sql.Open("oracle", getTestConnectionString())
	if err != nil {
		return nil, err
	}
	conn, err := db.Conn(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	return db, err
}

func getTestConnectionString() string {
	testConn := os.Getenv("CQ_SOURCE_ORACLE_DB_TEST_CONNECTION_STRING")
	if testConn == "" {
		return "oracle://cq:test@localhost:/cloudquery"
	}
	return testConn
}

func createTable(ctx context.Context, db *sql.DB, table *schema.Table) error {
	builder := strings.Builder{}
	builder.WriteString("CREATE TABLE ")
	builder.WriteString(client.Identifier(table.Name))
	builder.WriteString(" (\n  ")
	pk := make([]string, 0, len(table.PrimaryKeys()))
	for i, column := range table.Columns {
		if i > 0 {
			builder.WriteString(",\n  ")
		}
		builder.WriteString(client.Identifier(column.Name))
		builder.WriteString(" ")
		builder.WriteString(client.SQLType(column.Type))
		if column.PrimaryKey {
			switch client.SQLType(column.Type) {
			case "clob", "blob":
			// nop, ORA-02329: column of datatype LOB cannot be unique or a primary key
			default:
				pk = append(pk, client.Identifier(column.Name))
			}
			// In OracleDB primary keys are implicitly NOT NULL and UNIQUE
			// and it errors out if we try to do it explicitly
			continue
		}

		if column.NotNull {
			builder.WriteString(" NOT NULL")
		}
		if column.Unique {
			builder.WriteString(" UNIQUE")
		}
	}
	if len(pk) > 0 {
		// Need to move PK to a separate place
		// caused by https://github.com/cloudquery/plugin-sdk/pull/768
		builder.WriteString(",\n  CONSTRAINT ")
		builder.WriteString(client.Identifier(table.Name + "_cq_pk"))
		builder.WriteString(" PRIMARY KEY(")
		builder.WriteString(strings.Join(pk, ", ")) // already quoted
		builder.WriteString(")")
	}
	builder.WriteString("\n)")
	_, err := db.ExecContext(ctx, builder.String())
	return err
}

func insertTable(ctx context.Context, db *sql.DB, table *schema.Table, record arrow.Record) error {
	builder := strings.Builder{}
	builder.WriteString("INSERT INTO " + client.Identifier(table.Name))
	builder.WriteString(" (")
	for i, col := range table.Columns {
		builder.WriteString(client.Identifier(col.Name))
		if i < len(table.Columns)-1 {
			builder.WriteString(", ")
		}
	}
	builder.WriteString(") VALUES (")
	for i := 0; i < len(table.Columns); i++ {
		builder.WriteString(fmt.Sprintf(":%d", i+1))
		if i < len(table.Columns)-1 {
			builder.WriteString(", ")
		}
	}
	builder.WriteString(")")

	transformedRecords, err := client.TransformRecord(record)
	if err != nil {
		return err
	}
	for _, transformedRecord := range transformedRecords {
		if _, err := db.ExecContext(ctx, builder.String(), transformedRecord...); err != nil {
			return err
		}
	}
	return nil
}

func isNotExistsError(err error) bool {
	var dbError *network.OracleError
	if errors.As(err, &dbError) {
		return dbError.ErrCode == 942
	}

	return false
}

func sortResults(table *schema.Table, records []arrow.Record) {
	idIndex := table.Columns.Index("id")
	sort.Slice(records, func(i, j int) bool {
		firstUUID := records[i].Column(idIndex).ValueStr(0)
		secondUUID := records[j].Column(idIndex).ValueStr(0)
		return strings.Compare(firstUUID, secondUUID) < 0
	})
}

func TestPlugin(t *testing.T) {
	p := Plugin()
	ctx := context.Background()
	l := zerolog.New(zerolog.NewTestWriter(t)).Output(
		zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.StampMicro},
	).Level(zerolog.DebugLevel).With().Timestamp().Logger()
	p.SetLogger(l)
	spec := client.Spec{
		ConnectionString: getTestConnectionString(),
	}
	specBytes, err := json.Marshal(spec)
	if err != nil {
		t.Fatal(err)
	}
	db, err := getTestDB(ctx)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	testTable := schema.TestTable("test_oracledb_source", schema.TestSourceOptions{})
	if _, err := db.ExecContext(ctx, "DROP TABLE \"test_oracledb_source\""); err != nil {
		if !isNotExistsError(err) {
			t.Fatal(err)
		}
	}
	if err := createTable(ctx, db, testTable); err != nil {
		t.Fatal(err)
	}
	writtenRecord := schema.NewTestDataGenerator().Generate(testTable, schema.GenTestDataOptions{MaxRows: 2})
	if err := insertTable(ctx, db, testTable, writtenRecord); err != nil {
		t.Fatal(err)
	}

	otherTable := schema.TestTable("other_oracledb_table", schema.TestSourceOptions{})
	if _, err := db.ExecContext(ctx, "DROP TABLE \"other_oracledb_table\""); err != nil {
		if !isNotExistsError(err) {
			t.Fatal(err)
		}
	}
	if err := createTable(ctx, db, otherTable); err != nil {
		t.Fatal(err)
	}
	otherData := schema.NewTestDataGenerator().Generate(otherTable, schema.GenTestDataOptions{MaxRows: 1})
	if err := insertTable(ctx, db, otherTable, otherData); err != nil {
		t.Fatal(err)
	}

	if err := p.Init(ctx, specBytes, plugin.NewClientOptions{}); err != nil {
		t.Fatal(err)
	}
	res := make(chan message.SyncMessage, 1)
	g := errgroup.Group{}

	g.Go(func() error {
		defer close(res)
		opts := plugin.SyncOptions{Tables: []string{testTable.Name}, SkipTables: []string{otherTable.Name}}
		return p.Sync(ctx, opts, res)
	})
	actualRecords := make([]arrow.Record, 0)
	for r := range res {
		m, ok := r.(*message.SyncInsert)
		if ok {
			actualRecords = append(actualRecords, m.Record)
		}
	}
	err = g.Wait()
	if err != nil {
		t.Fatal("got unexpected error:", err)
	}
	if len(actualRecords) != 2 {
		t.Fatalf("expected 2 resource, got %d", len(actualRecords))
	}

	sortResults(testTable, actualRecords)

	for recordIndex := int64(0); recordIndex < writtenRecord.NumRows(); recordIndex++ {
		expectedRecord := writtenRecord.NewSlice(recordIndex, recordIndex+1)
		actualRecord := actualRecords[recordIndex]
		if expectedRecord.NumCols() != actualRecord.NumCols() {
			t.Fatalf("expected record %d to have %d columns, got %d", recordIndex, expectedRecord.NumCols(), actualRecord.NumCols())
		}
		for columnIndex, expectedCol := range expectedRecord.Columns() {
			actualColumn := actualRecord.Column(columnIndex)
			columnName := expectedRecord.ColumnName(columnIndex)
			if expectedCol.Len() != actualColumn.Len() {
				t.Fatalf("expected record %d column %d (%s) to have length %d, got %d", recordIndex, columnIndex, columnName, expectedCol.Len(), actualColumn.Len())
			}
			for arrayIndex := 0; arrayIndex < expectedCol.Len(); arrayIndex++ {
				expectedValue := expectedCol.ValueStr(arrayIndex)
				actualValue := actualColumn.ValueStr(arrayIndex)
				// Oracle treats empty strings as null, so we need to convert them for comparison
				// See https://stackoverflow.com/questions/203493/why-does-oracle-9i-treat-an-empty-string-as-null
				// Please note this means we can't distinguish between null and empty string
				if expectedValue == "" && actualColumn.IsNull(arrayIndex) && arrow.TypeEqual(expectedCol.DataType(), arrow.BinaryTypes.String) {
					actualValue = ""
				}

				if expectedValue != actualValue {
					t.Fatalf("expected record %d column %d (%s) array index %d to have value %s, got %s", recordIndex, columnIndex, columnName, arrayIndex, expectedValue, actualValue)
				}
			}
		}
	}
}

func TestPerformance(t *testing.T) {
	p := Plugin()
	ctx := context.Background()
	l := zerolog.New(zerolog.NewTestWriter(t)).Output(
		zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.StampMicro},
	).Level(zerolog.DebugLevel).With().Timestamp().Logger()
	p.SetLogger(l)
	spec := client.Spec{
		ConnectionString: getTestConnectionString(),
	}
	specBytes, err := json.Marshal(spec)
	if err != nil {
		t.Fatal(err)
	}
	db, err := getTestDB(ctx)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	group, gtx := errgroup.WithContext(ctx)
	group.SetLimit(5)
	const numTables = 20
	for i := 0; i < numTables; i++ {
		table := schema.TestTable(fmt.Sprintf("test_oracledb_source_performance_%d", i), schema.TestSourceOptions{})
		data := schema.NewTestDataGenerator().Generate(table, schema.GenTestDataOptions{MaxRows: 1})

		group.Go(func() error {
			if _, err := db.ExecContext(gtx, fmt.Sprintf("DROP TABLE \"%s\"", table.Name)); err != nil {
				if !isNotExistsError(err) {
					return err
				}
			}
			if err := createTable(gtx, db, table); err != nil {
				return err
			}

			return insertTable(gtx, db, table, data)
		})
	}

	if err := group.Wait(); err != nil {
		t.Fatal(err)
	}

	if err := p.Init(ctx, specBytes, plugin.NewClientOptions{}); err != nil {
		t.Fatal(err)
	}

	res := make(chan message.SyncMessage, 1)
	g := errgroup.Group{}
	g.Go(func() error {
		defer close(res)
		opts := plugin.SyncOptions{Tables: []string{"test_oracledb_source_performance_*"}}
		return p.Sync(ctx, opts, res)
	})
	totalResources := 0
	for r := range res {
		_, ok := r.(*message.SyncInsert)
		if ok {
			totalResources++
		}
	}
	err = g.Wait()
	if err != nil {
		t.Fatal("got unexpected error:", err)
	}

	require.Equal(t, numTables, totalResources)
}
