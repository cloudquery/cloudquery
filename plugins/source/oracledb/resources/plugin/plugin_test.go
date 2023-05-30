package plugin

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/cloudquery/plugins/source/oracledb/client"
	"github.com/cloudquery/plugin-pb-go/specs"
	"github.com/cloudquery/plugin-sdk/v3/scalar"
	"github.com/cloudquery/plugin-sdk/v3/schema"
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
		} else {
			// In OracleDB primary keys are implicitly NOT NULL and UNIQUE
			// and it errors out if we try to do it explicitly
			if column.NotNull {
				builder.WriteString(" NOT NULL")
			}
			if column.Unique {
				builder.WriteString(" UNIQUE")
			}
		}

		if i < len(table.Columns)-1 {
			builder.WriteString(",\n  ")
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

func insertTable(ctx context.Context, db *sql.DB, table *schema.Table, data []arrow.Record) error {
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

	for _, data := range data {
		transformedRecords, err := client.TransformRecord(data)
		if err != nil {
			return err
		}
		for _, transformedRecord := range transformedRecords {
			if _, err := db.ExecContext(ctx, builder.String(), transformedRecord...); err != nil {
				return err
			}
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

func TestPlugin(t *testing.T) {
	p := Plugin()
	ctx := context.Background()
	l := zerolog.New(zerolog.NewTestWriter(t)).Output(
		zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.StampMicro},
	).Level(zerolog.DebugLevel).With().Timestamp().Logger()
	p.SetLogger(l)
	spec := specs.Source{
		Name:         "test_oracledb_source",
		Path:         "cloudquery/oracledb",
		Version:      "vDevelopment",
		Destinations: []string{"test"},
		Tables:       []string{"test_oracledb_source"},
		Spec: client.Spec{
			ConnectionString: getTestConnectionString(),
		},
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
	data := schema.GenTestData(testTable, schema.GenTestDataOptions{MaxRows: 1})
	if err := insertTable(ctx, db, testTable, data); err != nil {
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
	otherData := schema.GenTestData(otherTable, schema.GenTestDataOptions{MaxRows: 1})
	if err := insertTable(ctx, db, otherTable, otherData); err != nil {
		t.Fatal(err)
	}

	if err := p.Init(ctx, spec); err != nil {
		t.Fatal(err)
	}
	res := make(chan *schema.Resource, 1)
	g := errgroup.Group{}
	syncTime := time.Now()
	g.Go(func() error {
		defer close(res)
		return p.Sync(ctx, syncTime, res)
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
		// Oracle treats empty strings as null, so we need to convert them for comparison
		// See https://stackoverflow.com/questions/203493/why-does-oracle-9i-treat-an-empty-string-as-null
		// Please note this means we can't distinguish between null and empty string
		if !v.IsValid() && v.DataType() == arrow.BinaryTypes.String {
			if err := v.Set(""); err != nil {
				t.Fatal(err)
			}
		}
		actualStrings[i] = v.String()
	}
	// Convert test data to scalar so we can do string comparison
	// We need this as OracleDB does not support all arrow types, so the type data is lost for some columns
	expectedColumns := data[0].Columns()
	expectedStrings := make([]string, len(expectedColumns))
	for i, col := range expectedColumns {
		if col.DataType() == arrow.PrimitiveTypes.Uint64 {
			// When reading from the OracleDB we can't know if the value is int64 or uint64 (without parsing the data)
			// So we skip the test for this column as we can't support it
			actualStrings[i] = ""
			expectedStrings[i] = ""
			continue
		}
		normalizedType := client.SchemaType(client.SQLType(col.DataType()))
		s := scalar.NewScalar(normalizedType)
		val, err := client.GetValue(col, 0)
		if err != nil {
			t.Fatal(err)
		}
		if err := s.Set(val); err != nil {
			t.Fatal(err)
		}
		expectedStrings[i] = s.String()
	}

	require.Equal(t, expectedStrings, actualStrings)
}

func TestPerformance(t *testing.T) {
	p := Plugin()
	ctx := context.Background()
	l := zerolog.New(zerolog.NewTestWriter(t)).Output(
		zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.StampMicro},
	).Level(zerolog.DebugLevel).With().Timestamp().Logger()
	p.SetLogger(l)
	spec := specs.Source{
		Name:         "test_oracledb_source",
		Path:         "cloudquery/oracledb",
		Version:      "vDevelopment",
		Destinations: []string{"test"},
		Tables:       []string{"test_oracledb_source_performance_*"},
		Spec: client.Spec{
			ConnectionString: getTestConnectionString(),
		},
	}
	db, err := getTestDB(ctx)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	group, gtx := errgroup.WithContext(ctx)
	const numTables = 20
	for i := 0; i < numTables; i++ {
		table := schema.TestTable(fmt.Sprintf("test_oracledb_source_performance_%d", i), schema.TestSourceOptions{})
		data := schema.GenTestData(table, schema.GenTestDataOptions{MaxRows: 1})

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

	if err := p.Init(ctx, spec); err != nil {
		t.Fatal(err)
	}

	res := make(chan *schema.Resource, 1)
	g := errgroup.Group{}
	syncTime := time.Now()
	g.Go(func() error {
		defer close(res)
		return p.Sync(ctx, syncTime, res)
	})
	totalResources := 0
	for range res {
		totalResources++
	}
	err = g.Wait()
	if err != nil {
		t.Fatal("got unexpected error:", err)
	}

	require.Equal(t, numTables, totalResources)
}
