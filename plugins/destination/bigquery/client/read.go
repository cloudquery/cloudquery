package client

import (
	"context"
	"fmt"
	"strings"
	"time"

	"cloud.google.com/go/bigquery"
	"cloud.google.com/go/civil"
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/apache/arrow/go/v13/arrow/array"
	"github.com/apache/arrow/go/v13/arrow/memory"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/goccy/go-json"
	"google.golang.org/api/iterator"
)

const (
	readSQL = "SELECT %s FROM `%s.%s.%s` WHERE `_cq_source_name` = @cq_source_name order by _cq_sync_time asc"
)

func (c *Client) Read(ctx context.Context, table *schema.Table, sourceName string, res chan<- arrow.Record) error {
	colSQL := "`" + strings.Join(table.Columns.Names(), "`, `") + "`"
	stmt := fmt.Sprintf(readSQL, colSQL, c.pluginSpec.ProjectID, c.pluginSpec.DatasetID, table.Name)
	q := c.client.Query(stmt)
	q.Parameters = []bigquery.QueryParameter{
		{
			Name:  "cq_source_name",
			Value: sourceName,
		},
	}
	q.Location = c.client.Location
	it, err := q.Read(ctx)
	if err != nil {
		return fmt.Errorf("failed to read table %s: %w", table.Name, err)
	}
	arrowSchema := table.ToArrowSchema()
	for {
		values := make([]bigquery.Value, len(table.Columns))
		err := it.Next(&values)
		if err == iterator.Done {
			break
		}
		if err != nil {
			return fmt.Errorf("failed to read from table %s: %w", table.Name, err)
		}
		rb := array.NewRecordBuilder(memory.DefaultAllocator, arrowSchema)
		for i := range values {
			readItem(values, rb, i)
		}
		res <- rb.NewRecord()
	}
	return nil
}

func readItem(values []bigquery.Value, rb *array.RecordBuilder, i int) error {
	fieldBuilder := rb.Field(i)
	value := values[i]
	return appendValue(fieldBuilder, value)
}

func appendValue(bldr array.Builder, value any) error {
	switch arr := bldr.(type) {
	case *array.ListBuilder:
		lst := value.([]bigquery.Value)
		for _, v := range lst {
			arr.Append(true)
			if err := appendValue(arr.ValueBuilder(), v); err != nil {
				return err
			}
		}
	}
	switch rbv := bldr.(type) {
	case *array.Time32Builder:
		t := value.(civil.Time)
		unit := bldr.Type().(*arrow.Time32Type).Unit
		ts := stringForTime(t, unit)
		t32, err := arrow.Time32FromString(ts, unit)
		if err != nil {
			return err
		}
		rbv.Append(t32)
	case *array.Time64Builder:
		t := value.(civil.Time)
		unit := bldr.Type().(*arrow.Time64Type).Unit
		ts := stringForTime(t, unit)
		t64, err := arrow.Time64FromString(ts, unit)
		if err != nil {
			return err
		}
		rbv.Append(t64)
	case *array.DayTimeIntervalBuilder:
		t := value.([]bigquery.Value)
		rbv.Append(arrow.DayTimeInterval{
			Days:         int32(t[0].(int64)),
			Milliseconds: int32(t[1].(int64)),
		})
	case *array.MonthDayNanoIntervalBuilder:
		t := value.([]bigquery.Value)
		rbv.Append(arrow.MonthDayNanoInterval{
			Months:      int32(t[0].(int64)),
			Days:        int32(t[1].(int64)),
			Nanoseconds: t[2].(int64),
		})
	case *array.MonthIntervalBuilder:
		t := value.([]bigquery.Value)
		rbv.Append(arrow.MonthInterval(t[0].(int64)))
	case *array.TimestampBuilder:
		unit := bldr.Type().(*arrow.TimestampType).Unit
		switch unit {
		case arrow.Nanosecond:
			v := value.([]bigquery.Value)
			t := v[0].(time.Time)
			nano := v[1].(int64)
			nanoTime := t.Add(time.Duration(nano) * time.Nanosecond)
			ts, err := arrow.TimestampFromString(nanoTime.Format(time.RFC3339Nano), arrow.Nanosecond)
			if err != nil {
				return fmt.Errorf("failed to call arrow.TimestampFromString: %w", err)
			}
			rbv.Append(ts)
		default:
			t := value.(time.Time)
			ts, err := arrow.TimestampFromString(t.Format(time.RFC3339Nano), unit)
			if err != nil {
				return err
			}
			rbv.Append(ts)
		}
	default:
		// catch-all case to keep the code simple; this is only for testing
		// so performance is not a big concern
		b, _ := json.Marshal(value)
		r := strings.NewReader(string(b))
		d := json.NewDecoder(r)
		err := bldr.UnmarshalOne(d)
		if err != nil {
			return fmt.Errorf("failed to unmarshal value of type %v: %w. Value was %v", bldr.Type(), err, string(b))
		}
	}
	return nil
}

func stringForTime(t civil.Time, unit arrow.TimeUnit) string {
	extra := ""
	switch unit {
	case arrow.Millisecond:
		extra = fmt.Sprintf(".%03d", t.Nanosecond/1e6)
	case arrow.Microsecond:
		extra = fmt.Sprintf(".%06d", t.Nanosecond/1e3)
	case arrow.Nanosecond:
		extra = fmt.Sprintf(".%09d", t.Nanosecond)
	}
	return fmt.Sprintf("%02d:%02d:%02d"+extra, t.Hour, t.Minute, t.Second)
}
