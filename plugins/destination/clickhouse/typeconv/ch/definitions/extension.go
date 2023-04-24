package definitions

import (
	"github.com/apache/arrow/go/v12/arrow"
	"github.com/cloudquery/plugin-sdk/v2/types"
)

func extensionType(extension arrow.ExtensionType) string {
	switch extension.(type) {
	// https://clickhouse.com/docs/en/sql-reference/data-types/uuid
	case *types.UUIDType:
		return "UUID"

	// https://clickhouse.com/docs/en/sql-reference/data-types/string
	case *types.InetType, *types.MacType:
		return "String"
	case *types.JSONType:
		// ClickHouse can't handle values like [{"x":{"y":"z"}}] at the moment.
		// https://github.com/ClickHouse/ClickHouse/issues/46190
		return "String"

	default:
		return "String"
	}
}
