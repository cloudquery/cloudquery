package client

import (
	"context"
	"database/sql/driver"
	"fmt"
	"strings"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/apache/arrow/go/v13/arrow/array"
	"github.com/apache/arrow/go/v13/arrow/decimal128"
	"github.com/apache/arrow/go/v13/arrow/memory"
	"github.com/cloudquery/plugin-sdk/v4/scalar"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

const (
	readSQL = "SELECT %s FROM %s"
)

func (c *Client) Read(ctx context.Context, table *schema.Table, res chan<- arrow.Record) error {
	colNames := make([]string, 0, len(table.Columns))
	for _, col := range table.Columns {
		colNames = append(colNames, pgx.Identifier{col.Name}.Sanitize())
	}
	cols := strings.Join(colNames, ",")
	tableName := table.Name
	sql := fmt.Sprintf(readSQL, cols, pgx.Identifier{tableName}.Sanitize())
	rows, err := c.conn.Query(ctx, sql)
	if err != nil {
		return err
	}
	for rows.Next() {
		values, err := rows.Values()
		if err != nil {
			return err
		}

		arrowSchema := table.ToArrowSchema()
		rb := array.NewRecordBuilder(memory.DefaultAllocator, arrowSchema)
		for i := range values {
			val, err := prepareValueForResourceSet(arrowSchema.Field(i).Type, values[i])
			if err != nil {
				return err
			}
			s := scalar.NewScalar(arrowSchema.Field(i).Type)
			if err := s.Set(val); err != nil {
				return err
			}
			scalar.AppendToBuilder(rb.Field(i), s)
		}
		res <- rb.NewRecord()
	}
	rows.Close()
	return nil
}

func prepareValueForResourceSet(dataType arrow.DataType, v any) (any, error) {
	if v == nil {
		return nil, nil
	}
	switch tp := dataType.(type) {
	case *arrow.Decimal128Type:
		if vt, ok := v.(pgtype.Numeric); ok {
			if vt.Valid {
				v = decimal128.FromBigInt(vt.Int)
			} else {
				v = nil
			}
		}
	case *arrow.ListType:
		vl := v.([]any)
		for i := range vl {
			var err error
			vl[i], err = prepareValueForResourceSet(tp.Elem(), vl[i])
			if err != nil {
				return nil, err
			}
		}
		return vl, nil
	case *arrow.LargeListType:
		vl := v.([]any)
		for i := range vl {
			var err error
			vl[i], err = prepareValueForResourceSet(tp.Elem(), vl[i])
			if err != nil {
				return nil, err
			}
		}
		return vl, nil
	case *arrow.StringType:
		if value, ok := v.(driver.Valuer); ok {
			if value == driver.Valuer(nil) {
				v = nil
			} else {
				val, err := value.Value()
				if err != nil {
					return nil, err
				}
				if s, ok := val.(string); ok {
					v = s
				}
			}
		}
		return v, nil
	case *arrow.Time32Type:
		switch vt := v.(type) {
		case pgtype.Time:
			t, err := vt.TimeValue()
			if err != nil {
				return nil, err
			}
			v = stringForTime(t, tp.Unit)
			return v, nil
		case string:
			v = vt
		}
	case *arrow.Time64Type:
		switch vt := v.(type) {
		case pgtype.Time:
			t, err := vt.TimeValue()
			if err != nil {
				return nil, err
			}
			v = stringForTime(t, tp.Unit)
			return v, nil
		case string:
			v = vt
		}
	case *arrow.Uint64Type:
		if vt, ok := v.(pgtype.Numeric); ok {
			if !vt.Valid {
				v = nil
			} else {
				v = vt.Int.Uint64()
			}
		}
		return v, nil
	}
	return v, nil
}

func stringForTime(t pgtype.Time, unit arrow.TimeUnit) string {
	extra := ""
	hour := t.Microseconds / 1e6 / 60 / 60
	minute := t.Microseconds / 1e6 / 60 % 60
	second := t.Microseconds / 1e6 % 60
	micros := t.Microseconds % 1e6
	switch unit {
	case arrow.Millisecond:
		extra = fmt.Sprintf(".%03d", (micros)/1e3)
	case arrow.Microsecond:
		extra = fmt.Sprintf(".%06d", micros)
	case arrow.Nanosecond:
		// postgres doesn't support nanosecond precision
		extra = fmt.Sprintf(".%06d", micros)
	}

	return fmt.Sprintf("%02d:%02d:%02d"+extra, hour, minute, second)
}
