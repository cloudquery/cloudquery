package client

import (
	"context"
	"crypto/rsa"
	"crypto/x509"
	"encoding/binary"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
	"github.com/cloudquery/plugin-sdk/v4/message"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"github.com/snowflakedb/gosnowflake/v2"
	"github.com/stretchr/testify/require"
	"golang.org/x/sync/errgroup"
)

// snowflakeTestDSN gates the Snowflake tests behind CQ_SNOWFLAKE_TEST, since they hit a live, paid account.
func snowflakeTestDSN(t *testing.T) string {
	t.Helper()
	if os.Getenv("CQ_SNOWFLAKE_TEST") == "" {
		t.Skip("Snowflake tests are disabled by default; set CQ_SNOWFLAKE_TEST=1 (and the SNOW_* secrets) to enable them")
	}

	account := requireSnowflakeEnv(t, "SNOW_ACCOUNT")
	user := requireSnowflakeEnv(t, "SNOW_USER")
	privateKeyPEM := requireSnowflakeEnv(t, "SNOW_PRIVATE_KEY")
	database := requireSnowflakeEnv(t, "SNOW_DATABASE")
	schemaName := requireSnowflakeEnv(t, "SNOW_SCHEMA")
	warehouse := requireSnowflakeEnv(t, "SNOW_WAREHOUSE")

	cfg := &gosnowflake.Config{
		Account:       account,
		User:          user,
		Database:      database,
		Schema:        schemaName,
		Warehouse:     warehouse,
		Authenticator: gosnowflake.AuthTypeJwt,
		PrivateKey:    parseSnowflakePrivateKey(t, privateKeyPEM),
	}
	// Some accounts encode the region in the account identifier, where also setting Region conflicts.
	if region := os.Getenv("SNOW_REGION"); region != "" {
		cfg.Region = region
	}

	dsn, err := gosnowflake.DSN(cfg)
	if err != nil {
		t.Fatalf("building Snowflake DSN: %v", err)
	}
	return dsn
}

func requireSnowflakeEnv(t *testing.T, name string) string {
	t.Helper()
	v := os.Getenv(name)
	if v == "" {
		t.Fatalf("%s must be set when CQ_SNOWFLAKE_TEST is enabled", name)
	}
	return v
}

func parseSnowflakePrivateKey(t *testing.T, pemStr string) *rsa.PrivateKey {
	t.Helper()
	// GitHub secrets normally preserve real newlines, but tolerate literal "\n".
	pemStr = strings.ReplaceAll(pemStr, `\n`, "\n")
	block, _ := pem.Decode([]byte(pemStr))
	if block == nil {
		t.Fatal("SNOW_PRIVATE_KEY is not a valid PEM block")
	}
	parsed, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		t.Fatalf("parsing SNOW_PRIVATE_KEY: %v", err)
	}
	key, ok := parsed.(*rsa.PrivateKey)
	if !ok {
		t.Fatalf("SNOW_PRIVATE_KEY: expected *rsa.PrivateKey, got %T", parsed)
	}
	return key
}

func TestPlugin(t *testing.T) {
	dsn := snowflakeTestDSN(t)
	ctx := context.Background()
	p := plugin.NewPlugin("snowflake", "development", New)
	spec := &Spec{
		ConnectionString: dsn,
	}
	specBytes, err := json.Marshal(spec)
	if err != nil {
		t.Fatal(err)
	}
	if err := p.Init(ctx, specBytes, plugin.NewClientOptions{}); err != nil {
		t.Fatal(err)
	}
	plugin.TestWriterSuiteRunner(t,
		p,
		plugin.WriterTestSuiteTests{
			SkipDeleteRecord: true,
			SkipSpecificWriteTests: plugin.WriteTests{
				DuplicatePK: true,
			},
			SafeMigrations: plugin.SafeMigrations{
				AddColumn:              true,
				AddColumnNotNull:       false,
				RemoveColumn:           true,
				RemoveColumnNotNull:    false,
				RemoveUniqueConstraint: true,
				MovePKToCQOnly:         true,
			},
		},
		plugin.WithTestDataOptions(schema.TestSourceOptions{
			SkipIntervals:  true,
			SkipMaps:       true,
			SkipLargeTypes: true,
			SkipLists:      true,
		}),
	)
}

