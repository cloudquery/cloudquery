package client

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"path"
	"reflect"
	"regexp"
	"strings"
	"time"

	"github.com/apache/arrow/go/v12/arrow"
	"github.com/apache/arrow/go/v12/arrow/array"
	"github.com/apache/arrow/go/v12/arrow/memory"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/types"
	"github.com/google/uuid"
)

const (
	PathVarTable = "{{TABLE}}"
	PathVarUUID  = "{{UUID}}"
	YearVar      = "{{YEAR}}"
	MonthVar     = "{{MONTH}}"
	DayVar       = "{{DAY}}"
	HourVar      = "{{HOUR}}"
	MinuteVar    = "{{MINUTE}}"
)

var reInvalidJSONKey = regexp.MustCompile(`\W`)

func (c *Client) WriteTableBatch(ctx context.Context, arrowSchema *arrow.Schema, data []arrow.Record) error {
	if len(data) == 0 {
		return nil
	}
	tableName := schema.TableName(arrowSchema)

	var err error
	mem := memory.DefaultAllocator
	if c.pluginSpec.Athena {
		for i, record := range data {
			data[i], err = sanitizeRecordJSONKeys(mem, record)
			if err != nil {
				return fmt.Errorf("failed to sanitize JSON keys for Athena in table %v: %w", tableName, err)
			}
		}
	}

	var b bytes.Buffer
	w := io.Writer(&b)

	timeNow := time.Now().UTC()

	if err := c.Client.WriteTableBatchFile(w, arrowSchema, data); err != nil {
		return err
	}
	// we don't upload in parallel here because AWS sdk moves the burden to the developer, and
	// we don't want to deal with that yet. in the future maybe we can run some benchmarks and see if adding parallelization helps.
	r := io.Reader(&b)
	if _, err := c.uploader.Upload(ctx, &s3.PutObjectInput{
		Bucket: aws.String(c.pluginSpec.Bucket),
		Key:    aws.String(replacePathVariables(c.pluginSpec.Path, tableName, uuid.NewString(), timeNow)),
		Body:   r,
	}); err != nil {
		return err
	}

	return nil
}

// sanitizeRecordJSONKeys replaces all invalid characters in JSON keys with underscores. This is required
// for compatibility with Athena. It returns a new record, and the old one is released.
func sanitizeRecordJSONKeys(mem memory.Allocator, record arrow.Record) (arrow.Record, error) {
	newRecordBuilder := array.NewRecordBuilder(mem, record.Schema())
	defer record.Release()
	defer newRecordBuilder.Release()
	for i := 0; i < int(record.NumCols()); i++ {
		col := record.Column(i)
		if arrow.TypeEqual(col.DataType(), types.NewJSONType()) {
			for r := 0; r < int(record.NumRows()); r++ {
				obj := col.GetOneForMarshal(r)
				sanitizeJSONKeysForObject(obj)
				newRecordBuilder.Field(i).(*types.JSONBuilder).Append(obj)
			}
			continue
		}
		for r := 0; r < int(record.NumRows()); r++ {
			b, err := record.Column(i).MarshalJSON()
			if err != nil {
				return nil, fmt.Errorf("failed to marshal JSON for col %v: %w", record.Schema().Field(i).Name, err)
			}
			err = newRecordBuilder.Field(i).UnmarshalJSON(b)
			if err != nil {
				return nil, fmt.Errorf("failed to unmarshal JSON into row for col %v: %w", record.Schema().Field(i).Name, err)
			}
		}
	}
	return newRecordBuilder.NewRecord(), nil
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

func replacePathVariables(specPath, table, fileIdentifier string, t time.Time) string {
	name := strings.ReplaceAll(specPath, PathVarTable, table)
	name = strings.ReplaceAll(name, PathVarUUID, fileIdentifier)
	name = strings.ReplaceAll(name, YearVar, t.Format("2006"))
	name = strings.ReplaceAll(name, MonthVar, t.Format("01"))
	name = strings.ReplaceAll(name, DayVar, t.Format("02"))
	name = strings.ReplaceAll(name, HourVar, t.Format("15"))
	name = strings.ReplaceAll(name, MinuteVar, t.Format("04"))
	return path.Clean(name)
}
