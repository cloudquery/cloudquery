package values

import (
	"github.com/apache/arrow/go/v14/arrow/array"
	"github.com/cloudquery/plugin-sdk/v4/types"
	"github.com/google/uuid"
)

func extensionValue(arr array.ExtensionArray) any {
	switch arr := arr.(type) {
	case *types.UUIDArray:
		return primitiveValue[uuid.UUID](arr)
	case *types.InetArray, *types.MACArray, *types.JSONArray:
		return valueStrData(arr)
	default:
		// we fall back here to string representation
		return valueStrData(arr)
	}
}
