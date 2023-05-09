package client

import (
	"context"
	"io"
	"path"
	"reflect"
	"regexp"
	"strings"
	"time"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/apache/arrow/go/v13/arrow/array"
	"github.com/apache/arrow/go/v13/arrow/memory"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/cloudquery/filetypes/v3"
	"github.com/cloudquery/plugin-pb-go/specs"
	"github.com/cloudquery/plugin-sdk/v3/plugins/destination/batchingwriter"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/cloudquery/plugin-sdk/v3/types"
	"github.com/google/uuid"
)

const (
	PathVarFormat = "{{FORMAT}}"
	PathVarTable  = "{{TABLE}}"
	PathVarUUID   = "{{UUID}}"
	YearVar       = "{{YEAR}}"
	MonthVar      = "{{MONTH}}"
	DayVar        = "{{DAY}}"
	HourVar       = "{{HOUR}}"
	MinuteVar     = "{{MINUTE}}"
)

var reInvalidJSONKey = regexp.MustCompile(`\W`)

var _ batchingwriter.OpenCloseWriter = (*Client)(nil)

func (c *Client) OpenTable(ctx context.Context, sourceSpec specs.Source, table *schema.Table, syncTime time.Time) error {
	c.logger.Debug().Str("source", sourceSpec.Name).Str("table", table.Name).Msg("OpenTable")

	c.tableWorkersMu.Lock()
	defer c.tableWorkersMu.Unlock()

	objKey := replacePathVariables(c.pluginSpec.Path, table.Name, uuid.NewString(), c.pluginSpec.Format, syncTime)

	pr, pw := io.Pipe()
	doneCh := make(chan error)

	go func() {
		_, err := c.uploader.Upload(ctx, &s3.PutObjectInput{
			Bucket: aws.String(c.pluginSpec.Bucket),
			Key:    aws.String(objKey),
			Body:   pr,
		})
		_ = pr.CloseWithError(err)
		doneCh <- err
		close(doneCh)
	}()

	h, err := c.Client.WriteHeader(pw, table)
	if err != nil {
		_ = pw.CloseWithError(err)
		<-doneCh
		return err
	}

	mk := mapKey(sourceSpec, table)
	c.tableWorkers[mk] = &worker{
		h:    h,
		pw:   pw,
		done: doneCh,
	}
	return nil
}

func (c *Client) CloseTable(_ context.Context, sourceSpec specs.Source, table *schema.Table, _ time.Time) error {
	c.logger.Debug().Str("source", sourceSpec.Name).Str("table", table.Name).Msg("CloseTable")

	c.tableWorkersMu.Lock()
	mk := mapKey(sourceSpec, table)
	wkr := c.tableWorkers[mk]
	c.tableWorkersMu.Unlock()

	err := wkr.h.WriteFooter()
	_ = wkr.pw.CloseWithError(err)

	err = <-wkr.done

	c.tableWorkersMu.Lock()
	delete(c.tableWorkers, mk)
	c.tableWorkersMu.Unlock()

	return err
}

func (c *Client) WriteTableBatch(_ context.Context, sourceSpec specs.Source, table *schema.Table, _ time.Time, data []arrow.Record) error {
	if len(data) == 0 {
		return nil
	}

	if c.pluginSpec.Athena {
		for i, record := range data {
			data[i] = sanitizeRecordJSONKeys(record)
		}
	}

	c.tableWorkersMu.RLock()
	key := mapKey(sourceSpec, table)
	wkr := c.tableWorkers[key]
	c.tableWorkersMu.RUnlock()

	return wkr.h.WriteContent(data)
}

// sanitizeRecordJSONKeys replaces all invalid characters in JSON keys with underscores. This is required
// for compatibility with Athena.
func sanitizeRecordJSONKeys(record arrow.Record) arrow.Record {
	cols := make([]arrow.Array, record.NumCols())
	for i, col := range record.Columns() {
		if arrow.TypeEqual(col.DataType(), types.NewJSONType()) {
			b := types.NewJSONBuilder(array.NewExtensionBuilder(memory.DefaultAllocator, types.NewJSONType()))
			for r := 0; r < int(record.NumRows()); r++ {
				if col.IsNull(r) {
					b.AppendNull()
					continue
				}
				obj := col.GetOneForMarshal(r)
				sanitizeJSONKeysForObject(obj)
				b.Append(obj)
			}
			cols[i] = b.NewArray()
			continue
		}
		cols[i] = col
	}
	return array.NewRecord(record.Schema(), cols, record.NumRows())
}

func sanitizeJSONKeysForObject(obj any) {
	value := reflect.ValueOf(obj)
	switch value.Kind() {
	case reflect.Map:
		iter := value.MapRange()
		for iter.Next() {
			k := iter.Key()
			if k.Kind() == reflect.String {
				nk := reInvalidJSONKey.ReplaceAllString(k.String(), "_")
				v := iter.Value()
				sanitizeJSONKeysForObject(v.Interface())
				value.SetMapIndex(k, reflect.Value{})
				value.SetMapIndex(reflect.ValueOf(nk), v)
			}
		}
	case reflect.Slice:
		for i := 0; i < value.Len(); i++ {
			sanitizeJSONKeysForObject(value.Index(i).Interface())
		}
	}
}

func replacePathVariables(specPath, table, fileIdentifier string, format filetypes.FormatType, t time.Time) string {
	name := strings.ReplaceAll(specPath, PathVarTable, table)
	name = strings.ReplaceAll(name, PathVarFormat, string(format))
	name = strings.ReplaceAll(name, PathVarUUID, fileIdentifier)
	name = strings.ReplaceAll(name, YearVar, t.Format("2006"))
	name = strings.ReplaceAll(name, MonthVar, t.Format("01"))
	name = strings.ReplaceAll(name, DayVar, t.Format("02"))
	name = strings.ReplaceAll(name, HourVar, t.Format("15"))
	name = strings.ReplaceAll(name, MinuteVar, t.Format("04"))
	return path.Clean(name)
}

func mapKey(sourceSpec specs.Source, table *schema.Table) string {
	return sourceSpec.Name + ":" + table.Name
}
