package plugin

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/cloudquery/cloudquery/plugins/source/mysql/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/cloudquery/plugin-sdk/testdata"
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
		if column.CreationOptions.Unique {
			builder.WriteString(" UNIQUE")
		}
		if column.CreationOptions.NotNull {
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
			if table.Columns.Get(pk).Type == schema.TypeString {
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

func insertTable(ctx context.Context, db *sql.DB, table *schema.Table, data schema.CQTypes) error {
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
	dbData := schema.TransformWithTransformer(&client.Transformer{}, data)
	if _, err := db.ExecContext(ctx, sb.String(), dbData...); err != nil {
		return err
	}
	return nil
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

	testTable := testdata.TestSourceTable("test_mysql_source")
	if _, err := db.ExecContext(ctx, "DROP TABLE IF EXISTS test_mysql_source"); err != nil {
		t.Fatal(err)
	}
	if err := createTable(ctx, db, testTable); err != nil {
		t.Fatal(err)
	}
	data := testdata.GenTestData(testTable)
	if err := insertTable(ctx, db, testTable, data); err != nil {
		t.Fatal(err)
	}

	otherTable := testdata.TestSourceTable("other_mysql_table")
	if _, err := db.ExecContext(ctx, "DROP TABLE IF EXISTS other_mysql_table"); err != nil {
		t.Fatal(err)
	}
	if err := createTable(ctx, db, otherTable); err != nil {
		t.Fatal(err)
	}
	otherData := testdata.GenTestData(otherTable)
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
		return p.Sync(ctx, res)
	})
	var resource *schema.Resource
	totalResources := 0
	for r := range res {
		resource = r
		totalResources++
	}
	err = g.Wait()
	if err != nil {
		t.Fatal("got unexpected error:", err)
	}
	if totalResources != 1 {
		t.Fatalf("expected 1 resource, got %d", totalResources)
	}
	gotData := resource.GetValues()
	actualStrings := make([]string, len(gotData))
	for i, v := range gotData {
		actualStrings[i] = v.String()
	}
	expectedStrings := make([]string, len(data))
	for i, v := range data {
		expectedStrings[i] = v.String()
	}
	require.Equal(t, expectedStrings, actualStrings)
}
