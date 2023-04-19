package arrow

import (
	"github.com/ClickHouse/clickhouse-go/v2/lib/column"
	"github.com/apache/arrow/go/v12/arrow"
)

func structField(_struct *column.Tuple) (*arrow.Field, error) {
	panic("implement me")
}
