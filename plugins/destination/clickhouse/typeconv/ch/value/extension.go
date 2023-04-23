package value

import (
	"github.com/apache/arrow/go/v12/arrow/array"
	"github.com/cloudquery/plugin-sdk/v2/types"
	"github.com/google/uuid"
)

func extensionValue(arr array.ExtensionArray) (any, error) {
	switch arr := arr.(type) {
	case *types.UUIDArray:
		return marshalValue[uuid.UUID](arr), nil
	case *types.InetArray, *types.MacArray, *types.JSONArray:
		return valueStrData(arr), nil
	default:
		// we fallback here to string representation
		return valueStrData(arr), nil
	}
}
