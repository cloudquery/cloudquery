package plugin

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strings"
	"testing"
	"time"

	"github.com/apache/arrow/go/v14/arrow"
	"github.com/cloudquery/cloudquery/plugins/source/mysql/client"
	"github.com/cloudquery/plugin-sdk/v4/message"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/types"
	"github.com/go-sql-driver/mysql"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/require"
	"golang.org/x/sync/errgroup"
)

func getTestDB(connectionString string) (*sql.DB, error) {
	dsn, err := mysql.ParseDSN(connectionString)
	if err != nil {
		return nil, fmt.Errorf("invalid MySQL connection string: %w", err)
	}
	if dsn.Params == nil {
		dsn.Params = map[string]string{}
	}
	dsn.Params["parseTime"] = "true"
	db, err := sql.Open("mysql", dsn.FormatDSN())
	return db, err
}

func getTestConnectionString() string {
	testConn := os.Getenv("CQ_SOURCE_MYSQL_TEST_CONNECTION_STRING")
	if testConn == "" {
		return "root:test@/cloudquery"
	}
	return testConn
}

func createTable(ctx context.Context, db *sql.DB, table *schema.Table) error {
	builder := strings.Builder{}
	builder.WriteString("CREATE TABLE ")
	builder.WriteString(client.Identifier(table.Name))
	builder.WriteString(" (\n  ")
	for i, column := range table.Columns {
		builder.WriteString(client.Identifier(column.Name))
		builder.WriteString(" ")
		builder.WriteString(client.SQLType(column.Type))
		if column.Unique {
			builder.WriteString(" UNIQUE")
		}
		if column.NotNull {
			builder.WriteString(" NOT NULL")
		}
		if i < len(table.Columns)-1 {
			builder.WriteString(",\n  ")
		}
	}
	pks := table.PrimaryKeys()
	if len(pks) > 0 {
		builder.WriteString(",\n  ")
		builder.WriteString(" PRIMARY KEY (")
		for i, pk := range pks {
			builder.WriteString(client.Identifier(pk))
			if arrow.TypeEqual(table.Columns.Get(pk).Type, arrow.BinaryTypes.String) {
				// Since we use `text` for strings we need to specify the prefix length to use for the primary key
				builder.WriteString("(64)")
			}
			if i < len(pks)-1 {
				builder.WriteString(", ")
			}
		}
		builder.WriteString(")\n")
	}
	builder.WriteString(") CHARACTER SET utf8mb4;")
	_, err := db.ExecContext(ctx, builder.String())
	return err
}

func insertTable(ctx context.Context, db *sql.DB, table *schema.Table, record arrow.Record) error {
	sb := strings.Builder{}
	sb.WriteString("INSERT INTO " + client.Identifier(table.Name))
	sb.WriteString(" (")
	for i, col := range table.Columns {
		sb.WriteString(client.Identifier(col.Name))
		if i < len(table.Columns)-1 {
			sb.WriteString(", ")
		}
	}
	sb.WriteString(") VALUES (")
	sb.WriteString(strings.TrimSuffix(strings.Repeat("?,", len(table.Columns)), ","))
	sb.WriteString(")")

	query := sb.String()
	transformedRecords, err := client.TransformRecord(record)
	if err != nil {
		return err
	}
	for _, transformedRecord := range transformedRecords {
		if _, err := db.ExecContext(ctx, query, transformedRecord...); err != nil {
			return err
		}
	}
	return nil
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
	db, err := getTestDB(spec.ConnectionString)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	testTable := schema.TestTable("test_mysql_source", schema.TestSourceOptions{})
	if _, err := db.ExecContext(ctx, "DROP TABLE IF EXISTS test_mysql_source"); err != nil {
		t.Fatal(err)
	}
	if err := createTable(ctx, db, testTable); err != nil {
		t.Fatal(err)
	}
	syncTime := time.Now()
	writtenRecord := schema.NewTestDataGenerator().Generate(testTable, schema.GenTestDataOptions{MaxRows: 2, SyncTime: syncTime})
	if err := insertTable(ctx, db, testTable, writtenRecord); err != nil {
		t.Fatal(err)
	}

	otherTable := schema.TestTable("other_mysql_table", schema.TestSourceOptions{})
	if _, err := db.ExecContext(ctx, "DROP TABLE IF EXISTS other_mysql_table"); err != nil {
		t.Fatal(err)
	}
	if err := createTable(ctx, db, otherTable); err != nil {
		t.Fatal(err)
	}
	otherData := schema.NewTestDataGenerator().Generate(otherTable, schema.GenTestDataOptions{MaxRows: 1, SyncTime: syncTime})
	if err := insertTable(ctx, db, otherTable, otherData); err != nil {
		t.Fatal(err)
	}

	// Init the plugin so we can call migrate
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
		for columnIndex, col := range expectedRecord.Columns() {
			actualColumn := actualRecord.Column(columnIndex)
			columnName := expectedRecord.ColumnName(columnIndex)
			if col.Len() != actualColumn.Len() {
				t.Fatalf("expected record %d column %d (%s) to have length %d, got %d", recordIndex, columnIndex, columnName, col.Len(), actualColumn.Len())
			}
			for arrayIndex := 0; arrayIndex < col.Len(); arrayIndex++ {
				var expectedValue, actualValue any
				switch c := col.(type) {
				case *types.JSONArray:
					// TODO: Remove this when https://bugs.mysql.com/bug.php?id=98135 is fixed
					// We can't do string comparison for JSON columns as MySQL saves them with extra spaces
					expectedValue = c.Value(arrayIndex)
					actualValue = actualColumn.(*types.JSONArray).Value(arrayIndex)
				default:
					expectedValue = col.ValueStr(arrayIndex)
					actualValue = actualColumn.ValueStr(arrayIndex)
				}

				require.Equal(t, expectedValue, actualValue, "expected record %d column %d (%s) array index %d to have value %s, got %s", recordIndex, columnIndex, columnName, arrayIndex, expectedValue, actualValue)
			}
		}
	}
}
