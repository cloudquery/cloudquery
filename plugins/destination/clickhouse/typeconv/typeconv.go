package typeconv

import (
	"github.com/apache/arrow/go/v12/arrow"
	_arrow "github.com/cloudquery/cloudquery/plugins/destination/clickhouse/typeconv/arrow"
	"github.com/cloudquery/cloudquery/plugins/destination/clickhouse/typeconv/ch"
)

func ArrowField(name, typ string) (*arrow.Field, error) {
	return _arrow.Field(name, typ)
}

func ClickHouseDefinitions(fields ...arrow.Field) ([]string, error) {
	return ch.Definitions(fields...)
}
