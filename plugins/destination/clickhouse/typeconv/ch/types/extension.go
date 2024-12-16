package types

import (
	"github.com/apache/arrow-go/v18/arrow"
	"github.com/cloudquery/plugin-sdk/v4/types"
)

func extensionType(extensionType arrow.ExtensionType) string {
	switch extensionType.(type) {
	// https://clickhouse.com/docs/en/sql-reference/data-types/uuid
	case *types.UUIDType:
		return "UUID"

	// https://clickhouse.com/docs/en/sql-reference/data-types/string
	case *types.InetType, *types.MACType:
		return "String"
	case *types.JSONType:
		// ClickHouse can't handle values like [{"x":{"y":"z"}}] at the moment.
		// https://github.com/ClickHouse/ClickHouse/issues/46190
		return "String"

	default:
		return "String"
	}
}
