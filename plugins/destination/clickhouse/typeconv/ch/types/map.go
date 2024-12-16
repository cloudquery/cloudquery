package types

import (
	"time"

	"github.com/ClickHouse/clickhouse-go/v2/lib/column"
	"github.com/apache/arrow-go/v18/arrow"
)

// https://clickhouse.com/docs/en/sql-reference/data-types/map
func mapType(mapType *arrow.MapType) (string, error) {
	keyType, err := FieldType(mapType.KeyField())
	if err != nil {
		return "", err
	}

	// Ensure CH can handle this as map
	// String, Integer, LowCardinality, FixedString, UUID, Date, DateTime, Date32, Enum.
	colType, err := column.Type(keyType).Column("key", time.UTC)
	if err != nil {
		return "", err
	}
	switch colType.(type) {
	case *column.String:
	case *column.UInt8, *column.UInt16, *column.UInt32, *column.UInt64:
	case *column.Int8, *column.Int16, *column.Int32, *column.Int64:
	case *column.FixedString:
	case *column.UUID:
	case *column.Date, *column.DateTime, *column.Date32:
	case *column.Enum8, *column.Enum16:
	default:
		// default to string representation
		return "String", nil
	}

	itemType, err := FieldType(mapType.ItemField()) // also sets nullable, just what we want
	if err != nil {
		return "", err
	}

	return "Map(" + keyType + ", " + itemType + ")", nil
}
