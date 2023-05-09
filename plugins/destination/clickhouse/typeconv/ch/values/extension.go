package values

import (
	"github.com/apache/arrow/go/v13/arrow/array"
	"github.com/cloudquery/plugin-sdk/v2/types"
	"github.com/google/uuid"
)

func extensionValue(arr array.ExtensionArray) any {
	switch arr := arr.(type) {
	case *types.UUIDArray:
		return primitiveValue[uuid.UUID](arr)
	case *types.InetArray, *types.MacArray, *types.JSONArray:
		return valueStrData(arr)
	default:
		// we fallback here to string representation
		return valueStrData(arr)
	}
}
