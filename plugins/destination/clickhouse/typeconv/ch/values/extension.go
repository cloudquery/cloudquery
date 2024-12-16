package values

import (
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/cloudquery/plugin-sdk/v4/types"
)

func extensionValue(arr array.ExtensionArray) any {
	switch arr := arr.(type) {
	case *types.UUIDArray:
		return primitiveValue(arr)
	case *types.InetArray, *types.MACArray, *types.JSONArray:
		return valueStrData(arr)
	default:
		// we fall back here to string representation
		return valueStrData(arr)
	}
}
