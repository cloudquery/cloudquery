package client

import (
	"context"
	"fmt"
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
	"github.com/cloudquery/filetypes/v4"
	ftypes "github.com/cloudquery/filetypes/v4/types"
	"github.com/cloudquery/plugin-sdk/v4/message"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/types"
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

type stream struct {
	h    ftypes.Handle
	wc   *writeCloser
	done chan error
}

type writeCloser struct {
	*io.PipeWriter
	closed bool
}

func (w *writeCloser) Close() error {
	w.closed = true
	return w.PipeWriter.Close()
}

func (c *Client) OpenTable(_ context.Context, sourceName string, table *schema.Table, syncTime time.Time) (any, error) {
	objKey := replacePathVariables(c.spec.Path, table.Name, uuid.NewString(), c.spec.Format, syncTime)

	pr, pw := io.Pipe()
	doneCh := make(chan error)

	go func() {
		_, err := c.uploader.Upload(c.uploadCtx, &s3.PutObjectInput{
			Bucket: aws.String(c.spec.Bucket),
			Key:    aws.String(objKey),
			Body:   pr,
		})
		_ = pr.CloseWithError(err)

		doneCh <- err
		close(doneCh)
	}()

	wc := &writeCloser{PipeWriter: pw}
	h, err := c.Client.WriteHeader(wc, table)
	if err != nil {
		_ = pw.CloseWithError(err)
		<-doneCh
		return nil, err
	}

	return &stream{
		h:    h,
		wc:   wc,
		done: doneCh,
	}, nil
}

func (*Client) CloseTable(_ context.Context, handle any) error {
	s := handle.(*stream)
	if err := s.h.WriteFooter(); err != nil {
		if !s.wc.closed {
			_ = s.wc.CloseWithError(err)
		}
		return fmt.Errorf("failed to write footer: %w", <-s.done)
	}

	// ParquetWriter likes to close the underlying writer, so we need to check if it's already closed
	if !s.wc.closed {
		if err := s.wc.Close(); err != nil {
			return err
		}
	}

	return <-s.done
}

func (c *Client) WriteTableStream(_ context.Context, handle any, upsert bool, msgs []*message.Insert) error {
	if len(msgs) == 0 {
		return nil
	}

	records := make([]arrow.Record, len(msgs))
	for i, msg := range msgs {
		if c.spec.Athena {
			records[i] = sanitizeRecordJSONKeys(msg.Record)
		} else {
			records[i] = msg.Record
		}
	}

	return handle.(*stream).h.WriteContent(records)
}

func (c *Client) Write(ctx context.Context, options plugin.WriteOptions, msgs <-chan message.Message) error {
	if err := c.writer.Write(ctx, msgs); err != nil {
		return err
	}
	return c.writer.Flush(ctx)
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
