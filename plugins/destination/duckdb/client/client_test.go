package client

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"path"
	"sync"
	"testing"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
	"github.com/cloudquery/plugin-sdk/v4/message"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/duckdb/duckdb-go/v2"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/require"
)

func TestPlugin(t *testing.T) {
	ctx := context.Background()
	p := plugin.NewPlugin("duckdb", "development", New)
	spec := Spec{
		ConnectionString: "?threads=1",
		Debug:            true,
	}
	specBytes, err := json.Marshal(spec)
	if err != nil {
		t.Fatal(err)
	}

	p.SetLogger(zerolog.New(zerolog.NewTestWriter(t)).Level(zerolog.DebugLevel))

	delayAfterDeleteStale = true
	if err := p.Init(ctx, specBytes, plugin.NewClientOptions{}); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if err := p.Close(ctx); err != nil {
			t.Logf("failed to close plugin: %v", err)
		}
	})

	plugin.TestWriterSuiteRunner(t,
		p,
		plugin.WriterTestSuiteTests{
			SkipDeleteRecord: true,
			SafeMigrations: plugin.SafeMigrations{
				AddColumn:    true,
				RemoveColumn: true,
			},
			SkipSpecificWriteTests: plugin.WriteTests{
				DuplicatePK: true,
			},
		},
		plugin.WithTestDataOptions(schema.TestSourceOptions{
			// not supported in Parquet Writer
			SkipDurations: true,
			SkipIntervals: true,
		}),
	)
}

type testingLog struct {
	testing.TB
	Buf bytes.Buffer
}

func (t *testingLog) Log(args ...any) {
	if _, err := t.Buf.WriteString(fmt.Sprint(args...)); err != nil {
		t.Error(err)
	}
}

func (t *testingLog) Logf(format string, args ...any) {
	if _, err := t.Buf.WriteString(fmt.Sprintf(format, args...)); err != nil {
		t.Error(err)
	}
}

func TestInsertDuplicateSameBatch(t *testing.T) {
	ctx := context.Background()
	p := plugin.NewPlugin("duckdb", "development", New)
	tempDB := path.Join(t.TempDir(), "test_insert_duplicate_same_batch.duckdb") + "?threads=1"

	spec := Spec{
		ConnectionString: tempDB,
		Debug:            true,
	}
	specBytes, err := json.Marshal(spec)
	if err != nil {
		t.Fatal(err)
	}

	testingLog := &testingLog{TB: t, Buf: bytes.Buffer{}}
	testWriter := zerolog.TestWriter{T: testingLog}
	p.SetLogger(zerolog.New(testWriter).Level(zerolog.DebugLevel))

	if err := p.Init(ctx, specBytes, plugin.NewClientOptions{}); err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() {
		if err := p.Close(ctx); err != nil {
			t.Logf("failed to close plugin: %v", err)
		}
	})

	table := &schema.Table{
		Name: "test_insert_duplicate_same_batch",
		Columns: []schema.Column{
			{Name: "id", Type: arrow.PrimitiveTypes.Int64, PrimaryKey: true},
			{Name: "name", Type: arrow.BinaryTypes.String, PrimaryKey: true},
			{Name: "age", Type: arrow.PrimitiveTypes.Int64},
		},
	}
	res := make(chan message.WriteMessage, 10)
	var writeErr error
	wg := sync.WaitGroup{}
	wg.Go(func() {
		writeErr = p.Write(ctx, res)
	})

	res <- &message.WriteMigrateTable{
		Table: table,
	}

	bldr := array.NewRecordBuilder(memory.DefaultAllocator, table.ToArrowSchema())
	bldr.Field(0).(*array.Int64Builder).Append(1)
	bldr.Field(1).(*array.StringBuilder).Append("foo")
	bldr.Field(2).(*array.Int64Builder).Append(20)
	bldr.Field(0).(*array.Int64Builder).Append(1)
	bldr.Field(1).(*array.StringBuilder).Append("foo")
	bldr.Field(2).(*array.Int64Builder).Append(20)

	record := bldr.NewRecordBatch()

	res <- &message.WriteInsert{
		Record: record,
	}
	close(res)

	wg.Wait()
	require.NoError(t, writeErr)

	require.NotContains(t, testingLog.Buf.String(), "error")
	connector, err := duckdb.NewConnector(tempDB, nil)
	require.NoError(t, err)
	defer connector.Close()
	db := sql.OpenDB(connector)
	defer db.Close()

	rows, err := db.QueryContext(ctx, "SELECT count(*) FROM test_insert_duplicate_same_batch")
	require.NoError(t, err)
	defer rows.Close()
	var count int64
	for rows.Next() {
		require.NoError(t, rows.Scan(&count))
	}
	require.NoError(t, rows.Err())
	require.Equal(t, int64(1), count)
}

