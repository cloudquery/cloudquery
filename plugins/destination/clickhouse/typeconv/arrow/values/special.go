package values

import (
	"github.com/apache/arrow/go/v12/arrow/float16"
	"github.com/cloudquery/plugin-sdk/v2/types"
	"github.com/google/uuid"
)

func buildFloat16(builder primitiveBuilder[float16.Num], value *float32) {
	if value == (*float32)(nil) {
		builder.AppendNull()
		return
	}
	builder.Append(float16.New(*value))
}

func buildBinary(builder primitiveBuilder[[]byte], value *string) {
	if value == (*string)(nil) {
		builder.AppendNull()
		return
	}
	builder.Append([]byte(*value))
}

func buildUUID(builder *types.UUIDBuilder, value *uuid.UUID) {
	if value == (*uuid.UUID)(nil) {
		builder.AppendNull()
		return
	}
	builder.Append(*value)
}
