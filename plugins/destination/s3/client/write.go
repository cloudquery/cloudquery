package client

import (
	"context"
	"io"
	"path"
	"reflect"
	"regexp"
	"strings"
	"time"

	"github.com/apache/arrow/go/v14/arrow"
	"github.com/apache/arrow/go/v14/arrow/array"
	"github.com/apache/arrow/go/v14/arrow/memory"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/cloudquery/filetypes/v4"
	"github.com/cloudquery/plugin-sdk/v4/message"
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

func (c *Client) WriteTable(ctx context.Context, msgs <-chan *message.WriteInsert) error {
	var s *filetypes.Stream

	for msg := range msgs {
		if s == nil {
			table := msg.GetTable()

			objKey := c.replacePathVariables(table.Name, uuid.NewString(), time.Now().UTC())

			var err error
			s, err = c.Client.StartStream(table, func(r io.Reader) error {
				_, err := c.uploader.Upload(ctx, &s3.PutObjectInput{
					Bucket: aws.String(c.spec.Bucket),
					Key:    aws.String(objKey),
					Body:   r,
				})
				return err
			})
			if err != nil {
				return err
			}
		}

		if c.spec.Athena {
			msg.Record = sanitizeRecordJSONKeys(msg.Record)
		}

		if err := s.Write([]arrow.Record{msg.Record}); err != nil {
			_ = s.FinishWithError(err)
			return err
		}
	}

	return s.Finish()
}

func (c *Client) Write(ctx context.Context, msgs <-chan message.WriteMessage) error {
	return c.writer.Write(ctx, msgs)
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

func (c *Client) replacePathVariables(table, fileIdentifier string, t time.Time) string {
	name := strings.ReplaceAll(c.spec.Path, PathVarTable, table)
	if strings.Contains(name, PathVarFormat) {
		e := string(c.spec.Format) + c.spec.Compression.Extension()
		name = strings.ReplaceAll(name, PathVarFormat, e)
	}
	name = strings.ReplaceAll(name, PathVarUUID, fileIdentifier)
	name = strings.ReplaceAll(name, YearVar, t.Format("2006"))
	name = strings.ReplaceAll(name, MonthVar, t.Format("01"))
	name = strings.ReplaceAll(name, DayVar, t.Format("02"))
	name = strings.ReplaceAll(name, HourVar, t.Format("15"))
	name = strings.ReplaceAll(name, MinuteVar, t.Format("04"))
	return path.Clean(name)
}
