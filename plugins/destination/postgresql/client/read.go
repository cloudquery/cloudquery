package client

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/decimal128"
	"github.com/apache/arrow-go/v18/arrow/memory"
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
		if c.pgType == pgTypeCrateDB {
			colNames = append(colNames, pgx.Identifier{strings.Trim(col.Name, "_")}.Sanitize())
			continue
		}
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
			return stringForTime32(t, tp.Unit), nil
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
			return stringForTime64(t, tp.Unit), nil
		case string:
			v = vt
		}
	case *arrow.Uint64Type:
		if vt, ok := v.(pgtype.Numeric); ok {
			if !vt.Valid {
				v = nil
			} else {
				v = numericToUint64(vt)
			}
		}
		return v, nil
	}
	return v, nil
}

var (
	big10 = big.NewInt(10)
)

func numericToUint64(n pgtype.Numeric) uint64 {
	if n.Exp < 0 {
		panic("unsupported negative exponent")
	}

	val := n.Int.Uint64()
	if n.Exp == 0 {
		return val
	}

	// exp can only be positive for our use-case
	return new(big.Int).Mul(
		new(big.Int).SetUint64(val),
		new(big.Int).Exp(big10, big.NewInt(int64(n.Exp)), nil),
	).Uint64()
}

func stringForTime32(t pgtype.Time, unit arrow.TimeUnit) string {
	return arrow.Time32((time.Duration(t.Microseconds) * time.Microsecond) / unit.Multiplier()).FormattedString(unit)
}

func stringForTime64(t pgtype.Time, unit arrow.TimeUnit) string {
	return arrow.Time64((time.Duration(t.Microseconds) * time.Microsecond) / unit.Multiplier()).FormattedString(unit)
}
