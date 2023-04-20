package arrow

import (
	"fmt"
	"strings"
	"time"

	"github.com/ClickHouse/clickhouse-go/v2/lib/column"
	"github.com/apache/arrow/go/v12/arrow"
)

func mapType(_map *column.Map) (*arrow.MapType, error) {
	kv := strings.SplitN(params(_map.Type()), ",", 2)
	if len(kv) != 2 {
		return nil, fmt.Errorf("unsupported CickHouse map type: %s", _map.Type())
	}

	keyCol, err := column.Type(strings.TrimSpace(kv[0])).Column(_map.Name(), time.UTC)
	if err != nil {
		return nil, err
	}
	valCol, err := column.Type(strings.TrimSpace(kv[1])).Column(_map.Name(), time.UTC)
	if err != nil {
		return nil, err
	}

	keyField, err := fieldFromColumn(keyCol)
	if err != nil {
		return nil, err
	}
	valField, err := fieldFromColumn(valCol)
	if err != nil {
		return nil, err
	}

	return arrow.MapOf(keyField.Type, valField.Type), nil
}
