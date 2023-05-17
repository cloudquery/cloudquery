package client

import (
	"context"
	"fmt"
	"strings"

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
			switch rbv := rb.Field(i).(type) {
			case *array.Time32Builder:
				t := values[i].(civil.Time)
				unit := arrowSchema.Field(i).Type.(*arrow.Time32Type).Unit
				ts := stringForTime(t, unit)
				t32, err := arrow.Time32FromString(ts, unit)
				if err != nil {
					return err
				}
				rbv.Append(t32)
			case *array.Time64Builder:
				t := values[i].(civil.Time)
				unit := arrowSchema.Field(i).Type.(*arrow.Time64Type).Unit
				ts := stringForTime(t, unit)
				t64, err := arrow.Time64FromString(ts, unit)
				if err != nil {
					return err
				}
				rbv.Append(t64)
			case *array.DayTimeIntervalBuilder:
				t := values[i].([]bigquery.Value)
				rbv.Append(arrow.DayTimeInterval{
					Days:         int32(t[0].(int64)),
					Milliseconds: int32(t[1].(int64)),
				})
			case *array.MonthDayNanoIntervalBuilder:
				t := values[i].([]bigquery.Value)
				rbv.Append(arrow.MonthDayNanoInterval{
					Months:      int32(t[0].(int64)),
					Days:        int32(t[1].(int64)),
					Nanoseconds: t[2].(int64),
				})
			case *array.MonthIntervalBuilder:
				t := values[i].([]bigquery.Value)
				rbv.Append(arrow.MonthInterval(t[0].(int64)))
			default:
				// catch-all case to keep the code simple; this is only for testing
				// so performance is not a big concern
				b, _ := json.Marshal(values[i])
				r := strings.NewReader(string(b))
				d := json.NewDecoder(r)
				err = rb.Field(i).UnmarshalOne(d)
				if err != nil {
					return fmt.Errorf("failed to unmarshal value for col %v of type %v: %w. Value was %v", table.Columns[i].Name, rb.Field(i).Type(), err, string(b))
				}
			}
		}
		res <- rb.NewRecord()
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