// TestConcurrentSyncsSameSchema reproduces the concurrent-write data loss: with `create or replace stage`,
// one sync's setup empties the shared stage and drops files another sync has PUT but not yet COPY'd.
func TestConcurrentSyncsSameSchema(t *testing.T) {
	dsn := snowflakeTestDSN(t)

	const (
		syncConcurrency = 4
		rowsPerSync     = 10
	)

	ctx := context.Background()
	// Snowflake unquoted identifiers can't contain hyphens, so strip them from the UUID.
	tableName := "cq_test_concurrent_syncs_" + strings.ReplaceAll(uuid.NewString(), "-", "")
	table := &schema.Table{
		Name: tableName,
		Columns: []schema.Column{
			schema.CqIDColumn,
			schema.CqSourceNameColumn,
			schema.CqSyncTimeColumn,
		},
	}

	newSyncPlugin := func(t *testing.T) *plugin.Plugin {
		t.Helper()
		p := plugin.NewPlugin("snowflake", "development", New)
		p.SetLogger(zerolog.New(zerolog.NewTestWriter(t)).Level(zerolog.WarnLevel))
		specBytes, err := json.Marshal(&Spec{ConnectionString: dsn})
		require.NoError(t, err)
		require.NoError(t, p.Init(ctx, specBytes, plugin.NewClientOptions{}))
		return p
	}

	p := newSyncPlugin(t)
	require.NoError(t, p.WriteAll(ctx, []message.WriteMessage{&message.WriteMigrateTable{Table: table}}))

	arrowSchema := table.ToArrowSchema()
	group, groupCtx := errgroup.WithContext(ctx)
	for w := range syncConcurrency {
		group.Go(func() error {
			sp := newSyncPlugin(t)

			// Build rows via the generic array.Builder interface: _cq_id is the types.UUID extension, so asserting a concrete builder type would panic.
			bldr := array.NewRecordBuilder(memory.DefaultAllocator, arrowSchema)
			syncTime := time.Now().UTC().Truncate(time.Microsecond)
			for i := range rowsPerSync {
				// Give every row across all writers a distinct _cq_id (PK); otherwise a MERGE destination would dedupe them and skew the count for reasons unrelated to concurrency.
				var id uuid.UUID
				binary.BigEndian.PutUint64(id[8:], uint64(w*rowsPerSync+i))
				for col, field := range arrowSchema.Fields() {
					var valueJSON string
					switch field.Name {
					case schema.CqIDColumn.Name:
						valueJSON = fmt.Sprintf("[%q]", id.String())
					case schema.CqSourceNameColumn.Name:
						valueJSON = `["source"]`
					case schema.CqSyncTimeColumn.Name:
						valueJSON = fmt.Sprintf("[%d]", syncTime.UnixMicro())
					default:
						return fmt.Errorf("unexpected column %q in test table schema", field.Name)
					}
					if err := bldr.Field(col).UnmarshalJSON([]byte(valueJSON)); err != nil {
						return fmt.Errorf("failed to build column %q: %w", field.Name, err)
					}
				}
			}
			record := bldr.NewRecordBatch()
			if err := sp.WriteAll(groupCtx, []message.WriteMessage{&message.WriteInsert{Record: record}}); err != nil {
				return fmt.Errorf("failed to insert records: %w", err)
			}
			return nil
		})
	}
	require.NoError(t, group.Wait())

	ch := make(chan arrow.RecordBatch)
	var readErr error
	go func() {
		defer close(ch)
		readErr = p.Read(ctx, table, ch)
	}()

	numRows := 0
	for record := range ch {
		numRows += int(record.NumRows())
	}
	require.NoError(t, readErr)
	require.Equal(t, syncConcurrency*rowsPerSync, numRows)
}

// TestMigrateMultipleTablesAddColumn covers a batch of tables all needing the same schema change:
// when getTableInfo saw only the first of them, the rest were treated as new and never altered.
func TestMigrateMultipleTablesAddColumn(t *testing.T) {
	dsn := snowflakeTestDSN(t)
	ctx := context.Background()

	specBytes, err := json.Marshal(&Spec{ConnectionString: dsn})
	require.NoError(t, err)
	pc, err := New(ctx, zerolog.New(zerolog.NewTestWriter(t)).Level(zerolog.WarnLevel), specBytes, plugin.NewClientOptions{})
	require.NoError(t, err)
	c := pc.(*Client)
	t.Cleanup(func() { require.NoError(t, c.Close(ctx)) })

	suffix := strings.ReplaceAll(uuid.NewString(), "-", "")
	tableNames := []string{
		"cq_test_migrate_a_" + suffix,
		"cq_test_migrate_b_" + suffix,
		"cq_test_migrate_c_" + suffix,
	}
	t.Cleanup(func() {
		for _, name := range tableNames {
			require.NoError(t, c.dropTable(context.Background(), name))
		}
	})

	const addedColumn = "added_after_create"
	buildTables := func(withAddedColumn bool) message.WriteMigrateTables {
		msgs := make(message.WriteMigrateTables, 0, len(tableNames))
		for _, name := range tableNames {
			table := &schema.Table{
				Name:    name,
				Columns: []schema.Column{schema.CqIDColumn, schema.CqSourceNameColumn},
			}
			if withAddedColumn {
				table.Columns = append(table.Columns, schema.Column{Name: addedColumn, Type: arrow.BinaryTypes.String})
			}
			msgs = append(msgs, &message.WriteMigrateTable{Table: table})
		}
		return msgs
	}

	require.NoError(t, c.MigrateTables(ctx, buildTables(false)))
	require.NoError(t, c.MigrateTables(ctx, buildTables(true)))

	existing, _, err := c.getTableInfo(ctx, tableNames)
	require.NoError(t, err)
	require.Len(t, existing, len(tableNames))
	for _, name := range tableNames {
		table := existing.Get(strings.ToUpper(name))
		require.NotNil(t, table, "table %s missing from information_schema", name)
		require.NotNil(t, table.Column(strings.ToUpper(addedColumn)), "table %s was never migrated to add %s", name, addedColumn)
	}
}
