package plugin

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"sort"
	"strings"
	"testing"
	"time"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/cloudquery/plugins/source/mysql/client"
	"github.com/cloudquery/plugin-pb-go/specs"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/cloudquery/plugin-sdk/v3/types"
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

func insertTable(ctx context.Context, db *sql.DB, table *schema.Table, records []arrow.Record) error {
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
	for _, record := range records {
		transformedRecords, err := client.TransformRecord(record)
		if err != nil {
			return err
		}
		for _, transformedRecord := range transformedRecords {
			if _, err := db.ExecContext(ctx, query, transformedRecord...); err != nil {
				return err
			}
		}
	}
	return nil
}

func getActualStringsForResource(resource *schema.Resource) []string {
	actualData := resource.GetValues()
	actualStrings := make([]string, len(actualData))
	for i, v := range actualData {
		// TODO: Remove this when https://bugs.mysql.com/bug.php?id=98135 is fixed
		// We can't do string comparison for JSON columns as MySQL saves them with extra spaces
		if v.DataType() != types.ExtensionTypes.JSON {
			actualStrings[i] = v.String()
		}
	}
	return actualStrings
}

func getActualStrings(resources []*schema.Resource) [][]string {
	actualStrings := make([][]string, len(resources))
	for i, resource := range resources {
		actualStrings[i] = getActualStringsForResource(resource)
	}
	return actualStrings
}

func getExpectedStringsForRecord(t *testing.T, table *schema.Table, record arrow.Record) []string {
	cloned := table.Copy(nil)
	for i := range cloned.Columns {
		// Normalize column types to the ones supported by MySQL
		sqlType := client.SQLType(cloned.Columns[i].Type)
		cloned.Columns[i].Type = client.SchemaType(sqlType, sqlType)
	}

	transformedRecord, _ := client.TransformRecord(record)
	values := transformedRecord[0]
	resource := schema.NewResourceData(cloned, nil, values)
	for i, col := range cloned.Columns {
		if err := resource.Set(col.Name, values[i]); err != nil {
			t.Fatal(err)
		}
	}

	return getActualStringsForResource(resource)
}

func getExpectedStrings(t *testing.T, table *schema.Table, records []arrow.Record) [][]string {
	expectedStrings := make([][]string, len(records))
	for i, record := range records {
		expectedStrings[i] = getExpectedStringsForRecord(t, table, record)
	}
	return expectedStrings
}

func sortResults(table *schema.Table, records [][]string) {
	cqIDIndex := table.Columns.Index(schema.CqIDColumn.Name)
	sort.Slice(records, func(i, j int) bool {
		firstUUID := records[i][cqIDIndex]
		secondUUID := records[j][cqIDIndex]
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
	spec := specs.Source{
		Name:         "test_mysql_source",
		Path:         "cloudquery/mysql",
		Version:      "vDevelopment",
		Destinations: []string{"test"},
		Tables:       []string{"test_mysql_source"},
		Spec: client.Spec{
			ConnectionString: getTestConnectionString(),
		},
	}
	db, err := getTestDB(spec.Spec.(client.Spec).ConnectionString)
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
	data := schema.GenTestData(testTable, schema.GenTestDataOptions{MaxRows: 2, SyncTime: syncTime})
	if err := insertTable(ctx, db, testTable, data); err != nil {
		t.Fatal(err)
	}

	otherTable := schema.TestTable("other_mysql_table", schema.TestSourceOptions{})
	if _, err := db.ExecContext(ctx, "DROP TABLE IF EXISTS other_mysql_table"); err != nil {
		t.Fatal(err)
	}
	if err := createTable(ctx, db, otherTable); err != nil {
		t.Fatal(err)
	}
	otherData := schema.GenTestData(otherTable, schema.GenTestDataOptions{MaxRows: 1, SyncTime: syncTime})
	if err := insertTable(ctx, db, otherTable, otherData); err != nil {
		t.Fatal(err)
	}

	// Init the plugin so we can call migrate
	if err := p.Init(ctx, spec); err != nil {
		t.Fatal(err)
	}
	res := make(chan *schema.Resource, 1)
	g := errgroup.Group{}

	g.Go(func() error {
		defer close(res)
		return p.Sync(ctx, syncTime, res)
	})
	resources := make([]*schema.Resource, 0)
	for r := range res {
		resources = append(resources, r)
	}
	err = g.Wait()
	if err != nil {
		t.Fatal("got unexpected error:", err)
	}
	if len(resources) != 2 {
		t.Fatalf("expected 2 resource, got %d", len(resources))
	}
	expectedStrings := getExpectedStrings(t, testTable, data)
	actualStrings := getActualStrings(resources)
	sortResults(testTable, actualStrings)

	require.Equal(t, expectedStrings, actualStrings)
}
