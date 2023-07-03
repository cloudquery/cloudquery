package client

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/apache/arrow/go/v13/arrow/array"
	"github.com/apache/arrow/go/v13/arrow/memory"
	"github.com/cloudquery/filetypes/v4"
	"github.com/cloudquery/plugin-sdk/v4/message"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
)

const (
	bucket = "cq-playground-test"
	region = "us-east-1"
)

func TestPlugin(t *testing.T) {
	zero := int64(0)
	for _, ft := range []filetypes.FormatType{
		filetypes.FormatTypeCSV,
		filetypes.FormatTypeJSON,
		filetypes.FormatTypeParquet,
	} {
		spec := Spec{
			Bucket:         bucket,
			Region:         region,
			Path:           t.TempDir()[1:],
			NoRotate:       true,
			BatchSizeBytes: &zero,
			BatchSize:      &zero,
			FileSpec: &filetypes.FileSpec{
				Format: ft,
			},
		}

		t.Run("generic/"+string(ft), func(t *testing.T) {
			testPlugin(t, &spec)
		})

		t.Run("write/"+string(ft), func(t *testing.T) {
			testPluginCustom(t, &spec)
		})
	}
}

func testPlugin(t *testing.T, spec *Spec) {
	ctx := context.Background()
	p := plugin.NewPlugin("s3", "development", New)
	b, err := json.Marshal(spec)
	if err != nil {
		t.Fatal(err)
	}
	if err := p.Init(ctx, b, plugin.NewClientOptions{}); err != nil {
		t.Fatal(err)
	}
	plugin.TestWriterSuiteRunner(t,
		p,
		plugin.WriterTestSuiteTests{
			SkipInsert:      true,
			SkipUpsert:      true,
			SkipMigrate:     true,
			SkipDeleteStale: true,
		},
	)
}

func testPluginCustom(t *testing.T, spec *Spec) {
	ctx := context.Background()

	var client plugin.Client

	p := plugin.NewPlugin("s3", "development", func(ctx context.Context, logger zerolog.Logger, spec []byte, opts plugin.NewClientOptions) (plugin.Client, error) {
		var err error
		client, err = New(ctx, logger, spec, opts)
		return client, err
	})
	b, err := json.Marshal(spec)
	if err != nil {
		t.Fatal(err)
	}
	if err := p.Init(ctx, b, plugin.NewClientOptions{}); err != nil {
		t.Fatal(err)
	}

	tableName := fmt.Sprintf("cq_test_custom_insert_%d", time.Now().UnixNano())
	table := &schema.Table{
		Name: tableName,
		Columns: []schema.Column{
			{Name: "name", Type: arrow.BinaryTypes.String},
		},
	}
	if err := p.WriteAll(ctx, []message.WriteMessage{
		&message.WriteMigrateTable{
			Table: table,
		},
	}); err != nil {
		t.Fatal(fmt.Errorf("failed to create table: %w", err))
	}

	bldr := array.NewRecordBuilder(memory.DefaultAllocator, table.ToArrowSchema())
	bldr.Field(0).(*array.StringBuilder).Append("foo")
	record := bldr.NewRecord()

	if err := p.WriteAll(ctx, []message.WriteMessage{
		&message.WriteInsert{
			Record: record,
		},
		&message.WriteInsert{
			Record: record,
		},
	}); err != nil {
		t.Fatal(fmt.Errorf("failed to insert record: %w", err))
	}

	if err := client.Close(ctx); err != nil {
		t.Fatal(fmt.Errorf("failed to close client: %w", err))
	}

	readRecords, err := readAll(ctx, client, table)
	if err != nil {
		t.Fatal(fmt.Errorf("failed to sync: %w", err))
	}

	totalItems := plugin.TotalRows(readRecords)
	assert.Equalf(t, int64(2), totalItems, "expected 2 items, got %d", totalItems)
}

func readAll(ctx context.Context, client plugin.Client, table *schema.Table) ([]arrow.Record, error) {
	var err error
	ch := make(chan arrow.Record)
	go func() {
		defer close(ch)
		err = client.Read(ctx, table, ch)
	}()
	// nolint:prealloc
	var records []arrow.Record
	for record := range ch {
		records = append(records, record)
	}
	return records, err
}
