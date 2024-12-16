package client

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
	"github.com/cloudquery/cloudquery/plugins/destination/file/v5/client/spec"
	"github.com/cloudquery/filetypes/v4"
	"github.com/cloudquery/filetypes/v4/csv"
	"github.com/cloudquery/plugin-sdk/v4/message"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
)

func testFormats() []filetypes.FileSpec {
	return []filetypes.FileSpec{
		{
			Format: filetypes.FormatTypeCSV,
			FormatSpec: csv.CSVSpec{
				SkipHeader: true,
				Delimiter:  ",",
			},
		},
		{
			Format: filetypes.FormatTypeJSON,
		},
		{
			Format: filetypes.FormatTypeParquet,
		},
	}
}

type testSpec struct {
	spec.Spec
	testName string
	baseDir  string
}

func testSpecsWithoutFormat() []testSpec {
	var (
		ret  []testSpec
		zero int64
	)

	ret = append(ret, testSpec{
		testName: "Path",
		Spec: spec.Spec{
			Path:           filepath.Join("{{TABLE}}.{{FORMAT}}"),
			BatchSize:      &zero,
			BatchSizeBytes: &zero,
		},
	})

	ret = append(ret, testSpec{
		testName: "PathWithTable",
		Spec: spec.Spec{
			Path:           filepath.Join("{{TABLE}}", "data.{{FORMAT}}"),
			BatchSize:      &zero,
			BatchSizeBytes: &zero,
		},
	})

	return ret
}

func testSpecs(t *testing.T) []testSpec {
	var ret []testSpec
	formats := testFormats()
	for _, s := range testSpecsWithoutFormat() {
		s := s
		s.NoRotate = true
		for i := range formats {
			s2 := s
			s2.testName += ":" + string(formats[i].Format)
			s2.FileSpec = formats[i]
			ret = append(ret, s2)

			if formats[i].Format != filetypes.FormatTypeParquet {
				s2.testName += ":gzip"
				fs := s2.FileSpec
				fs.Compression = filetypes.CompressionTypeGZip
				s2.FileSpec = fs
				ret = append(ret, s2)
			}
		}
	}

	for i := range ret {
		bd := t.TempDir()
		ret[i].baseDir = bd
		ret[i].Spec.Path = filepath.Join(bd, ret[i].Spec.Path)
	}

	return ret
}

func TestPlugin(t *testing.T) {
	syncAfterWrite = true // turn on sync after write and wait before read
	for _, ts := range testSpecs(t) {
		ts := ts
		t.Run(ts.testName, func(t *testing.T) {
			t.Parallel()
			if ts.Spec.Format == filetypes.FormatTypeParquet || ts.Spec.Compression != filetypes.CompressionTypeNone {
				testPluginCustom(t, &ts.Spec)
			} else {
				testPlugin(t, &ts.Spec)
			}

			fi, err := os.Stat(ts.baseDir)
			assert.NoError(t, err)
			assert.Truef(t, fi.IsDir(), "basedir %s is not a directory", ts.baseDir)

			fileCount := 0
			assert.NoError(t, filepath.WalkDir(ts.baseDir, func(path string, d os.DirEntry, err error) error {
				assert.NoError(t, err)
				if err != nil {
					return err
				}
				t.Log("walking path", path)
				if !d.IsDir() {
					fileCount++
				}
				if !assert.NotContainsf(t, path, "{", "path %s still contains template", path) {
					return fmt.Errorf("test failed")
				}
				return nil
			}))

			assert.NotZero(t, fileCount, "no files written to %s", ts.baseDir)
		})
	}
}

func testPlugin(t *testing.T, s *spec.Spec) {
	ctx := context.Background()
	p := plugin.NewPlugin("file", "development", New)
	b, err := json.Marshal(s)
	if err != nil {
		t.Fatal(err)
	}
	if err := p.Init(ctx, b, plugin.NewClientOptions{}); err != nil {
		t.Fatal(err)
	}
	plugin.TestWriterSuiteRunner(t,
		p,
		plugin.WriterTestSuiteTests{
			SkipUpsert:       true,
			SkipMigrate:      true,
			SkipDeleteStale:  true,
			SkipDeleteRecord: true,
		},
	)
}

func testPluginCustom(t *testing.T, s *spec.Spec) {
	ctx := context.Background()

	var client plugin.Client

	p := plugin.NewPlugin("file", "development", func(ctx context.Context, logger zerolog.Logger, spec []byte, opts plugin.NewClientOptions) (plugin.Client, error) {
		var err error
		client, err = New(ctx, logger, spec, opts)
		return client, err
	})
	b, err := json.Marshal(s)
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
		t.Fatal(fmt.Errorf("failed to insert records: %w", err))
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
