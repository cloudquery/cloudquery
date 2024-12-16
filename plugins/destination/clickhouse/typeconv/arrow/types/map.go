package types

import (
	"fmt"
	"strings"
	"time"

	"github.com/ClickHouse/clickhouse-go/v2/lib/column"
	"github.com/apache/arrow-go/v18/arrow"
)

func mapType(name string, col *column.Map) (*arrow.Field, error) {
	// need to parse
	params := strings.SplitN(params(col.Type()), ",", 2)
	if len(params) < 2 {
		return nil, fmt.Errorf("unexpected ClickHouse Map type: %s", col.Type())
	}

	keyCol, err := column.Type(strings.TrimSpace(params[0])).Column(name, time.UTC)
	if err != nil {
		return nil, err
	}
	keyType, err := fieldFromColumn(keyCol)
	if err != nil {
		return nil, err
	}

	itemCol, err := column.Type(strings.TrimSpace(params[1])).Column(name, time.UTC)
	if err != nil {
		return nil, err
	}
	itemType, err := fieldFromColumn(itemCol)
	if err != nil {
		return nil, err
	}

	return &arrow.Field{Name: name, Type: arrow.MapOf(keyType.Type, itemType.Type)}, nil
}