func TestListPrimaryKey(t *testing.T) {
	ctx := context.Background()
	p := plugin.NewPlugin("duckdb", "development", New)
	tempDB := path.Join(t.TempDir(), "test_list_primary_key.duckdb") + "?threads=1"

	spec := Spec{
		ConnectionString: tempDB,
		Debug:            true,
	}
	specBytes, err := json.Marshal(spec)
	if err != nil {
		t.Fatal(err)
	}

	testingLog := &testingLog{TB: t, Buf: bytes.Buffer{}}
	testWriter := zerolog.TestWriter{T: testingLog}
	p.SetLogger(zerolog.New(testWriter).Level(zerolog.DebugLevel))

	if err := p.Init(ctx, specBytes, plugin.NewClientOptions{}); err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() {
		if err := p.Close(ctx); err != nil {
			t.Logf("failed to close plugin: %v", err)
		}
	})

	table := &schema.Table{
		Name: "test_list_primary_key",
		Columns: []schema.Column{
			{Name: "name", Type: arrow.BinaryTypes.String, PrimaryKey: true},
			{Name: "locations", Type: arrow.ListOf(arrow.BinaryTypes.String), PrimaryKey: true},
		},
	}
	res := make(chan message.WriteMessage, 10)
	var writeErr error
	wg := sync.WaitGroup{}
	wg.Go(func() {
		writeErr = p.Write(ctx, res)
	})

	res <- &message.WriteMigrateTable{
		Table: table,
	}

	bldr := array.NewRecordBuilder(memory.DefaultAllocator, table.ToArrowSchema())
	appendRow := func(name, location string) {
		bldr.Field(0).(*array.StringBuilder).Append(name)
		lb := bldr.Field(1).(*array.ListBuilder)
		lb.Append(true)
		lb.ValueBuilder().(*array.StringBuilder).Append(location)
	}
	appendRow("standard", "eastus")
	appendRow("standard", "westus")
	appendRow("standard", "eastus")

	record := bldr.NewRecordBatch()

	res <- &message.WriteInsert{
		Record: record,
	}
	close(res)

	wg.Wait()
	require.NoError(t, writeErr)

	require.NotContains(t, testingLog.Buf.String(), "Invalid type for index key")
	connector, err := duckdb.NewConnector(tempDB, nil)
	require.NoError(t, err)
	defer connector.Close()
	db := sql.OpenDB(connector)
	defer db.Close()

	var colType string
	require.NoError(t, db.QueryRowContext(ctx,
		"SELECT data_type FROM information_schema.columns WHERE table_name = 'test_list_primary_key' AND column_name = 'locations'").
		Scan(&colType))
	require.Equal(t, "VARCHAR", colType)

	rows, err := db.QueryContext(ctx, "SELECT count(*) FROM test_list_primary_key")
	require.NoError(t, err)
	defer rows.Close()
	var count int64
	for rows.Next() {
		require.NoError(t, rows.Scan(&count))
	}
	require.NoError(t, rows.Err())
	require.Equal(t, int64(2), count)
}

func TestListColumnStringRoundTrip(t *testing.T) {
	listType := arrow.ListOf(arrow.BinaryTypes.String)
	lb := array.NewListBuilder(memory.DefaultAllocator, arrow.BinaryTypes.String)
	defer lb.Release()
	lb.Append(true)
	lb.ValueBuilder().(*array.StringBuilder).Append("eastus")
	lb.ValueBuilder().(*array.StringBuilder).Append("westus")
	listArr := lb.NewArray()
	defer listArr.Release()

	strArr, ok := transformToStringArray(listArr).(*array.String)
	require.True(t, ok)

	back, ok := reverseTransformFromString(listType, strArr).(*array.List)
	require.True(t, ok)
	vals := back.ListValues().(*array.String)
	require.Equal(t, "eastus", vals.Value(0))
	require.Equal(t, "westus", vals.Value(1))
}
